package main

import (
	"context"

	"github.com/bisal-dd/go-backend-boilerplate/helper/contextUtil"
	"github.com/gin-gonic/gin"
)


func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// tokenString, err := jwtUtil.TokenFromRequest(c.Request)
		// if err != nil {
		// 	c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		// 	return
		// }
		// token, err := jwtUtil.ParseToken(tokenString)
		// if err != nil {
		// 	c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		// 	return
		// }
		// claims, err := jwtUtil.ClaimsFromToken(token)
		// if err != nil {
		// 	c.AbortWithStatusJSON(401, gin.H{"error": "invalid token"})
		// 	return
		// }

		// userID, ok := claims["user_id"].(string)
		// if !ok {
		// 	c.AbortWithStatusJSON(401, gin.H{"error": "user_id not found in token"})
		// 	return
		// }
		// ctx := contextUtil.SetContextValue(c.Request.Context(), contextUtil.UserIDKey, userID)
		// ctx = contextUtil.SetContextValue(ctx, contextUtil.GinContextKey, c)
		// c.Request = c.Request.WithContext(ctx)
		// c.Next()

		ctx := context.WithValue(c.Request.Context(), contextUtil.GinContextKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
