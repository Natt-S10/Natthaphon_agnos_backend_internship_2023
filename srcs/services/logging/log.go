package logging

import (
	"database/sql"
	"fmt"
	_ "fmt"
	"net/http"
	"time"

	"github.com/Natt-S10/Natthaphon_agnos_backend_internship_2023/srcs/domain/models"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	INSERTSTATEMENT = `INSERT INTO log
			VALUE ($1, $2, $3, $4, $5, $6)`
)

func MakeSuccessLog(ctx *gin.Context, password string, steps int) models.LogRecord {
	lR := models.LogRecord{
		Timestamp:    time.Now(),
		Route:        ctx.Request.URL.Path,
		Status:       http.StatusOK,
		InitPassword: password,
		NumOfSteps:   steps,
		Error:        0,
	}
	return lR
}
func MakeErrorLog(ctx *gin.Context, password string) models.LogRecord {
	lR := models.LogRecord{
		Timestamp:    time.Now(),
		Route:        ctx.Request.URL.Path,
		Status:       http.StatusBadRequest,
		InitPassword: password,
		NumOfSteps:   0,
		Error:        1,
	}
	return lR
}

func ExecLog(db *sql.DB, lR models.LogRecord) {
	_, err := db.Exec(INSERTSTATEMENT,
		time.Now(),
		lR.Timestamp,
		lR.Route,
		lR.Status,
		lR.InitPassword,
		lR.NumOfSteps,
		lR.Error)
	if err != nil {
		fmt.Println("Error occurs at logging request")
		panic(err)
	}
}
