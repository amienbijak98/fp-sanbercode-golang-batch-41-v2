package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wisedevguy/fp-sanbercode-golang-batch-41-v2/database"
	"github.com/wisedevguy/fp-sanbercode-golang-batch-41-v2/models"
	"github.com/wisedevguy/fp-sanbercode-golang-batch-41-v2/repository"
)

func GetAllDepartements(c *gin.Context) {
	var (
		result gin.H
	)

	departements, err := repository.GetAllDepartements(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": departements,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetDepartementByID(c *gin.Context) {
	var (
		result gin.H
	)

	idDepartement, _ := strconv.Atoi(c.Param("id"))

	departement, err := repository.GetDepartementByID(database.DbConnection, idDepartement)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": departement,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertDepartement(c *gin.Context) {
	var departement models.Departement

	err := c.ShouldBindJSON(&departement)
	if err != nil {
		panic(err)
	}

	err = repository.InsertDepartement(database.DbConnection, departement)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success insert new departement",
	})
}

func UpdateDepartement(c *gin.Context) {
	var departement models.Departement

	err := c.ShouldBindJSON(&departement)
	if err != nil {
		panic(err)
	}

	departement.ID, _ = strconv.Atoi(c.Param("id"))

	err = repository.UpdateDepartement(database.DbConnection, departement)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success update departement",
	})
}

func DeleteDepartement(c *gin.Context) {
	idDepartement, _ := strconv.Atoi(c.Param("id"))

	err := repository.DeleteDepartement(database.DbConnection, idDepartement)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success delete departement",
	})
}
