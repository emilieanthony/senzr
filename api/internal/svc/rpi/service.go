package rpi

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"google.golang.org/api/iterator"

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

type dailyAverage struct {
	Value float64 `json:"value"`
}

func (s *Server) GetDailyAverageCarbonDioxide(ctx *gin.Context) {
	client, err := db.NewClient(ctx)
	defer func() {
		client.Close()
	}()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "GetDailyAverageCarbonDioxide: creating client")
		return
	}
	year, month, day := time.Now().Date()
	midnight := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	query := client.Collection(db.CollectionCO2).Where("created_at", ">=", midnight)
	documents, err := query.Documents(ctx).GetAll()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "could not get CO2 data")
	}
	totalRecordsCount := len(documents)
	totalCo2 := 0
	for _, doc := range documents {
		var e *entry
		if err := doc.DataTo(&e); err != nil {
			ctx.String(http.StatusInternalServerError, "could not get CO2 data")
			return
		}
		totalCo2 += int(e.Value)
	}
	if totalRecordsCount == 0 {
		totalRecordsCount = 1
	}
	average := dailyAverage{
		Value: float64(totalCo2) / float64(totalRecordsCount),
	}
	ctx.JSON(http.StatusOK, average)
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

// GetDurationCarbonDioxide takes a query parameter "seconds" that is the duration from the current time
// to get CO2 data over.
// Defaults to "43200" (12h) if not set. Max is 2592000 (30 days).
func (s *Server) GetDurationCarbonDioxide(ctx *gin.Context) {
	const maxDurationSeconds = 2592000
	durationParam := ctx.DefaultQuery("seconds", "43200")
	durationSeconds, err := strconv.Atoi(durationParam)
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid duration. Duration can only contain numbers")
		return
	}
	if durationSeconds > maxDurationSeconds {
		ctx.String(
			http.StatusBadRequest,
			fmt.Sprintf("invalid duration. Duration cannot be greater than %d seconds", 2592000),
		)
		return
	}
	client, err := db.NewClient(ctx)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "could not get CO2 data")
		return
	}
	defer func() {
		client.Close()
	}()
	duration := time.Duration(durationSeconds) * time.Second
	now := time.Now()
	query := client.Collection(db.CollectionCO2).Where("created_at", ">", now.Add(-duration)).OrderBy("created_at", firestore.Asc)
	iter := query.Documents(ctx)
	defer iter.Stop()
	data := make([]*entry, 0)
	i := 0
	for {
		i++
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			ctx.String(http.StatusInternalServerError, "could not get CO2 data")
			return
		}
		var e *entry
		if err := doc.DataTo(&e); err != nil {
			ctx.String(http.StatusInternalServerError, "could not get CO2 data")
			return
		}
		if i % 20 == 0 {
			e.Value = math.Round(e.Value)
			data = append(data, e)
		}
	}
	ctx.JSON(http.StatusOK, data)
}
