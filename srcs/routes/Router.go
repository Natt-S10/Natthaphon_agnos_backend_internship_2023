package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

const (
	DBTYPE     = "postgres"
	DBUSER     = "dev"
	DBNAME     = "log"
	DBPASSWORD = "12345678"
)

var Database *sqlx.DB

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	{
		var err error
		Database, err = sqlx.Connect(DBTYPE, fmt.Sprintf("user=%s dbname=%s password=%s", DBUSER, DBNAME, DBPASSWORD))
		if err != nil {
			fmt.Println("Fatal error on database connect")
		}
		defer Database.Close()
	}
	Database.MustExec("DROP TABLE IF EXISTS log;")

	r := gin.Default()

	// Group: api
	apiGroup := r.Group("/api")
	{
		apiGroup.POST("/strong_password_steps", CorrectionResponse)
	}

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// r.POST("/api/strong_password_steps", func(ctx *gin.Context) {
	// 	var passStepReq models.PasswordStepsReqest

	// 	if err := ctx.BindJSON(&passStepReq); err != nil {
	// 		fmt.Print("shit happened", passStepReq.Password)
	// 		return
	// 	}

	// 	ctx.String(http.StatusOK, "Oh u found me"+passStepReq.Password)
	// })

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/

	return r
}
