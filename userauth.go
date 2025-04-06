package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int       `db:"id" json:"id"`
	Email        string    `db:"email" json:"email"`
	PasswordHash string    `db:"password_hash" json:"-"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:updated_at json:updated_at`
}

var DB *sqlx.DB

func initDB() {
	var err error
	DB, err = sqlx.connect("postgres", "user= postgres dbname=database sslmode=disable")
	if err != nil {
		panic(err)
	}

}

const jwtSecret = "eehicarburbur"

func Register(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ :=
		bcrypt.GenerateFromPassword([]byte(user.PasswordHash),
			bcrypt.DefaultCost)
	user.PasswordHash = string(hashedPassword)

	_, err := DB.NamedExec(`INSERT INTO users (email, password_hash)
VALUES (:email, :password_hash)`, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User already exists"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User registered succesfully"})

}

func Login(c *gin.Context) {
	var user User
	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		c
	}
}
