package server

import (
	"fmt"
	"go-world/router"
)

func Start() {
	r := router.SetupRouter()
	fmt.Println("Server is running on port 8080")
	r.Run(":8080")
}
