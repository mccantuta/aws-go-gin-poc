package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func main() {
	isLambdaEnable := os.Getenv("LAMBDA_ENABLED")
	router := gin.Default()
	router.GET("/users", GetUsers)
	router.GET("/roles", GetRoles)
	if isLambdaEnable == "false" {
		router.Run()
	} else {
		ginLambda = ginadapter.New(router)
		lambda.Start(Handler)
	}
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
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
