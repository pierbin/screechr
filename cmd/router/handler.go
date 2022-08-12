package router

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pierbin/screechr/internal/controller"
	"github.com/pierbin/screechr/internal/models"
)

type Handler struct {
	scCtl *controller.ScreechrCtl
}

func NewHandler(scCtl *controller.ScreechrCtl) *Handler {
	return &Handler{
		scCtl: scCtl,
	}
}

func (handler *Handler) getProfile(c *gin.Context) {
	token := c.Request.Header["Authorization"]
	if len(token) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	profile, err := handler.scCtl.GetProfile(id, token[0])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "process failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": profile})
}

func (handler *Handler) updateProfile(c *gin.Context) {
	token := c.Request.Header["Authorization"]
	if len(token) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	var bodyParam *models.Profile
	if err := c.ShouldBindWith(&bodyParam, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err = handler.scCtl.UpdateProfile(id, token[0], bodyParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "process failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": true})
}

func (handler *Handler) createScreech(c *gin.Context) {
	creatorid, err := strconv.ParseInt(c.Query("creatorid"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	var bodyParam *models.Screech
	if err := c.ShouldBindWith(&bodyParam, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err = handler.scCtl.CreateScreech(creatorid, bodyParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "process failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": true})
}

func (handler *Handler) getScreech(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	screech, err := handler.scCtl.GetScreech(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "process failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": screech})
}

func (handler *Handler) updateScreech(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	var bodyParam *models.Screech
	if err := c.ShouldBindWith(&bodyParam, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err = handler.scCtl.UpdateScreech(id, bodyParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "process failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": true})
}

func (handler *Handler) getScreechList(c *gin.Context) {
	order := c.Query("order")
	if strings.ToUpper(order) != "ASC" {
		order = "desc"
	}

	creatorid, err := strconv.ParseInt(c.Query("creatorid"), 10, 64)

	size, err := strconv.ParseInt(c.Query("size"), 10, 64)
	if size == 0 {
		size = 50
	} else if size > 500 {
		size = 500
	}

	screechs, err := handler.scCtl.GetScreechList(creatorid, size, order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "process failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": screechs})
}
