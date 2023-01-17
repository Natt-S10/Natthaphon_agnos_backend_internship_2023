package main

import (
	"github.com/Natt-S10/Natthaphon_agnos_backend_internship_2023/srcs/routes"
)

func main() {
	// r := routes.setupRouter()
	r := routes.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
