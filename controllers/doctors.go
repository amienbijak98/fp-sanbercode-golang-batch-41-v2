package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wisedevguy/fp-sanbercode-golang-batch-41-v2/database"
	"github.com/wisedevguy/fp-sanbercode-golang-batch-41-v2/models"
	"github.com/wisedevguy/fp-sanbercode-golang-batch-41-v2/repository"
)

func GetAllDoctors(c *gin.Context) {
	var (
		result gin.H
	)

	doctors, err := repository.GetAllDoctors(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": doctors,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetDoctorByID(c *gin.Context) {
	var (
		result gin.H
	)

	idDoctor, _ := strconv.Atoi(c.Param("id"))

	doctor, err := repository.GetDoctorByID(database.DbConnection, idDoctor)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": doctor,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertDoctor(c *gin.Context) {
	var doctor models.Doctor

	err := c.ShouldBindJSON(&doctor)
	if err != nil {
		panic(err)
	}

	err = repository.InsertDoctor(database.DbConnection, doctor)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success insert new doctor",
	})
}

func UpdateDoctor(c *gin.Context) {
	var doctor models.Doctor

	err := c.ShouldBindJSON(&doctor)
	if err != nil {
		panic(err)
	}

	doctor.ID, _ = strconv.Atoi(c.Param("id"))

	err = repository.UpdateDoctor(database.DbConnection, doctor)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success update doctor",
	})
}

func DeleteDoctor(c *gin.Context) {
	idDoctor, _ := strconv.Atoi(c.Param("id"))

	err := repository.DeleteDoctor(database.DbConnection, idDoctor)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success delete doctor",
	})
}
