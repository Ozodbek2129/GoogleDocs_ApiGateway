package middleware

import (
	auth "api_gateway/api/token"
	"errors"
	"fmt"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type casbinPermission struct {
	enforcer *casbin.Enforcer
}

func Check(c *gin.Context) {

	accessToken := c.GetHeader("Authorization")
	fmt.Println(1)
	if accessToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization is required",
		})
		return
	}
	fmt.Println(2)

	_, err := auth.ValidateAccessToken(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token provided",
		})
		return
	}
	fmt.Println(3)
	c.Next()
}

func (casb *casbinPermission) GetRole(c *gin.Context) (string, int) {
	token := c.GetHeader("Authorization")
	fmt.Println(4)
	if token == "" {
		return "unauthorized", http.StatusUnauthorized
	}
	fmt.Println(5)

	_, _, _, role, err := auth.GetUserInfoFromAccessToken(token)
	if err != nil {
		return "error while reding role", 500
	}
	fmt.Println(6)

	return role, 0
}

func (casb *casbinPermission) CheckPermission(c *gin.Context) (bool, error) {

	act := c.Request.Method
	sub, status := casb.GetRole(c)
	fmt.Println(7)

	if status != 0 {
		return false, errors.New("error in get role")
	}
	obj := c.FullPath()
	fmt.Println(8)

	ok, err := casb.enforcer.Enforce(sub, obj, act)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Internal server error",
		})
		c.Abort()
		return false, err
	}
	fmt.Println(9)

	return ok, nil
}

func CheckPermissionMiddleware(enf *casbin.Enforcer) gin.HandlerFunc {
	casbHandler := &casbinPermission{
		enforcer: enf,
	}

	return func(c *gin.Context) {
		fmt.Println(10)

		result, err := casbHandler.CheckPermission(c)

		if err != nil {
			c.AbortWithError(500, err)
		}
		if !result {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "Forbidden",
			})
		}

		c.Next()
	}
}
