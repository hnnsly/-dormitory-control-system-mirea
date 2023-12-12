package loginPage

import (
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"hackaton/pkg/database"
	"hackaton/pkg/loggers"
	"hackaton/pkg/models"
	"net/http"
	"strconv"
	"time"
)

const SecretKey = "secret"

func Register(c *gin.Context) {
	var data map[string]string

	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{"message": "invalid request"})
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Email:    data["email"],
		Password: []byte(string(password)),
	}

	err := createUser(&user)

	if err != nil {
		loggers.ErrorLogger.Println(err)
		c.JSON(500, gin.H{"message": "could not register user"})
		return
	}

	c.JSON(200, user)
}

func createUser(user *models.User) error {
	_, err := database.DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)",
		user.Email, user.Password)
	if err != nil {
		return err
	}

	return err
}

func Login(c *gin.Context) {
	var data map[string]string

	if err := c.BindJSON(&data); err != nil {
		loggers.ErrorLogger.Println(err)
		c.JSON(400, gin.H{"message": "invalid request"})
		return
	}

	var user models.User
	err := database.DB.QueryRow("SELECT id, username, password FROM users WHERE username = $1", data["email"]).
		Scan(&user.Id, &user.Email, &user.Password)

	if err == sql.ErrNoRows {
		loggers.ErrorLogger.Println(err)
		c.JSON(404, gin.H{"message": "user not found"})
		return
	} else if err != nil {
		loggers.ErrorLogger.Println(err)
		c.JSON(500, gin.H{"message": "could not retrieve user"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"]))

	if err != nil {
		loggers.ErrorLogger.Println(err)
		c.JSON(400, gin.H{"message": "incorrect password"})
		return
	}

	token, err := generateToken(strconv.Itoa(int(user.Id)))

	if err != nil {
		c.JSON(500, gin.H{"message": "could not login"})
		return
	}

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 730),
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
	}

	http.SetCookie(c.Writer, cookie)
	//c.Writer.Header().Add("access-control-expose-headers", "Set-Cookie")

	c.JSON(200, gin.H{"message": "success"})
}

func generateToken(issuer string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: time.Now().Add(time.Hour * 730).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func User(c *gin.Context) {
	cookie, err := c.Request.Cookie("jwt")
	if err != nil {
		loggers.ErrorLogger.Println(err)
		c.JSON(401, gin.H{"message": "unauthenticated"})
		return
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		loggers.ErrorLogger.Println(err)
		c.JSON(401, gin.H{"message": "unauthenticated"})
		return
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	err = database.DB.QueryRow("SELECT id, username, password FROM users WHERE id = $1", claims.Issuer).
		Scan(&user.Id, &user.Name, &user.Email)

	if err != nil {
		loggers.ErrorLogger.Println(err)
		c.JSON(500, gin.H{"message": "could not retrieve user"})
		return
	}

	c.JSON(200, user)
}

func Logout(c *gin.Context) {
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(c.Writer, cookie)

	c.JSON(200, gin.H{"message": "success"})
}
