// main_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"
)

func TestDefoultHundlr(t *testing.T) {
	// Test for correct path
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	testing.Init()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DefoultHundlr)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Test for incorrect path
	req, err = http.NewRequest("GET", "/wrong", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func init() {
	var err error
	Tp1, err = template.ParseFiles("templates/pageDefoult.html")
	if err != nil {
		panic(err)
	}
}

func TestAsciiHundler(t *testing.T) {
	// Test for correct POST request
	postData := strings.NewReader("font=standard&inputText=Hello")
	req, err := http.NewRequest("POST", "/ascii-art", postData)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AsciiHundler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Test for incorrect method (GET instead of POST)
	req, err = http.NewRequest("GET", "/ascii-art", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}
}
