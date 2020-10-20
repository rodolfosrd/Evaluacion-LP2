package apis

import (
	"crudrod/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CursosGetId(c *gin.Context) {
	id := c.Params.ByName("id")
	var cur models.Cursos
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	if err := conn.First(&cur, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &cur)
}

func CursosIndex(c *gin.Context) {
	var lis []models.Cursos
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	conn.Find(&lis)
	c.JSON(http.StatusOK, lis)
}

func CursosPost(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	cur := models.Cursos{Name: c.PostForm("name"), Period: c.PostForm("period"), State: c.PostForm("state")}
	conn.Create(&cur)
	c.JSON(http.StatusOK, &cur)
}

func CursosPut(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var cur models.Cursos
	if err := conn.First(&cur, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	cur.Name = c.PostForm("name")
	cur.Period = c.PostForm("age")
	conn.Save(&cur)
	c.JSON(http.StatusOK, &cur)
}

func CursosDelete(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var cur models.Cursos
	if err := conn.Where("id = ?", id).First(&cur).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Unscoped().Delete(&cur)
}
