package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Natt-S10/Natthaphon_agnos_backend_internship_2023/srcs/domain/models"
	passwd_service "github.com/Natt-S10/Natthaphon_agnos_backend_internship_2023/srcs/services/passwd"
	"github.com/gin-gonic/gin"
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

func ExecLogUpdate(lR models.LogRecord) {
	Database.MustExec("INSERT INTO $1 (timestamp, route, status, initpassword, numofsteps, error) VALUES ($2, $3, $4, $5, $6, $7)",
		DBNAME, lR.Timestamp, lR.Route, lR.Status, lR.InitPassword, lR.NumOfSteps, lR.Error)
}

func CorrectionResponse(ctx *gin.Context) {
	var passStepReq models.PasswordStepsReqest
	//if error
	var lR models.LogRecord
	if err := ctx.BindJSON(&passStepReq); err != nil {
		fmt.Print("Invalid Request: ", passStepReq.Password)
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.String(http.StatusBadRequest, "Invalid Request or password string")
		lR = MakeErrorLog(ctx, passStepReq.Password)
		ExecLogUpdate(lR)
		return
	}

	steps := passwd_service.PasswordCorrectSteps(passStepReq.Password)

	passStepRes := models.PasswordStepsResponse{NumOfSteps: steps}
	ctx.IndentedJSON(http.StatusOK, passStepRes)

	lR = MakeSuccessLog(ctx, passStepReq.Password, steps)
	ExecLogUpdate(lR)

}
