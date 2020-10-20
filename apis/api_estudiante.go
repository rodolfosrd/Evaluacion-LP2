package apis

import (
	"net/http"
	"crudjos/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func EstudianteGetId(c *gin.Context) {
	id := c.Params.ByName("id")
	var est models.Estudiante
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	if err := conn.First(&est, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &est)
}

func EstudianteIndex(c *gin.Context) {
	var lis []models.Estudiante
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	conn.Find(&lis)
	c.JSON(http.StatusOK, lis)
}

func EstudiantePost(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	est := models.Estudiante{Name:		c.PostForm("name"), 
							Paternal: 	c.PostForm("paternal"), 
							Maternal: 	c.PostForm("maternal"),
							Age:		c.PostForm("age"),
							State:		c.PostForm("state"),
		}
	conn.Create(&est)
	c.JSON(http.StatusOK, &est)
}

func EstudiantePut(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var est models.Estudiante
	if err := conn.First(&est, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	est.Name = c.PostForm("name")
	est.Paternal = c.PostForm("paternal")
	est.Maternal = c.PostForm("maternal")
	conn.Save(&est)
	c.JSON(http.StatusOK, &est)
}

func EstudianteDelete(c *gin.Context){
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var est models.Estudiante
	if err := conn.Where("id = ?", id).First(&est).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Unscoped().Delete(&est)
}