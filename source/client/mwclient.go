package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

var apiURL = "https://www.dictionaryapi.com/api/v3/references/collegiate/json/%s?key=%s"

// Definition holds the relevant fields from the API response
// Only the fields we need are included
// See: https://dictionaryapi.com/products/api-collegiate-dictionary

type MWEntry struct {
	Meta struct {
		ID string `json:"id"`
	} `json:"meta"`
	Hwi struct {
		Hw  string `json:"hw"`
		Prs []struct {
			Mw string `json:"mw"`
		} `json:"prs"`
	} `json:"hwi"`
	Fl       string   `json:"fl"`
	Shortdef []string `json:"shortdef"`
}

type Definition struct {
	Word          string
	Pronunciation string
	PartOfSpeech  string
	Meaning       string
}

type SuggestionError struct {
	Suggestions []string
}

func (e *SuggestionError) Error() string {
	return "no definition found, but suggestions are available"
}

func extractSuggestions(body []byte) ([]string, error) {
	var suggestions []string
	if err := json.Unmarshal(body, &suggestions); err == nil && len(suggestions) > 0 {
		return suggestions, nil
	}
	return nil, errors.New("no suggestions found")
}

func Lookup(word string) (*Definition, error) {
	apiKey := os.Getenv("MERRIAM_WEBSTER_API_KEY")
	if apiKey == "" {
		return nil, errors.New("MERRIAM_WEBSTER_API_KEY environment variable not set")
	}

	url := fmt.Sprintf(apiURL, word, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API request failed with status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var entries []MWEntry
	if err := json.Unmarshal(body, &entries); err == nil && len(entries) > 0 && len(entries[0].Shortdef) > 0 {
		entry := entries[0]
		pron := ""
		if len(entry.Hwi.Prs) > 0 {
			pron = entry.Hwi.Prs[0].Mw
		}
		return &Definition{
			Word:          entry.Hwi.Hw,
			Pronunciation: pron,
			PartOfSpeech:  entry.Fl,
			Meaning:       entry.Shortdef[0],
		}, nil
	}

	suggestions, err := extractSuggestions(body)
	if err == nil {
		return nil, &SuggestionError{Suggestions: suggestions}
	}

	return nil, errors.New("no definition found")
}
