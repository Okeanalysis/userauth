package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int       `db:"id" json:"id"`
	Email        string    `db:"email" json:"email"`
	PasswordHash string    `db:"password_hash" json:"-"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}

type Project struct {
	ID     int `db:"id" json:"id"`
	UserID int `db:"user_id" json:"user_id"`
}

var DB *sqlx.DB

func initDB() {
	var err error
	DB, err = sqlx.Connect("postgres", "user= postgres password=Emmanuel247@ dbname=userauth sslmode=disable")
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dbuser User
	err := DB.Get(&dbuser, "SELECT * FROM users WHERE email=$1", user.Email)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(dbuser.PasswordHash), []byte(user.PasswordHash)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  dbuser.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenstring, _ := token.SignedString([]byte(jwtSecret))

	c.JSON(http.StatusOK, gin.H{"token": tokenstring})

}

func main() {
	initDB()
	defer DB.Close()

	r := gin.Default()

	r.POST("/register", Register)
	r.POST("/login", Login)

	r.Run(":8080")
}
