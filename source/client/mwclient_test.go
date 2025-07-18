package client

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func mockServer(response string, statusCode int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		io.WriteString(w, response)
	}))
}

func TestLookup_MissingAPIKey(t *testing.T) {
	os.Unsetenv("MERRIAM_WEBSTER_API_KEY")

	_, err := Lookup("exercise")
	if err == nil || !strings.Contains(err.Error(), "environment variable not set") {
		t.Errorf("expected missing API key error, got %v", err)
	}
}

func TestLookup_Non200Response(t *testing.T) {
	ts := mockServer("Internal Server Error", 500)
	defer ts.Close()

	os.Setenv("MERRIAM_WEBSTER_API_KEY", "test")
	apiURL = ts.URL + "/%s?key=%s"

	_, err := Lookup("exercise")
	if err == nil || !strings.Contains(err.Error(), "status 500") {
		t.Errorf("expected 500 status error, got %v", err)
	}
}

func TestLookup_ValidDefinition(t *testing.T) {
	response := `[{"meta": {"id": "exercise"}, "hwi": {"hw": "exercise", "prs": [{"mw": "ˈek-sər-ˌsīz"}]}, "fl": "noun", "shortdef": ["a physical activity"]}]`
	ts := mockServer(response, 200)
	defer ts.Close()

	os.Setenv("MERRIAM_WEBSTER_API_KEY", "test")
	apiURL = ts.URL + "/%s?key=%s"

	def, err := Lookup("exercise")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if def.Word != "exercise" || def.Meaning != "a physical activity" {
		t.Errorf("unexpected definition: %+v", def)
	}
}

func TestLookup_WithSuggestions(t *testing.T) {
	response := `["exorcise", "exert", "executive"]`
	ts := mockServer(response, 200)
	defer ts.Close()

	os.Setenv("MERRIAM_WEBSTER_API_KEY", "test")
	apiURL = ts.URL + "/%s?key=%s"

	_, err := Lookup("exrsize")
	if err == nil {
		t.Fatalf("expected suggestion error")
	}
	if se, ok := err.(*SuggestionError); !ok || len(se.Suggestions) == 0 {
		t.Errorf("expected SuggestionError, got %v", err)
	}
}

func TestLookup_MalformedJSON(t *testing.T) {
	response := `{{bad json]]`
	ts := mockServer(response, 200)
	defer ts.Close()

	os.Setenv("MERRIAM_WEBSTER_API_KEY", "test")
	apiURL = ts.URL + "/%s?key=%s"

	_, err := Lookup("badword")
	if err == nil || !strings.Contains(err.Error(), "no definition found") {
		t.Errorf("expected JSON unmarshal fallback error, got %v", err)
	}
}
