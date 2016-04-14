package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

var testVal = "angryMonkey"
var expectedResult = "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q=="

func Test_EncodedHash_returns_correct_result(t *testing.T) {
	hashedBase64 := EncodedHash([]byte(testVal))
	if hashedBase64 != expectedResult {
		t.Error("EncodedHash output is incorrect")
	}
}

func Test_hashStringhandler_returns_400_on_missing_password(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()
	r := &http.Request{
		Method: "GET",
		URL:    &url.URL{Path: "/"},
		Form:   url.Values{},
	}
	handler := hashStringhandler{}
	handler.ServeHTTP(w, r)
	expectedValue := 400
	if w.Code != expectedValue {
		t.Errorf("\nexpected: %d\ngot: %d", expectedValue, w.Code)
	}
}

func Test_hashStringhandler_returns_normal_on_response(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()
	r := &http.Request{
		Method: "GET",
		URL:    &url.URL{Path: "/"},
		Form:   url.Values{},
	}
	r.Form.Set("password", testVal)
	handler := hashStringhandler{}
	handler.ServeHTTP(w, r)
	expectedValue := 200
	if w.Code != expectedValue {
		t.Errorf("\nexpected: %d\ngot: %d", expectedValue, w.Code)
	}
	result := w.Body.String()
	if result != expectedResult {
		t.Errorf("\nexpected: %s\ngot: %s", expectedResult, result)
	}
}

func Test_hashStringhandler_returns_after_5_seconds(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()
	r := &http.Request{
		Method: "GET",
		URL:    &url.URL{Path: "/"},
		Form:   url.Values{},
	}
	r.Form.Set("password", testVal)
	handler := hashStringhandler{}
	defer TimeTrack(t, time.Now(), "delay test")
	handler.ServeHTTP(w, r)
}

func TimeTrack(t *testing.T, start time.Time, name string) {
	elapsed := time.Since(start)
	if elapsed < 5*time.Second {
		t.Error("socket open for less then 5 seconds")
	}
}
