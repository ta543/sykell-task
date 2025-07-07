package handlers

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/require"
)

func TestAuthMiddleware(t *testing.T) {
    gin.SetMode(gin.TestMode)
    r := gin.New()
    r.Use(AuthMiddleware())
    r.GET("/protected", func(c *gin.Context) { c.Status(http.StatusOK) })

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/protected", nil)
    r.ServeHTTP(w, req)
    require.Equal(t, http.StatusUnauthorized, w.Code)

    w = httptest.NewRecorder()
    req, _ = http.NewRequest("GET", "/protected", nil)
    req.Header.Set("Authorization", authToken)
    r.ServeHTTP(w, req)
    require.Equal(t, http.StatusOK, w.Code)
}

