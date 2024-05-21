package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestMainApp(t *testing.T) {
	t.Run("Test hostname", func(t *testing.T) {
		for i := 0; i < 20; i++ {
			req, err := http.NewRequest("GET", "http://localhost:8080/", nil)
			if err != nil {
				t.Fatal(err)
			}
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				t.Fatal(err)
			}

			if i > 9 {
				assert.Equal(t, http.StatusTooManyRequests, resp.StatusCode)
			} else {
				assert.Equal(t, http.StatusOK, resp.StatusCode)
			}
		}
	})

	t.Run("Test with token", func(t *testing.T) {
		for i := 0; i < 200; i++ {
			req, err := http.NewRequest("GET", "http://localhost:8080/", nil)
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("X-API-KEY", "AAAAA")
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				t.Fatal(err)
			}

			if i > 99 {
				assert.Equal(t, http.StatusTooManyRequests, resp.StatusCode)
			} else {
				assert.Equal(t, http.StatusOK, resp.StatusCode)
			}
		}
	})
}
