package logincontroller

import (
	"net/http"

	"github.com/ainmtsn1999/go-restapi-login/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	var user models.User
	uname := c.Query("uname")
	pass := c.Query("pass")

	if err := models.DB.Where("uname = ? AND pass = ?", uname, pass).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	//c.String(http.StatusOK, "Success login %s", uname)
	c.JSON(http.StatusOK, gin.H{"message": "Successfully login"})
}

func Register(c *gin.Context) {
	var user models.User

	user.Name = c.PostForm("name")
	user.Uname = c.PostForm("uname")
	user.Pass = c.PostForm("pass")

	if models.DB.Create(&user).Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Fsailed to update data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully insert data"})

	//c.JSON(http.StatusOK, gin.H{"message": user})
}
