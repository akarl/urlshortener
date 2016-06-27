package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func addURL(app *App, code, url string) {
	conn := app.URLStorage.RedisPool.Get()
	defer conn.Close()
	conn.Do("SET", code, url)
}

func TestGotoURL(t *testing.T) {
	app := NewApp()
	target := "http://example.com"
	addURL(app, "thecode", target)
	setupHandlers(app)

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/r/thecode", nil)
	app.Mux.ServeHTTP(resp, req)

	if resp.Code != http.StatusFound {
		t.Errorf("Expected redirect found %d", resp.Code)
	}

	if resp.Header().Get("Location") != target {
		t.Errorf("Expected redirect to %s.", target)

	}
}

func TestGETAddURL(t *testing.T) {
	app := NewApp()
	setupHandlers(app)

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/add", nil)
	app.Mux.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.Code)
	}
}

func TestPOSTAddURL(t *testing.T) {
	app := NewApp()
	setupHandlers(app)

	form := url.Values{
		"url": []string{"example.com"},
	}

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/add", strings.NewReader(form.Encode()))

	app.Mux.ServeHTTP(resp, req)

	if l := resp.Header().Get("Location"); !strings.HasPrefix(l, "/view?c=") {
		t.Errorf("Invalid redirect: %s", l)
	}
}

func TestViewURL(t *testing.T) {
	app := NewApp()
	setupHandlers(app)
	addURL(app, "thecode", "http://example.com")

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/view?c=thecode", nil)

	app.Mux.ServeHTTP(resp, req)
	body := resp.Body.String()
	if !strings.Contains(body, "/r/thecode") {
		t.Error(body)
	}
}
