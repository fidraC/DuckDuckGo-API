package duckduckgo

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/acheong08/DuckDuckGo-API/internal/types"
	"github.com/acheong08/DuckDuckGo-API/internal/utils"
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
