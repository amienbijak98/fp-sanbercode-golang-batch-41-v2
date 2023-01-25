package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wisedevguy/fp-sanbercode-golang-batch-41-v2/database"
	"github.com/wisedevguy/fp-sanbercode-golang-batch-41-v2/models"
	"github.com/wisedevguy/fp-sanbercode-golang-batch-41-v2/repository"
)

func GetAllPatients(c *gin.Context) {
	var (
		result gin.H
	)

	patients, err := repository.GetAllPatients(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": patients,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetPatientByID(c *gin.Context) {
	var (
		result gin.H
	)

	idPatient, _ := strconv.Atoi(c.Param("id"))

	patients, err := repository.GetPatientByID(database.DbConnection, idPatient)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": patients,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertPatient(c *gin.Context) {
	var patient models.Patient

	err := c.ShouldBindJSON(&patient)
	if err != nil {
		panic(err)
	}

	err = repository.InsertPatient(database.DbConnection, patient)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success insert new patient",
	})
}

func UpdatePatient(c *gin.Context) {
	var patient models.Patient

	err := c.ShouldBindJSON(&patient)
	if err != nil {
		panic(err)
	}

	patient.ID, _ = strconv.Atoi(c.Param("id"))

	err = repository.UpdatePatient(database.DbConnection, patient)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success update patient",
	})
}

func DeletePatient(c *gin.Context) {
	idPatient, _ := strconv.Atoi(c.Param("id"))

	err := repository.DeletePatient(database.DbConnection, idPatient)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success delete patient",
	})
}
