package main

import (
	"github.com/ainmtsn1999/go-restapi-login/controllers/logincontroller"
	"github.com/ainmtsn1999/go-restapi-login/controllers/usercontroller"
	"github.com/ainmtsn1999/go-restapi-login/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDB()

	//initialize static page
	r.Static("/log", "./views/login")
	r.Static("/reg", "./views/register")

	//login & register
	r.GET("/", logincontroller.Login)
	r.POST("/register", logincontroller.Register)

	//api 
	r.GET("/api/users", usercontroller.Index)
	r.GET("/api/user/:id", usercontroller.Show)
	r.POST("/api/user", usercontroller.Create)
	r.PUT("/api/user/:id", usercontroller.Update)
	r.DELETE("/api/user", usercontroller.Delete)

	r.Run()

}
