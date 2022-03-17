package healthModule

import (
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

// TODO: e2e 테스트 환경 리서치 & 만들기

func TestHealthCheck(t *testing.T) {
	response, _ := CreateModule().Test(httptest.NewRequest("GET", "/", nil))

	if response.StatusCode != fiber.StatusNoContent {
		t.Fatalf("HealthCheck statusCode must be 204")
	}
}

func TestGetVersion(t *testing.T) {
	response, _ := CreateModule().Test(httptest.NewRequest("GET", "/version", nil))

	body, _ := ioutil.ReadAll(response.Body)

	if string(body) != "go1.18" {
		t.Fatalf("Current version is go1.18")
	}
}
