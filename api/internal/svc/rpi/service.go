package rpi

import (
	"fmt"
	"net/http"
	"time"

	"github.com/emilieanthony/senzr/internal/db"
	"github.com/gin-gonic/gin"
)

type carbonDioxideEntry struct {
	Id        string    `json:"id" db:"id"`
	Value     float64   `json:"value" db:"value"`
	Timestamp time.Time `json:"createdAt" db:"created_at"`
}

type Server struct {
	Db db.Database
}

func (s *Server) HelloWorld(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello world!")
}

func (s *Server) GetLatestCarbonDioxideEntry(ctx *gin.Context) {
	tx, err := s.Db.BeginTransaction()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "GetLatestDataEntry: connecting to database")
		return
	}
	query := "SELECT * FROM carbon_dioxide ORDER BY created_at DESC LIMIT 1"
	data := make([]*carbonDioxideEntry, 0)
	if err := tx.Select(&data, query); err != nil {
		ctx.String(http.StatusInternalServerError, "GetLatestDataEntry: getting from database")
		return
	}
	fmt.Printf("data: %v \n", data)
	if err := tx.Commit(); err != nil {
		ctx.String(http.StatusInternalServerError, "GetLatestDataEntry: committing transaction")
	}

	ctx.JSON(http.StatusOK, data)
}
