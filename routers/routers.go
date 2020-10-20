package routers

import (
	"crudjos/apis"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"crudjos/models"
)

func SetupRouter() *gin.Engine {

	conn, err := connectDB()
	if err != nil {
		panic("failed to connect database"+err.Error())
	}

	r := gin.Default()
	r.Use(dbMiddleware(*conn))

	stu := r.Group("/stu")
	{
		stu.GET("/students/:id", apis.EstudianteGetId)
		stu.GET("/students/", apis.EstudianteIndex)
		stu.POST("/students/", apis.EstudiantePost)
		stu.PUT("/students/:id", apis.EstudiantePut)
		stu.DELETE("/students/:id", apis.EstudianteDelete)
	}
	cour := r.Group("/cour")
	{
		cour.GET("/courses/:id", 	apis.CursosGetId)
		cour.GET("/courses/", 		apis.CursosIndex)
		cour.POST("/courses/", 		apis.CursosPost)
		cour.PUT("/courses/:id", 	apis.CursosPut)
		cour.DELETE("/courses/:id", apis.CursosDelete)
	}

	return r
}

func connectDB() (c *gorm.DB, err error) {

	dsn := "docker:docker@tcp(localhost:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	conn.AutoMigrate(&models.Cursos{})
	conn.AutoMigrate(&models.Estudiante{})

	if err != nil {
		panic("failed to connect database"+err.Error())
	}
	return conn, err
}

func dbMiddleware(conn gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", conn)
		c.Next()
	}
}