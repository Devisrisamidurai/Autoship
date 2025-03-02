package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetTasks(t *testing.T) {
	os.Setenv("TEST_MODE", "true")
	initDB()
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/tasks", getTasks)

	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
