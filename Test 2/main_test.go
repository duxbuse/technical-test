package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVersionHandler(t *testing.T) {

	//Test Data
	VersionNumber = `0.0.0-rc.0`
	LastCommitSha = `abc123`
	Message = VersionMessage{VersionNumber, LastCommitSha, "pre-interview technical test"}

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/version", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(VersionHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code\n got: %v\nwant: %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "{\"Version\":\"" + VersionNumber +
		"\",\"LastCommitSha\":\"" + LastCommitSha +
		"\",\"Description\":\"pre-interview technical test\"}\n"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body.\n got: %v\nwant: %v",
			rr.Body.String(), expected)
	}
}
