package handler

import (
	"log"
	"net/http"

	"github.com/archon42x/agora/auth/logic"
	"github.com/archon42x/agora/common/errs"
	"github.com/archon42x/agora/common/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterHandler(c *gin.Context) {
	req := &RegisterRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		log.Printf("json parse error: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": errs.REGISTER_ERROR,
			"msg":  "json parse error",
		})
		return
	}

	username := req.Username
	if username == "" {
		log.Printf("username is empty\n")
		c.JSON(http.StatusOK, gin.H{
			"code": errs.REGISTER_ERROR,
			"msg":  "username is empty",
		})
		return
	}
	_, err := logic.FindUserByUsername(username)
	if err == nil {
		log.Printf("username already exists\n")
		c.JSON(http.StatusOK, gin.H{
			"code": errs.REGISTER_ERROR,
			"msg":  "username already exists",
		})
		return
	} else if err != gorm.ErrRecordNotFound {
		log.Printf("find user error: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": errs.LOGIN_ERROR,
			"msg":  "find user error",
		})
		return
	}

	password := req.Password
	if password == "" {
		log.Printf("password is empty\n")
		c.JSON(http.StatusOK, gin.H{
			"code": errs.REGISTER_ERROR,
			"msg":  "password is empty",
		})
		return
	}

	encPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("encode password error: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": errs.REGISTER_ERROR,
			"msg":  "encode password error",
		})
		return
	}

	_, err = logic.CreateUser(&model.User{
		Username: username,
		Password: string(encPassword),
		Role:     model.UserRoleUser,
	})
	if err != nil {
		log.Printf("create user error: %v\n", err)
		c.JSON(http.StatusOK, gin.H{
			"code": errs.LOGIN_ERROR,
			"msg":  "create user error",
		})
		return
	}

	log.Printf("register success\n")
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "register success",
	})
}
