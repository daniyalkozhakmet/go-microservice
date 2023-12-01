package template

import (
	"errors"
	"fmt"
	"microservice/handler"
	"microservice/model"
	"microservice/repository"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type TemplateUserHandler struct {
	UserRepo *repository.UserRepo
}

var secretKey = []byte("your-secret-key")

func (o *TemplateUserHandler) RegisterView(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}
func (o *TemplateUserHandler) LoginView(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}
func (o *TemplateUserHandler) Register(c *gin.Context) {

	requestBody := model.UserRegister{}
	if err := c.BindJSON(&requestBody); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]handler.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = handler.ErrorMsg{fe.Field(), handler.GetErrorMsg(fe)}
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}
	fields := map[string]string{
		"email": requestBody.Email}
	exist, err := o.UserRepo.IsExist(fields)
	if err != nil {
		handler.ErrorHandler(c, err)
		return
	}
	if exist {
		handler.ErrorHandler(c, fmt.Errorf("user with that email already exists"))
		return
	}
	hashedPassword, err := o.HashPassword(requestBody.Password)
	if err != nil {
		handler.ErrorHandler(c, err)
		return
	}
	requestBody.Password = hashedPassword
	o.UserRepo.Insert(requestBody)

	token, err := o.GenerateToken(requestBody.Email)
	if err != nil {
		handler.ErrorHandler(c, fmt.Errorf("could not generate token"))
		return
	}
	o.SetCookie(c, token)

	c.JSON(200, gin.H{
		"user":  requestBody,
		"token": token,
	})
}
func (o *TemplateUserHandler) Login(c *gin.Context) {
	requestBody := model.UserLogin{}
	user := model.User{}
	if err := c.BindJSON(&requestBody); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]handler.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = handler.ErrorMsg{fe.Field(), handler.GetErrorMsg(fe)}
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	user, err := o.UserRepo.Get(requestBody.Email)
	if err != nil {
		handler.ErrorHandler(c, err)
		return
	}
	passwordMatch := o.CheckPasswordHash(requestBody.Password, user.Password)
	if !passwordMatch {
		handler.ErrorHandler(c, fmt.Errorf("invalid credentials"))
		return
	}
	token, err := o.GenerateToken(user.Email)
	if err != nil {
		handler.ErrorHandler(c, fmt.Errorf("could not generate token"))
		return
	}
	o.SetCookie(c, token)
	c.JSON(200, gin.H{
		"user":  user,
		"token": token,
	})

}
func (o *TemplateUserHandler) GenerateToken(email string) (string, error) {
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
func (o *TemplateUserHandler) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
func (o *TemplateUserHandler) CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
func (o *TemplateUserHandler) SetCookie(c *gin.Context, token string) {

	_, err := c.Cookie("auth_cookie")

	if err != nil {
		cookie := token
		c.SetCookie("gin_cookie", cookie, 3600, "/", "localhost", false, true)
	}
}
