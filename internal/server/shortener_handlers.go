package server

import (
	"net/http"
	"url/internal/database"
	"url/internal/errors"
	"url/internal/shortener"

	"github.com/gin-gonic/gin"
)

func (s *Server) ShortenHandler(c *gin.Context) {
	link := database.Link{Domain: s.domain}
	err := c.ShouldBindJSON(&link)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if link.Alias == "" {
		link.Alias, err = shortener.GenerateShortLink(s.DBManager.CheckLink)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	} else {
		ok, err := s.DBManager.CheckLink(link.Domain + link.Alias)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		} else if ok {
			c.String(http.StatusBadRequest, errors.NotUniqueLink.Error())
			return
		}
	}

	err = s.DBManager.AddLink(link)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"your link:": link.Domain + link.Alias})
}

func (s *Server) RedirectHandler(c *gin.Context) {
	shortUrl := c.Params.ByName("link")
	url, err := s.DBManager.GetFullUrl(s.domain + shortUrl)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	}
	c.Redirect(http.StatusPermanentRedirect, url)
}
