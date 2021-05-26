package cloud_receiver

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

const (
	driver = "postgres"
)

type Data struct {
	Co2         float64 `json:"co2"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
	Timestamp   string  `json:"timestamp"`
}

// Message can contain any of these fields:
// https://cloud.google.com/pubsub/docs/reference/rest/v1/PubsubMessage
type Message struct {
	Data []byte `json:"data"`
}

func SensorDataReceiver(ctx context.Context, m Message) error {
	var data Data
	if err := json.Unmarshal(m.Data, &data); err != nil {
		return fmt.Errorf("error unmarshalling data: %w", err)
	}

	db, err := newDBConnection()
	if err != nil {
		return fmt.Errorf("could not connect to database: %w", err)
	}
	tx, err := db.Beginx()
	if err != nil {
		return fmt.Errorf("could not start transaction: %w", err)
	}
	tx.MustExec("INSERT INTO carbon_dioxide(value, created_at) VALUES($1, $2)", data.Co2, data.Timestamp)
	tx.MustExec("INSERT INTO temperature(value, created_at) VALUES($1, $2)", data.Temperature, data.Timestamp)
	tx.MustExec("INSERT INTO humidity(value, created_at) VALUES($1, $2)", data.Humidity, data.Timestamp)

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}
	return nil
}

func newDBConnection() (_ *sqlx.DB, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("connection error: %v", err)
		}
	}()
	env, err := Env()
	if err != nil {
		return nil, fmt.Errorf("getting environment: %w", err)
	}
	db, err := sqlx.Connect(
		driver,
		fmt.Sprintf("user=%s password=%s database=%s host=%s",
			env.DBUser, env.DBPassword, env.DBName, env.DBHost,
		),
	)
	if err != nil {
		return nil, fmt.Errorf("connecting: %w", err)
	}
	return db, nil
}

type environment struct {
	DBHost     string `envconfig:"db_host"`     // SENZR_DB_HOST
	DBName     string `envconfig:"db_name"`     // SENZR_DB_NAME
	DBPort     string `envconfig:"db_port"`     // SENZR_DB_PORT
	DBUser     string `envconfig:"db_user"`     // SENZR_DB_USER
	DBPassword string `envconfig:"db_password"` // SENZR_DB_PASSWORD
}

func Env() (*environment, error) {
	var e environment
	err := envconfig.Process("senzr", &e)
	if err != nil {
		return nil, err
	}
	return &e, nil
}
