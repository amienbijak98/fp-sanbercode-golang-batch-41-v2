package middleware

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/wisedevguy/fp-sanbercode-golang-batch-41-v2/database"
	// "github.com/appleboy/gin-jwt"
)

var (
	secret = os.Getenv("JWT_SECRET")
)

func VerifyToken(c *gin.Context) (interface{}, error) {
	errCode := errors.New("please use correct authentication")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")
	if !bearer {
		log.Printf("Error 1 : %s", errCode)
		return nil, errCode
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, err := jwt.Parse(stringToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Printf("Error 2 : %s", errCode)
			return nil, errCode
		}
		return []byte(secret), nil
	})
	if err != nil {
		log.Printf("Error 3 : %s", err)
		return nil, err
	}

	if !token.Valid {
		log.Printf("Error 4 : %s", errCode)
		return nil, errCode
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok {
		log.Printf("Error 5 : %s", errCode)
		return nil, errCode
	}

	checkLogout := token.Claims.(jwt.MapClaims)
	sqlStatement := "SELECT * FROM user_credentials SET WHERE uuid = $1"

	db := database.DbConnection
	row := db.QueryRow(sqlStatement, checkLogout["uuid"])

	storedUserCredential := struct {
		ID      int
		User_id int
		Uuid    string
	}{}

	err = row.Scan(&storedUserCredential.ID, &storedUserCredential.User_id, &storedUserCredential.Uuid)
	if err != nil {
		return nil, errors.New("user has been logout")
	}

	// fmt.Println(token.Claims.(jwt.MapClaims))
	return token.Claims.(jwt.MapClaims), nil
}

func GenerateToken(id int, uuid string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix()
	claims["uuid"] = uuid
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifiedToken, err := VerifyToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"err":     "unauthorized",
				"message": err.Error(),
			})
		}
		data := verifiedToken.(jwt.MapClaims)
		c.Set("id", data["id"])
		c.Set("uuid", data["uuid"])
		c.Set("user_data", verifiedToken)
		c.Next()
	}
}
