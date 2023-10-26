// auth_middleware.go

package middleware

import (
	"go-server/utils"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Claims represents the JWT claims.
type Claims struct {
	UserId    int    `json:"user_id"`
	UserEmail string `json:"user_email"`
	UserRole  string `json:"user_role"`
	jwt.StandardClaims
}

// AuthMiddleware checks JWT tokens from cookies and enforces user roles.
func AuthMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the JWT token from the cookie
		cookie, err := c.Request.Cookie("jwt-token")
		if err != nil {
			// Token not found, redirect to login page
			c.Redirect(http.StatusSeeOther, "http://localhost:8080")
			c.Abort()
			return
		}

		token := cookie.Value

		claims, valid := utils.VerifyJWTToken(token)
		if !valid {
			// Token is invalid or expired, redirect to login page
			c.Redirect(http.StatusSeeOther, "http://localhost:8080/")
			c.Abort()
			return
		}

		// Check if the user has the required role
		hasRequiredRole := false
		userRole := claims.UserRole // Access the user role from the claims

		for _, role := range roles {
			if userRole == role {
				hasRequiredRole = true
				break
			}
		}

		if !hasRequiredRole {
			c.JSON(http.StatusForbidden, gin.H{"message": "Access Forbidden"})
			c.Abort()
			return
		}

		c.Next()
	}
}
