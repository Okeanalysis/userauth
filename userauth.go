package main

import (
	"net/http"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)


type User struct{
	ID int `db:"id" json:"id"`
	Email string `db:"email" json:"email"`
	PasswordHash string `db:"password_hash" json:"-"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:updated_at json:updated_at`
}

var DB* sqlx.DB

func initDB() {
	var err error
	DB, err= sqlx.connect("postgres", "user= postgres dbname=database sslmode=disable")
	if err !=nil{
		panic(err)
	}

}

const jwtSecret= "eehicarburbur"

func Register(c *gin.Context) {
	var user User 
	if err:= c.ShouldBindJSON(&user); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedpassword, _:= 
bcrypt.GenerateFromPassword([]byte(user.PasswordHash)),
}
