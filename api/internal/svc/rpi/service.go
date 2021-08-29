package rpi

import (
	"net/http"
	"time"

	"github.com/emilieanthony/senzr/internal/db"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

type entry struct {
	Id        string     `json:"id" firestore:"id"`
	Value     float64    `json:"value" firestore:"value"`
	Timestamp *time.Time `json:"createdAt" firestore:"created_at"`
}

type Server struct{}

func (s *Server) Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

func (s *Server) GetLatestCarbonDioxideEntry(ctx *gin.Context) {
	client, err := db.NewClient(ctx)
	defer func() {
		client.Close()
	}()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "GetLatestCarbonDioxideEntry: creating client")
		return
	}
	query := client.Collection(db.CollectionCO2).OrderBy("created_at", firestore.Desc).Limit(1)
	document, err := query.Documents(ctx).Next()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "GetLatestCarbonDioxideEntry: getting from database")
		return
	}
	var data *entry
	if err := document.DataTo(&data); err != nil {
		ctx.String(http.StatusInternalServerError, "GetLatestCarbonDioxideEntry: reading from database")
		return
	}
	data.Id = document.Ref.ID
	ctx.JSON(http.StatusOK, data)
}

func (s *Server) GetLatestTemperatureEntry(ctx *gin.Context) {
	client, err := db.NewClient(ctx)
	defer func() {
		client.Close()
	}()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "GetLatestTemperatureEntry: creating client")
		return
	}
	query := client.Collection(db.CollectionTemperature).OrderBy("created_at", firestore.Desc).Limit(1)
	document, err := query.Documents(ctx).Next()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "GetLatestTemperatureEntry: getting from database")
		return
	}
	var data *entry
	if err := document.DataTo(&data); err != nil {
		ctx.String(http.StatusInternalServerError, "GetLatestTemperatureEntry: reading from database")
		return
	}
	data.Id = document.Ref.ID
	ctx.JSON(http.StatusOK, data)
}

func (s *Server) GetLatestHumidityEntry(ctx *gin.Context) {
	client, err := db.NewClient(ctx)
	defer func() {
		client.Close()
	}()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "GetLatestHumidityEntry: creating client")
		return
	}
	query := client.Collection(db.CollectionHumidity).OrderBy("created_at", firestore.Desc).Limit(1)
	document, err := query.Documents(ctx).Next()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "GetLatestHumidityEntry: getting from database")
		return
	}
	var data *entry
	if err := document.DataTo(&data); err != nil {
		ctx.String(http.StatusInternalServerError, "GetLatestHumidityEntry: reading from database")
		return
	}
	data.Id = document.Ref.ID
	ctx.JSON(http.StatusOK, data)
}
