package api // import "github.com/eriol/wp24-deities/api"

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"

	"github.com/eriol/wp24-deities/cfg"
)

var oauthsrv *server.Server

func Serve() {

	manager := manage.NewDefaultManager()
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	clientStore := store.NewClientStore()
	clientId := cfg.GetClientId()
	clientStore.Set(clientId, &models.Client{
		ID:     clientId,
		Secret: cfg.GetClientSecret(),
		Domain: cfg.GetDomain(),
	})
	manager.MapClientStorage(clientStore)
	srv := server.NewDefaultServer(manager)
	oauthsrv = srv
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)
	srv.SetUserAuthorizationHandler(useraAuthorizationHandler)
	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})
	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	router := http.NewServeMux()

	router.HandleFunc("GET /authorize", authorize)
	router.HandleFunc("GET /token", token)

	router.HandleFunc("GET /", info)
	router.HandleFunc("GET /deities", getDeities)
	router.HandleFunc("GET /deities/{slug}", getDeity)
	router.HandleFunc("GET /deities/{slug}/influence", getDeityInfluence)
	router.HandleFunc("GET /random", random)
	router.HandleFunc("OPTIONS /random", cors)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func toJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if status != http.StatusOK {
		w.WriteHeader(status)
	}

	if err := json.NewEncoder(w).Encode(value); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
