package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/3n0ugh/GoCrudBook/cmd/web/config"
	"github.com/3n0ugh/GoCrudBook/cmd/web/router"
	"github.com/3n0ugh/GoCrudBook/pkg/models"
	"github.com/3n0ugh/GoCrudBook/pkg/models/mysql"
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

func NewTestDB(t *testing.T) (*sql.DB, func()) {
	db, err := sql.Open("mysql", "root@/test_library?multiStatements=true")
	if err != nil {
		t.Fatal(err)
	}

	script, err := os.ReadFile("../../pkg/models/mysql/testdata/setup.sql")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec(string(script))
	if err != nil {
		t.Fatal(err)
	}

	return db, func() {
		script, err := os.ReadFile("../../pkg/models/mysql/testdata/teardown.sql")
		if err != nil {
			t.Fatal(err)
		}
		_, err = db.Exec(string(script))
		if err != nil {
			t.Fatal(err)
		}

		db.Close()
	}
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
		{"Non-existent ID", "/book/id?id=2", http.StatusNotFound, nil},
		{"Negative ID", "/book/id?id=-1", http.StatusBadRequest, nil},
		{"Decimal ID", "/book/id?id=1.23", http.StatusBadRequest, nil},
		{"String ID", "/book/id?id=foo", http.StatusBadRequest, nil},
		{"Empty ID", "/book/id?id=", http.StatusBadRequest, nil},
		{"Trailing slash", "/book/id?id=/", http.StatusBadRequest, nil},
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

			if tt.name != "Valid ID" && fmt.Sprintf("%s", string(body)) != http.StatusText(tt.wantCode)+"\n" {
				t.Errorf("want body to contain %q, %q", http.StatusText(tt.wantCode)+"\n", string(body))
			} else if tt.name == "Valid ID" && fmt.Sprintf("%v", string(body)) != fmt.Sprintf("%v", tt.wantBody) {
				t.Errorf("want body to contain %v but got %v", tt.wantBody, string(body))
			}
		})
	}

}

func TestGetBokkAll(t *testing.T) {

	bookmodels := []models.Book{
		{
			ISBN:        1933988673,
			BookName:    "Unlocking Android: A Developer Guide",
			Author:      "Charlie Collins",
			PageCount:   416,
			BookCount:   1,
			BorrowTimes: 0,
		},
		{
			ISBN:        1933988746,
			BookName:    "Flex 3 in Action",
			Author:      "Tariq Ahmed with Jon Hirschi",
			PageCount:   576,
			BookCount:   1,
			BorrowTimes: 0,
		},
	}

	test := struct {
		name     string
		urlPath  string
		wantCode int
		wantBody []models.Book
	}{"GetBookAll", "/book", http.StatusOK, bookmodels}

	app := newTestApplication(t)
	ts := newTestServer(t, router.SetRoutes(app))
	defer ts.Close()

	t.Run(test.name, func(t *testing.T) {
		db, teardown := NewTestDB(t)
		defer teardown()

		m := mysql.BookModel{
			DB: db,
		}

		books, _ := m.GetAll()

		var booklist []models.Book
		for _, book := range books {
			booklist = append(booklist, *book)
		}

		if !reflect.DeepEqual(booklist, test.wantBody) {
			t.Errorf("want \n%v; got \n%v", test.wantBody, booklist)
		}
	})
}
