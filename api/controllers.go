package api // import "github.com/eriol/wp24-deities/api"

import (
	"database/sql"
	"log"
	"net/http"
	"strings"

	"github.com/eriol/wp24-deities/database"
)

type ApiInfo struct {
	Description string `json:"description"`
	Version     string `json:"version"`
}

type ApiError struct {
	Error string `json:"error"`
}

func preflight(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Method", "GET, OPTIONS")
	(*w).Header().Set("Access-Control-Max-Age", "86400")
}

// Return API description.
// This endpoint is the root of the API.
func info(w http.ResponseWriter, r *http.Request) {
	preflight(&w)

	// The "/" pattern matches everything, so check if we are at the
	// root and return a 403 otherwise (we blame the client for endpoints that
	// don't exist!).
	//
	// It's not possible to specify a custom NotFound(), because in
	// https://golang.org/src/pkg/net/http/server.go NotFoundHandler()
	// returns a hardcoded function called NotFound(). So we need to do this to
	// use JSON instead.
	if r.URL.Path != "/" {
		toJSON(w, http.StatusForbidden, ApiError{Error: "Forbidden"})
		return
	}

	api := ApiInfo{
		Description: "A simple open REST API for deities!",
		Version:     "0.1",
	}
	toJSON(w, http.StatusOK, api)
}

// Return an array with all the deities.
func getDeities(w http.ResponseWriter, r *http.Request) {
	preflight(&w)

	deities, err := database.GetDeities()

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	toJSON(w, http.StatusOK, deities)
}

// Return the specified (in the path) deity.
func getDeity(w http.ResponseWriter, r *http.Request) {
	preflight(&w)

	slug := strings.TrimSpace(r.PathValue("slug"))
	if slug == "" {
		toJSON(w, http.StatusBadRequest, ApiError{Error: "No deity slug provided"})
		return
	}

	deity, err := database.GetDeity(slug)

	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			toJSON(w, http.StatusNotFound, ApiError{Error: "No deity found"})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	toJSON(w, http.StatusOK, deity)
}

// Return the specified (in the path) deity influence.
func getDeityInfluence(w http.ResponseWriter, r *http.Request) {
	preflight(&w)

	slug := strings.TrimSpace(r.PathValue("slug"))
	if slug == "" {
		toJSON(w, http.StatusBadRequest, ApiError{Error: "No deity slug provided"})
		return
	}

	influence, err := database.GetDeityInfluence(slug)

	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			toJSON(w, http.StatusNotFound, ApiError{Error: "No deity influence found"})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	toJSON(w, http.StatusOK, influence)
}

// Return a random deity.
func random(w http.ResponseWriter, r *http.Request) {
	preflight(&w)

	deity, err := database.GetRandomDeity()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	toJSON(w, http.StatusOK, deity)
}
