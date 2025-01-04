package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/users", GetUsers)
	router.GET("/roles", GetRoles)
	router.Run()
}

func GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"UserID":   "001",
		"UserName": "Miguel",
	})
}

func GetRoles(c *gin.Context) {
	c.JSON(200, gin.H{
		"RoleID":   "100",
		"RoleName": "Admin",
	})
}
