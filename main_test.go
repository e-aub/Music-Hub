package main

import (
	"musicHub/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetch(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/artist?id=1", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.Artist)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("not 200")
	}

}
