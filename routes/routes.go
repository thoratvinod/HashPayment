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
		Path:    "/webhook/success",
		Methods: []string{http.MethodGet},
		Handler: handlers.SuccessWebhook,
	},
	{
		Path:    "/webhook/cancel",
		Methods: []string{http.MethodGet},
		Handler: handlers.CanceledWebhook,
	},
}

func InitRoutes() *mux.Router {
	r := mux.NewRouter()
	for _, route := range Routes {
		r.HandleFunc(route.Path, route.Handler).Methods(route.Methods...)
	}
	return r
}
