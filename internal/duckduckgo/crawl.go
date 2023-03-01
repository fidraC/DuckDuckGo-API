package duckduckgo

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/acheong08/DuckDuckGo-API/internal/types"
	"github.com/acheong08/DuckDuckGo-API/internal/utils"
	"github.com/anaskhan96/soup"
	_ "github.com/anaskhan96/soup"
)

func Get_html(search types.Search) (string, error) {
	var base_url string = "html.duckduckgo.com"
	// POST form data
	var formdata = map[string]string{
		"q":  search.Query,
		"df": search.TimeRange,
		"kl": search.Region,
	}
	// URL encode form data
	var form string = utils.Url_encode(formdata)
	// Create POST request
	var request = http.Request{
		Method: "POST",
		URL: &url.URL{
			Host:   base_url,
			Path:   "/html/",
			Scheme: "https",
		},
		Header: map[string][]string{
			"Content-Type": {"application/x-www-form-urlencoded"},
		},
		Body: utils.StringToReadCloser(form),
	}
	// Send POST request
	var client = http.Client{}
	var response, err = client.Do(&request)
	if err != nil {
		return "", err
	}
	if response.StatusCode != 200 {
		return "", errors.New("Status code: " + strconv.Itoa(response.StatusCode))
	}
	// Read response body
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	// Close response body
	err = response.Body.Close()
	if err != nil {
		return "", err
	}
	return string(bodyBytes), nil
}

func Parse_html(html string) ([]types.Result, error) {
	// Results is an array of Result structs
	var final_results []types.Result = []types.Result{}
	// Parse
	doc := soup.HTMLParse(html)
	// Find each result__body
	result_bodies := doc.FindAll("div", "class", "result__body")
	// Loop through each result__body
	for _, item := range result_bodies {
		// Get text of result__title
		var title string = item.Find("a", "class", "result__a").Text()
		// Get href of result__a
		var link string = item.Find("a", "class", "result__a").Attrs()["href"]
		// Get text of result__snippet
		var snippet string = item.Find("div", "class", "result__snippet").Text()
		// Append to final_results
		final_results = append(final_results, types.Result{
			Title:   title,
			Link:    link,
			Snippet: snippet,
		})
	}
	return final_results, nil
}
