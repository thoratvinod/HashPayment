package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thoratvinod/HashPayment/handlers"
)

type Route struct {
	Path    string
	Methods []string
	Handler func(w http.ResponseWriter, r *http.Request)
}

var Routes = []Route{
	{
		Path:    "/ping",
		Methods: []string{http.MethodPost, http.MethodPost},
		Handler: func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("pong!"))
		},
	},
	{
		Path:    "/payment",
		Methods: []string{http.MethodPost},
		Handler: handlers.CreatePaymentSession,
	},
	{
		Path:    "/payment/status/{uniqueKey}",
		Methods: []string{http.MethodGet},
		Handler: handlers.CheckPaymentStatus,
	},
	{
		Path:    "/webhook", // this webhook is common
		Methods: []string{http.MethodGet},
		Handler: handlers.WebhookHandler,
	},
	{
		Path:    "/setapikeys",
		Methods: []string{http.MethodPost},
		Handler: handlers.SetAPIKeysHandler,
	},
}

func InitRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Use(LoggingMiddleware)
	for _, route := range Routes {
		r.HandleFunc(route.Path, route.Handler).Methods(route.Methods...)
	}
	return r
}
