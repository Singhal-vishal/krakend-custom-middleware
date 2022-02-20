package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"

	"krakend-custom-middleware/internal/constants"
)

func InitApiKeyValidationMiddleware() *ApiKeyValidationMiddleware {
	return &ApiKeyValidationMiddleware{}
}

type ApiKeyValidationMiddleware struct{}

func (rakv *ApiKeyValidationMiddleware) Apply(c *gin.Context) {
	apiKey := c.GetHeader(constants.AUTHORIZATION_HEADER)
	if len(apiKey) == 0 {
		c.AbortWithStatus(401)
		return
	}
	if apiKey != constants.API_KEY {
		log.Println("Api Key does not match.")
		c.AbortWithStatus(401)
		return
	}
	log.Println("Api key matched, sending call to next middleware(in our case to downstream service)")
	c.Next()
}
