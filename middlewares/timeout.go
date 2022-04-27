package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
)

func Filter(c *gin.Context) func(c *gin.Context) {
	t := time.Second * 10
	return func(c *gin.Context) {
		finish := make(chan struct{})

		go func() {
			c.Next()
			finish <- struct{}{}
		}()

		select {
		case <-time.After(t):
			c.JSON(504, "timeout")
			c.Abort()
			return
		case <-finish:

		}
	}
}
