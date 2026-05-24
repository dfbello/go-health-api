package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthEndpointHappyPath (t *testing.T) {
	// create the request
	req := httptest.NewRequest("GET", "/health", nil)
	// create response recorder
	rec := httptest.NewRecorder()

	// call the Handler
	healthHandler(rec, req)	

	res := rec.Result()

	// Assert
	expectedStatusCode := 200
	gotStatusCode := res.StatusCode
	if expectedStatusCode != gotStatusCode {
		t.Errorf("Expected %v but got %v", expectedStatusCode, gotStatusCode)
	}

	expectedContentType := "application/json; charset=utf-8"
	gotContentType := res.Header.Get("Content-Type")
	if expectedContentType != gotContentType {
		t.Errorf("Expected %v but got %v", expectedContentType, gotContentType)
	}

	expectedBody := "{\"status\": \"ok\"}"
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Could not read Body: %v", err)
	}

	gotBody := string(body)
	if expectedBody != gotBody {
		t.Errorf("Expected %v but got %v", expectedBody, gotBody)
	}

}

func TestHealthEndpointWrongMethod(t *testing.T) {
	
	// for this test serve mux is necessary cause it's the one to enforce the valid http methods
	sMux := http.NewServeMux()
	sMux.HandleFunc("GET /health", healthHandler)

	// create the request
	req := httptest.NewRequest("GET", "/health", nil)
	// create response recorder
	rec := httptest.NewRecorder()

	sMux.ServeHTTP(rec, req)

	res := rec.Result()

	//Assert
	expectedStatusCode := http.StatusMethodNotAllowed
	gotStatusCode := res.StatusCode
	if expectedStatusCode != gotStatusCode {
		t.Errorf("Expected %v but got %v", expectedStatusCode, gotStatusCode)
	}

}
