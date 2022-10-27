package controllers

import "github.com/gin-gonic/gin"

type UserController struct{}

func (a UserController) Add(c *gin.Context) {
	c.String(200, "ADD")
}
func (a UserController) Search(c *gin.Context) {
	c.String(200, "Search")
}
