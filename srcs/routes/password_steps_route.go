package routes

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Natt-S10/Natthaphon_agnos_backend_internship_2023/srcs/domain/models"
	logging "github.com/Natt-S10/Natthaphon_agnos_backend_internship_2023/srcs/services/logging"
	passwd_service "github.com/Natt-S10/Natthaphon_agnos_backend_internship_2023/srcs/services/passwd"
	"github.com/gin-gonic/gin"
)

func CorrectionResponse(ctx *gin.Context, db *sql.DB) {
	var passStepReq models.PasswordStepsReqest
	//if error

	if err := ctx.BindJSON(&passStepReq); err != nil {
		fmt.Print("Invalid Request: ", passStepReq.Password)
		// ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.String(http.StatusBadRequest, "Invalid Request or password string")
		lR := logging.MakeErrorLog(ctx, passStepReq.Password)
		logging.ExecLog(db, lR)
		return
	}

	steps := passwd_service.PasswordCorrectSteps(passStepReq.Password)

	passStepRes := models.PasswordStepsResponse{NumOfSteps: steps}
	ctx.IndentedJSON(http.StatusOK, passStepRes)

	lR := logging.MakeSuccessLog(ctx, passStepReq.Password, steps)
	logging.ExecLog(db, lR)
}

func GetStrongPasswordStepsHandler(db *sql.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		CorrectionResponse(ctx, db)
	}
}
