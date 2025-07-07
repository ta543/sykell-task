package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/ta543/task/backend/crawler"
	"github.com/ta543/task/backend/models"
)

const authToken = "secret-token"

var cancelMap = make(map[int64]context.CancelFunc)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != authToken {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}

func AddURL(c *gin.Context) {
	var req struct {
		Address string `json:"address"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	url := &models.URL{Address: req.Address, Status: models.StatusQueued}
	if err := models.CreateURL(url); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancelMap[url.ID] = cancel
	go processURL(ctx, url)

	c.JSON(http.StatusOK, url)
}

func processURL(ctx context.Context, u *models.URL) {
	u.Status = models.StatusProcessing
	models.UpdateURL(u)

	result, err := crawler.Crawl(ctx, u.Address)
	if err != nil {
		u.Status = models.StatusError
		models.UpdateURL(u)
		return
	}

	u.Title = result.Title
	u.HTMLVersion = result.HTMLVersion
	u.H1Count = result.H1Count
	u.H2Count = result.H2Count
	u.H3Count = result.H3Count
	u.InternalLinks = result.InternalLinks
	u.ExternalLinks = result.ExternalLinks
	u.BrokenLinks = result.BrokenLinksCount
	u.HasLoginForm = result.HasLoginForm
	u.Status = models.StatusDone
	models.UpdateURL(u)

	var list []models.BrokenLink
	for _, b := range result.BrokenLinks {
		list = append(list, models.BrokenLink{
			URLID:      u.ID,
			Link:       b.URL,
			StatusCode: b.StatusCode,
		})
	}
	if len(list) > 0 {
		models.AddBrokenLinks(u.ID, list)
	}
}

func ListURLs(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.DefaultQuery("search", "")
	status := c.DefaultQuery("status", "")
	html := c.DefaultQuery("html", "")

	offset := (page - 1) * limit
	list, err := models.GetURLs(limit, offset, search, status, html)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}

func GetURL(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	u, err := models.GetURL(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if u == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	if links, err := models.GetBrokenLinksForURL(id); err == nil {
		u.BrokenLinksDetail = links
	}
	c.JSON(http.StatusOK, u)
}

func DeleteURL(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := models.DeleteURL(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func RestartURL(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	u, err := models.GetURL(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if u == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	u.Status = models.StatusQueued
	models.UpdateURL(u)
	ctx, cancel := context.WithCancel(context.Background())
	cancelMap[id] = cancel
	go processURL(ctx, u)
	c.JSON(http.StatusOK, u)
}

func StopURL(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if cancel, ok := cancelMap[id]; ok {
		cancel()
		delete(cancelMap, id)
	}
	u, _ := models.GetURL(id)
	if u != nil {
		u.Status = models.StatusStopped
		models.UpdateURL(u)
	}
	c.Status(http.StatusOK)
}
