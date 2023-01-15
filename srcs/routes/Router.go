package routes

import (
	"fmt"
	"net/http"

	"github.com/Natt-S10/Natthaphon_agnos_backend_internship_2023/srcs/domain/models"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Group: api
	apiGroup := r.Group("/api")
	{
		apiGroup.POST("/strong_password_steps", func(ctx *gin.Context) {
			var passStepReq models.PasswordStepsReqest

			if err := ctx.BindJSON(&passStepReq); err != nil {
				fmt.Print("shit happened", passStepReq.Password)
				return
			}

			ctx.String(http.StatusOK, "Oh u found me"+passStepReq.Password)
		})
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
