package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"purpleschool-go/advanced/internal/auth"
	"testing"
)

func TestLoginSuccess(t *testing.T) {
	ts := httptest.NewServer(App("test"))

	defer ts.Close()

	loginPayload, _ := json.Marshal(&auth.LoginRequest{
		Email:    "a2@a.ru",
		Password: "123",
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(loginPayload))

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected %d got %d", http.StatusOK, res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		t.Fatal(err)
	}

	var data auth.LoginResponse

	err = json.Unmarshal(body, &data)

	if err != nil {
		t.Fatal(err)
	}

	if data.Token == "" {
		t.Fatalf("Expected jwt got %s", data.Token)
	}
}

func TestLoginFail(t *testing.T) {
	ts := httptest.NewServer(App("test"))

	defer ts.Close()

	loginPayload, _ := json.Marshal(&auth.LoginRequest{
		Email:    "a2@a.ru",
		Password: "123456789",
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(loginPayload))

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("Expected %d got %d", http.StatusUnauthorized, res.StatusCode)
	}
}
