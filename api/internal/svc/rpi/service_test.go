package rpi

import (
	"database/sql"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/emilieanthony/senzr/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gotest.tools/assert"
)

func createFx(data *mockData) *Server {
	return &Server{
		Db: &database{
			data.CarbonDioxideData,
		},
	}
}

func TestRpiService_GetLatestCarbonDioxideEntry(t *testing.T) {
	t.Run("happy", func(t *testing.T) {
		timestamp := time.Date(2020, 5, 20, 20, 00, 00, 00, &time.Location{})
		data := []*carbonDioxideEntry{
			{
				Id:        "1234",
				Value:     15.0,
				Timestamp: timestamp,
			},
		}
		fx := createFx(&mockData{CarbonDioxideData: data})
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		fx.GetLatestCarbonDioxideEntry(c)
		assert.Equal(t, 200, w.Code)
		var got []*carbonDioxideEntry
		err := json.Unmarshal(w.Body.Bytes(), &got)
		assert.NilError(t, err)
		assert.DeepEqual(
			t,
			got,
			data,
		)
	})
}

type mockData struct {
	CarbonDioxideData []*carbonDioxideEntry
}

type database struct {
	CarbonDioxideData []*carbonDioxideEntry
}

type transaction struct {
	CarbonDioxideData []*carbonDioxideEntry
}

func (d *database) BeginTransaction() (_ db.Transaction, err error) {
	return &transaction{
		d.CarbonDioxideData,
	}, nil
}

func (t *transaction) Get(dest interface{}, query string, args ...interface{}) error {
	return nil
}
func (t *transaction) Select(dest interface{}, query string, args ...interface{}) error {
	pbs := dest.(*[]*carbonDioxideEntry)
	*pbs = append(*pbs, t.CarbonDioxideData...)
	return nil
}
func (t *transaction) NamedExec(query string, arg interface{}) (sql.Result, error) {
	return nil, nil
}
func (t *transaction) Query(query string) (*sqlx.Rows, error) {
	return &sqlx.Rows{}, nil
}
func (t *transaction) Commit() error {
	return nil
}
