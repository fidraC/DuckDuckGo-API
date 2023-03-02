package utils

import (
	"io"
	url "net/url"
)

// URL encode form data
func Url_encode(formdata map[string]string) string {
	var form string
	for key, value := range formdata {
		form += key + "=" + url.QueryEscape(value) + "&"
	}
	return form
}

func StringToReadCloser(s string) *readCloser {
	return &readCloser{s: s}
}

type readCloser struct {
	s string
}

func (r *readCloser) Read(p []byte) (n int, err error) {
	n = copy(p, r.s)
	r.s = r.s[n:]
	if n == 0 {
		err = io.EOF
	}
	return
}

func (r *readCloser) Close() error {
	return nil
}
