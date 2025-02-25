package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	//setup a request and response
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	//conver the home function to a http.Handler type
	handler := http.HandlerFunc(home)
	//pass the fake HTTP reques to the handler
	handler.ServeHTTP(rr, req)

	//check the status code
	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("got %v expected %v", status, http.StatusOK)
	}
	//check the response body
	expected := "Welcome to our newsletter"
	if rr.Body.String() != expected {
		t.Errorf("got %v expected %v", rr.Body.String(), expected)
	}
}

func TestAboutHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/about", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(about)
	handler.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("got %v expected %v", status, http.StatusOK)
	}

	expected := "This is the about page"
	if rr.Body.String() != expected {
		t.Errorf("got %v expected %v", rr.Body.String(), expected)
	}
}

func TestSignupHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/signup", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(signup)
	handler.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("got %v expected %v", status, http.StatusOK)
	}

	expected := "This is the signup page"
	if rr.Body.String() != expected {
		t.Errorf("got %v expected %v", rr.Body.String(), expected)
	}
}
