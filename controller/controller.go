package controller

import "github.com/gin-gonic/gin"

// Controller : holds router for route
type Controller struct {
	BaseURL string
	Router  *gin.RouterGroup
}

func (c *Controller) SetBaseURL(url string) {
	c.BaseURL = url
}

func (c *Controller) createURL(pattern string) string {
	return c.BaseURL + pattern
}
