package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wisedevguy/fp-sanbercode-golang-batch-41-v2/database"
	"github.com/wisedevguy/fp-sanbercode-golang-batch-41-v2/models"
	"github.com/wisedevguy/fp-sanbercode-golang-batch-41-v2/repository"
)

func GetAllAppointments(c *gin.Context) {
	var (
		result gin.H
	)

	appointment, err := repository.GetAllAppointments(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": appointment,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetAppointmentByID(c *gin.Context) {
	var (
		result gin.H
	)

	idAppoint, _ := strconv.Atoi(c.Param("id"))

	appointment, err := repository.GetAppointmentByID(database.DbConnection, idAppoint)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": appointment,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertAppointment(c *gin.Context) {
	var appointment models.Appointment

	err := c.ShouldBindJSON(&appointment)
	if err != nil {
		panic(err)
	}

	err = repository.InsertAppointment(database.DbConnection, appointment)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success insert new appointment",
	})
}

func UpdateAppointment(c *gin.Context) {
	var appointment models.Appointment

	err := c.ShouldBindJSON(&appointment)
	if err != nil {
		panic(err)
	}

	appointment.ID, _ = strconv.Atoi(c.Param("id"))

	err = repository.UpdateAppointment(database.DbConnection, appointment)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success update appointment",
	})
}

func DeleteAppointment(c *gin.Context) {
	idAppoint, _ := strconv.Atoi(c.Param("id"))

	err := repository.DeleteAppointment(database.DbConnection, idAppoint)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success delete appointment",
	})
}
