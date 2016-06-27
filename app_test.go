package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCanAddHandler(t *testing.T) {
	app := NewApp()

	app.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request, a *App) {
		fmt.Fprintln(w, "Called")
	})

	resp := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	app.Mux.ServeHTTP(resp, req)

	b, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fail()
	}

	if !strings.Contains(string(b), "Called") {
		t.Error("Handler was not called.")
	}
}
