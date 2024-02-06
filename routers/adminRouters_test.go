package routers

import (
	"HMS/payment/controllers/admin"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAdminRoutersInit(t *testing.T) {
	r := gin.Default()

	r.GET("/pay", admin.BasePaymentControl{}.Pay)

	ts := httptest.NewServer(r)

	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/pay", ts.URL))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]

	if !ok {
		t.Fatal("Expected Content-Type header to be set")
	}

	if val[0] != "text/plain; charset=utf-8" {
		t.Fatalf("Expected \"text/plain;; charset=utf-8\", got %s", val[0])
	}
}
