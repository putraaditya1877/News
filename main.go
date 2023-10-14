package main

import (
	"backend-berita/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// CreateData

	// ShowData

	// GetData

	// UpdateData

	// DeleteData

	r.Run()
}
