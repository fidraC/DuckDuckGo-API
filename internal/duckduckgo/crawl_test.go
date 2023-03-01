package duckduckgo_test

import (
	"testing"

	"github.com/acheong08/DuckDuckGo-API/internal/duckduckgo"
	types "github.com/acheong08/DuckDuckGo-API/internal/types"
)

func TestGet_html(t *testing.T) {
	search := types.Search{
		Query:     "test",
		Region:    "",
		TimeRange: "",
	}
	response, err := duckduckgo.Get_html(search)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if response == "" {
		t.Errorf("Response is empty")
	}
	t.Log(response)
}
