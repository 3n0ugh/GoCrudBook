package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/3n0ugh/GoCrudBook/cmd/web/config"
	"github.com/3n0ugh/GoCrudBook/cmd/web/router"
	"github.com/3n0ugh/GoCrudBook/pkg/models"
)

func newTestApplication(t *testing.T) *config.Application {
	return &config.Application{
		ErrorLog: log.New(io.Discard, "", 0),
		InfoLog:  log.New(io.Discard, "", 0),
	}
}

type testServer struct {
	*httptest.Server
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewServer(h)
	return &testServer{ts}
}

func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, []byte) {
	rs, err := ts.Client().Get("http://localhost:5000" + urlPath)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	return rs.StatusCode, rs.Header, body
}

func TestHome(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, router.SetRoutes(app))
	defer ts.Close()

	code, _, body := ts.get(t, "/")

	if code != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, code)
	}

	if string(body) != "Home Page" {
		t.Errorf("want body to equal %q", "Home Page")
	}
}

func TestGetBookByID(t *testing.T) {
	bookmodel := &models.Book{
		ISBN:        1932394160,
		BookName:    "Show Must Go On",
		Author:      "Show Must Go On",
		PageCount:   311,
		BookCount:   2,
		BorrowTimes: 7,
	}

	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody *models.Book
	}{
		{"Valid ID", "/book/id?id=1932394160", http.StatusOK, bookmodel},
	}

	app := newTestApplication(t)
	ts := newTestServer(t, router.SetRoutes(app))
	defer ts.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.get(t, tt.urlPath)

			if code != tt.wantCode {
				t.Errorf("want %d; got %d", tt.wantCode, code)
			}

			if fmt.Sprintf("%v", string(body)) != fmt.Sprintf("%v", tt.wantBody) {
				t.Errorf("want body to contain %v but got %v", tt.wantBody, string(body))
			}
		})
	}

}
