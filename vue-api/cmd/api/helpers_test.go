package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_readJSON(t *testing.T) {
	// create sample JSON
	sampleJSON := map[string]interface{}{
		"foo": "bar",
	}

	body, _ := json.Marshal(sampleJSON)

	// declare a variable that we can read into
	var decodedJSON struct {
		FOO string `json:"foo"`
	}

	// create a request
	req, err := http.NewRequest("POST", "/", bytes.NewReader(body))
	if err != nil {
		t.Log(err)
	}

	// create a test response recorder
	rr := httptest.NewRecorder()
	defer req.Body.Close()

	// call readJSON
	err = testApp.readJSON(rr, req, &decodedJSON)
	if err != nil {
		t.Error("failed to decode json", err)
	}
}

func Test_writeJSON(t *testing.T) {
	rr := httptest.NewRecorder()
	payload := jsonResponse{
		Error:   false,
		Message: "foo",
	}

	headers := make(http.Header)
	headers.Add("FOO", "BAR")
	err := testApp.writeJSON(rr, http.StatusOK, payload, headers)
	if err != nil {
		t.Errorf("failed to write JSON: %v", err)
	}

	testApp.environment = "production"
	err = testApp.writeJSON(rr, http.StatusOK, payload, headers)
	if err != nil {
		t.Errorf("failed to write JSON in production environment: %v", err)
	}

	testApp.environment = "development"
}

func Test_errorJSON(t *testing.T) {
	rr := httptest.NewRecorder()
	err := testApp.errorJSON(rr, errors.New("new error."))
	if err != nil {
		t.Error(err)
	}

	testJSONPayload(t, rr)

	errSlice := []string{
		"(SQLSTATE 23505)",
		"(SQLSTATE 22001)",
		"(SQLSTATE 23503)",
	}

	for _, x := range errSlice {
		customErr := testApp.errorJSON(rr, errors.New(x), http.StatusUnauthorized)
		if customErr != nil {
			t.Error(customErr)
		}
		testJSONPayload(t, rr)
	}
}

func testJSONPayload(t *testing.T, rr *httptest.ResponseRecorder) {
	var requestPayload jsonResponse
	decoder := json.NewDecoder(rr.Body)
	err := decoder.Decode(&requestPayload)
	if err != nil {
		t.Error("received error when decoding errorJSON payload: ", err)
	}

	if !requestPayload.Error {
		t.Error("error set to false in response from errorJSON, and it should be set to true.")
	}

}
