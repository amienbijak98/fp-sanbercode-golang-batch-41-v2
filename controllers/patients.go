package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wisedevguy/fp-sanbercode-golang-batch-41-v2/database"
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
