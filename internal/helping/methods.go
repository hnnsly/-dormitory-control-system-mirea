package helping

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"hackaton/internal/keys"
	"hackaton/pkg/database"
	"hackaton/pkg/loggers"
	"hackaton/pkg/models"
	"net/http"
	"strconv"
	"time"
)

func CheckJWTAuth(c *gin.Context) (*models.User, error) {
	cookie, err := c.Request.Cookie("jwt")
	if err != nil {
		return nil, err
	}
	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(keys.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}
	claims := token.Claims.(*jwt.StandardClaims)
	var user models.User

	err = database.DB.QueryRow("SELECT id, username, password FROM users WHERE id = $1", claims.Issuer).
		Scan(&user.Id, &user.Name, &user.Email)

	if err != nil {
		loggers.ErrorLogger.Println(err)
		return nil, err
	}
	newToken, err := GenerateToken(strconv.Itoa(int(user.Id)))
	if err != nil {
		loggers.ErrorLogger.Println(err)

		return nil, err
	}

	cookie = &http.Cookie{
		Name:     "jwt",
		Value:    newToken,
		Expires:  time.Now().Add(time.Hour * 730),
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
	}

	http.SetCookie(c.Writer, cookie)
	return &user, nil
}

func GenerateToken(issuer string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: time.Now().Add(time.Hour * 730).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(keys.SecretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}
