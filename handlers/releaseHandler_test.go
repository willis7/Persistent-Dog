package handlers

import (
	"testing"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strings"
	"fmt"
)


// User Story - User should be able to view list of Release entity
func TestGetReleases(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/releases", GetReleases).Methods("GET")

	req, err := http.NewRequest("GET", "/releases", nil)

	if err != nil {
		t.Error(err)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != 200 {
			t.Errorf("HTTP Status expected: 200, got: %d", w.Code)
		}
	}
}

// User Story - User should be able to create a Release entity
func TestCreateRelease(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/releases", CreateRelease).Methods("POST")
	releaseJson := `{"name": "name", "description": "description", "version": "1.0"}`

	req, err := http.NewRequest("POST", "/releases", strings.NewReader(releaseJson))
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != 201 {
		t.Errorf("HTTP Status expected: 201, got: %d", w.Code)
	}
}

//User Story - The Name of a Release entity should be unique
func TestUniqueName(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/releases", CreateRelease).Methods("POST")
	releaseJson := `{"name": "name", "description": "description", "version": "1.0"}`

	req, err := http.NewRequest("POST", "/releases", strings.NewReader(releaseJson))
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != 400 {
		t.Error("Bad Request expected, got: %d", w.Code)
	}
}

func TestGetReleasesClient(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/releases", GetReleases).Methods("GET")
	server := httptest.NewServer(r)
	defer server.Close()

	releasesUrl := fmt.Sprintf("%s/releases", server.URL)
	request, err := http.NewRequest("GET", releasesUrl, nil)

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", res.StatusCode)
	}
}

func TestCreateReleaseClient(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/releases", CreateRelease).Methods("POST")
	server := httptest.NewServer(r)
	defer server.Close()

	releasesUrl := fmt.Sprintf("%s/releases", server.URL)
	fmt.Println(releasesUrl)
	releaseJson := `{"name": "another name", "description": "description", "version": "1.0"}`
	request, err := http.NewRequest("POST", releasesUrl, strings.NewReader(releaseJson))

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 201 {
		t.Errorf("HTTP Status expected: 201, got: %d", res.StatusCode)
	}
}
