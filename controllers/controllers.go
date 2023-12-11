package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/testing_ap/models"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.Ussankutti
		if err := c.Bind(&user); err != nil {
			c.JSON(500, gin.H{
				"message": "binding error",
			})
			return
		}

		err := db.Create(&models.Ussankutti{
			Name:  user.Name,
			Email: user.Email,
			Phone: user.Phone,
		}).Error

		if err != nil {
			c.JSON(501, gin.H{
				"error": "failed create user",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "successfully crated " + user.Name,
		})
	}
}

func GetUserById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var user models.Ussankutti
		if err := db.First(&user, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"user": user,
		})
	}
}

func UpdateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		var user models.Ussankutti

		if err := c.Bind(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "binding error",
			})
			return
		}
		err := db.Model(&models.Ussankutti{}).Where("id=?", id).Updates(map[string]interface{}{"name": user.Name, "email": user.Email, "phone": user.Phone}).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "successfully updated user",
		})
	}
}

func DeleteUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		err := db.Where("id=?", id).Delete(&models.Ussankutti{}).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "successfully delete user",
		})
	}
}
