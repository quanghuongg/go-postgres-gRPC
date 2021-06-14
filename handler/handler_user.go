package handler

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	models "go-postgres/model"
	"net/http"
	"strings"
	"time"
)

var jwtKey = []byte("abcdefghijklmnopq")

type Claims struct {
	Name string `json:"Name"`
	Id   uint    `json:"ID"`
	jwt.StandardClaims
}

func LoginUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable password=huongnq")
		if err != nil {
			panic("failed to connect database")
		}
		defer db.Close()
		var users []models.User

		db.Where("Name = ? AND ID = ?", user.Name, user.Id).Find(&users)
		if len(users) > 0 {
			token, _ := GenToken(users[0])
			c.JSON(http.StatusOK, gin.H{"message": token})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		}
	}
}

func RegisterUser(c *gin.Context) {
	var register models.User

	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable password=huongnq")
		if err != nil {
			panic("failed to connect database")
		}
		defer db.Close()

		person := models.User{
			Name: register.Name,
		}
		db.Create(&person)

		msg := fmt.Sprintf("Register User %s success", register.Name)
		c.JSON(http.StatusOK, gin.H{"message": msg})
	}
}

type Header struct {
	Authorization string `header:"Authorization"`
}

func InfoUser(c *gin.Context) {
	h := Header{}

	if err := c.ShouldBindHeader(&h); err != nil {
		c.JSON(200, err)
	}
	tokenHeader := h.Authorization

	splits := strings.Split(tokenHeader, " ") // Bearer jwt_token
	if len(splits) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	tokenPart := splits[1]
	tk := &Claims{}

	token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "InternalServerError"})
	}

	if token.Valid {
		c.JSON(http.StatusOK, gin.H{"message": token.Claims})
	}
}

func GenToken(user models.User) (string, error) {
	expirationTime := time.Now().Add(120 * time.Second)
	claims := &Claims{
		Name: user.Name,
		Id:   user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
