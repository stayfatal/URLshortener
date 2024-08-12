package server

import (
	"net/http"
	"url/internal/auth"
	"url/internal/database"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (s Server) CreateUserHandler(c *gin.Context) {
	user := database.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	bcPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	user.Password = string(bcPassword)

	id, err := s.DBManager.CreateUser(user)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	token, err := auth.CreateNewToken(id)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token:": token})
}

func (s *Server) LoginHandler(c *gin.Context) {
	user := database.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := s.DBManager.VerifyPassword(user.Email, user.Password); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	token, err := auth.CreateNewToken(user.ID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token:": token})
}
