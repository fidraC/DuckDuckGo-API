package duckduckgo_test

import (
	"testing"

	"github.com/acheong08/DuckDuckGo-API/duckduckgo"
	types "github.com/acheong08/DuckDuckGo-API/types"
)

func TestGet_html(t *testing.T) {
	search := types.Search{
		Query:     "test",
		Region:    "",
		TimeRange: "",
	}
	response, err := duckduckgo.Get_results(search)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	t.Log(response)
}
