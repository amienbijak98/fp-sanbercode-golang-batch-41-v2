package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wisedevguy/fp-sanbercode-golang-batch-41-v2/database"
	"github.com/wisedevguy/fp-sanbercode-golang-batch-41-v2/models"
	"github.com/wisedevguy/fp-sanbercode-golang-batch-41-v2/repository"
)

func Register(c *gin.Context) {
	var admin models.UserAdmin

	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := repository.Register(database.DbConnection, admin)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Register Success!"})
}

func Login(c *gin.Context) {

	var admin models.UserAdmin

	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := repository.Login(database.DbConnection, admin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

func Logout(c *gin.Context) {
	uuid, _ := c.Get("uuid")
	err := repository.Logout(database.DbConnection, uuid.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "admin has been logout"})
}

func GetAllAdmin(c *gin.Context) {
	var (
		result gin.H
	)

	admins, err := repository.GetAllAdmin(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": admins,
		}
	}

	c.JSON(http.StatusOK, result)
}
