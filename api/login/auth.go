package login

import (
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"hackaton/log"
	"hackaton/storage"
	"hackaton/types"
	"hackaton/utils"
	"net/http"
	"strconv"
	"time"
)

func Register(c *gin.Context) {
	var data map[string]string

	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{"message": "invalid request"})
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := types.User{
		Email:    data["email"],
		Password: []byte(string(password)),
	}

	err := storage.Store.AddUser(&user)

	if err != nil {
		log.ErrorLogger.Println(err)
		c.JSON(500, gin.H{"message": "could not register user"})
		return
	}

	c.JSON(200, user)
}

func Login(c *gin.Context) {
	var data map[string]string

	if err := c.BindJSON(&data); err != nil {
		log.ErrorLogger.Println(err)
		c.JSON(400, gin.H{"message": "invalid request"})
		return
	}

	var user types.User
	err := storage.Store.Db.QueryRow("SELECT id, username, password FROM users WHERE username = $1", data["email"]).
		Scan(&user.Id, &user.Email, &user.Password)

	if err == sql.ErrNoRows {
		log.ErrorLogger.Println(err)
		c.JSON(404, gin.H{"message": "user not found"})
		return
	} else if err != nil {
		log.ErrorLogger.Println(err)
		c.JSON(500, gin.H{"message": "could not retrieve user"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"]))

	if err != nil {
		log.ErrorLogger.Println(err)
		c.JSON(400, gin.H{"message": "incorrect password"})
		return
	}

	token, err := utils.GenerateToken(strconv.Itoa(int(user.Id)))

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
	c.Redirect(302, "/students/find")
}

func User(c *gin.Context) {
	cookie, err := c.Request.Cookie("jwt")
	if err != nil {
		log.ErrorLogger.Println(err)
		c.JSON(401, gin.H{"message": "unauthenticated"})
		return
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.SecretKey), nil
	})

	if err != nil {
		log.ErrorLogger.Println(err)
		c.JSON(401, gin.H{"message": "unauthenticated"})
		return
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user types.User

	err = storage.Store.SearchUser(&user, claims)

	if err != nil {
		log.ErrorLogger.Println(err)
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
