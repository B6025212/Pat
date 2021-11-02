package main

import (
	"github.com/b6025212/sa-64-example/controller"
	"github.com/b6025212/sa-64-example/entity"
	"github.com/b6025212/sa-64-example/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{
			// Nurses Routes
			protected.POST("/nurse", controller.CreateNurse)
			protected.GET("/nurse/:id", controller.GetNurse)
			protected.GET("/nurses", controller.ListNurses)

			// Diseases Routes
			protected.GET("/diseases", controller.ListDiseases)
			protected.GET("/disease/:id", controller.GetDisease)
			protected.POST("/disease", controller.CreateDisease)
			protected.PATCH("/disease", controller.UpdateDisease)
			protected.DELETE("/disease/:id", controller.DeleteDisease)

			// Medicines Routes
			protected.POST("/medicine", controller.CreateMedicine)
			protected.GET("/medicine/:id", controller.GetMedicine)
			protected.GET("/medicine", controller.ListMedicine)
			protected.DELETE("/medicine/:id", controller.DeleteMedicine)
			protected.PATCH("/medicine", controller.UpdateMedicine)

			// Patients Routes
			protected.GET("/patients", controller.ListPatients)
			protected.GET("/patient/:id", controller.GetPatient)
			protected.POST("/patient", controller.CreatePatient)
			protected.PATCH("/patient", controller.UpdatePatient)
			protected.DELETE("/patient/:id", controller.DeletePatient)

		}
	}

	// Nurse Routes
	r.POST("/nurses", controller.CreateNurse)

	// Authentication Routes
	r.POST("/login", controller.Login)

	// Run the server
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
