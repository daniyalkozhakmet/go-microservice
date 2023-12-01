package handler

import (
	"errors"
	"fmt"
	"microservice/model"
	"microservice/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	UserRepo *repository.UserRepo
}

var secretKey = []byte("your-secret-key")

func (o *UserHandler) SignUp(c *gin.Context) {

	requestBody := model.UserRegister{}
	if err := c.BindJSON(&requestBody); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}
	fields := map[string]string{
		"email": requestBody.Email}
	exist, err := o.UserRepo.IsExist(fields)
	if err != nil {
		ErrorHandler(c, err)
		return
	}
	if exist {
		ErrorHandler(c, fmt.Errorf("user with that email already exists"))
		return
	}
	hashedPassword, err := o.HashPassword(requestBody.Password)
	if err != nil {
		ErrorHandler(c, err)
		return
	}
	requestBody.Password = hashedPassword
	o.UserRepo.Insert(requestBody)

	token, err := o.GenerateToken(requestBody.Email)
	if err != nil {
		ErrorHandler(c, fmt.Errorf("could not generate token"))
		return
	}
	o.SetCookie(c, token)
	
	c.JSON(200, gin.H{
		"user":  requestBody,
		"token": token,
	})
}
func (o *UserHandler) SignIn(c *gin.Context) {
	requestBody := model.UserLogin{}
	user := model.User{}
	if err := c.BindJSON(&requestBody); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	user, err := o.UserRepo.Get(requestBody.Email)
	if err != nil {
		ErrorHandler(c, err)
		return
	}
	passwordMatch := o.CheckPasswordHash(requestBody.Password, user.Password)
	if !passwordMatch {
		ErrorHandler(c, fmt.Errorf("invalid credentials"))
		return
	}
	token, err := o.GenerateToken(user.Email)
	if err != nil {
		ErrorHandler(c, fmt.Errorf("could not generate token"))
		return
	}
	o.SetCookie(c, token)
	c.JSON(200, gin.H{
		"user":  user,
		"token": token,
	})

}
func (o *UserHandler) UpdateByID(c *gin.Context) {
	fmt.Println("Update by id")
}
func (o *UserHandler) GenerateToken(email string) (string, error) {
	claims := model.CustomClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // Token expires in 24 hours
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func (o *UserHandler) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
func (o *UserHandler) CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
func (o *UserHandler) SetCookie(c *gin.Context, token string) {

	_, err := c.Cookie("auth_cookie")

	if err != nil {
		cookie := token
		c.SetCookie("gin_cookie", cookie, 3600, "/", "localhost", false, true)
	}
}
