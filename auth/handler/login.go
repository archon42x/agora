package handler

import (
	"log"
	"net/http"

	"github.com/archon42x/agora/auth/logic"
	"github.com/archon42x/agora/common/errs"
	"github.com/archon42x/agora/common/jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(c *gin.Context) {
	req := &LoginRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		log.Printf("json parse error: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": errs.LOGIN_ERROR,
			"msg":  "json parse error",
		})
		return
	}

	username := req.Username
	if username == "" {
		log.Printf("username is empty\n")
		c.JSON(http.StatusOK, gin.H{
			"code": errs.LOGIN_ERROR,
			"msg":  "username is empty",
		})
		return
	}
	user, err := logic.FindUserByUsername(username)
	if err == gorm.ErrRecordNotFound {
		log.Printf("username not exists: %v\n", username)
		c.JSON(http.StatusOK, gin.H{
			"code": errs.LOGIN_ERROR,
			"msg":  "username not exists",
		})
		return
	} else if err != nil {
		log.Printf("find user exists: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": errs.LOGIN_ERROR,
			"msg":  "find user exists",
		})
		return
	}
	encPassword := user.Password
	password := req.Password
	if password == "" {
		log.Printf("password is empty\n")
		c.JSON(http.StatusOK, gin.H{
			"code": errs.LOGIN_ERROR,
			"msg":  "password is empty",
		})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(encPassword), []byte(password)) != nil {
		log.Printf("password incorrect\n")
		c.JSON(http.StatusOK, gin.H{
			"code": errs.LOGIN_ERROR,
			"msg":  "password incorrect",
		})
		return
	}

	token, err := jwt.GenerateToken(username)
	if err != nil {
		log.Printf("generate token error: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": errs.LOGIN_ERROR,
			"msg":  "generate token error",
		})
		return
	}

	log.Printf("login success\n")
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "login success",
		"data": token,
	})
}
