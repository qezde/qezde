package http

import (
	"bytes"
	"gateway/internal/config"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type ProxyHandler struct{}

func NewProxyHandler() *ProxyHandler {
	return &ProxyHandler{}
}

func (h *ProxyHandler) Routes(routerGroup *gin.RouterGroup, config config.Config) {
	routerGroup.Any("/auth/*action", h.handleRequest(config.API.Auth))
}

func (h *ProxyHandler) handleRequest(targetURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Param("action")
		method := c.Request.Method
		query := c.Request.URL.RawQuery

		target := targetURL + path
		if query != "" {
			target += "?" + query
		}

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to read request body"})
			return
		}

		req, err := http.NewRequest(method, target, bytes.NewBuffer(body))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to create request"})
			return
		}

		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		for key, values := range c.Request.Header {
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "request to backend service failed"})
			return
		}
		defer resp.Body.Close()

		c.Writer.WriteHeader(resp.StatusCode)
		for key, values := range resp.Header {
			for _, value := range values {
				c.Writer.Header().Add(key, value)
			}
		}
		io.Copy(c.Writer, resp.Body)
	}
}