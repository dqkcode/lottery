package main

import (
	"log"
	"lottery/lottery"
	"lottery/server"
	"lottery/types"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(server.LoggerMiddleware)

	routes := make([]types.Route, 0)

	lotteryHandler := lottery.NewHandler()

	routes = append(routes, lotteryHandler.Routes()...)

	for _, route := range routes {
		handler := http.Handler(route.HandlerFunc)
		for i := len(route.Middlewares) - 1; i >= 0; i-- {
			handler = route.Middlewares[i](handler)
		}
		router.Path(route.Path).Methods(route.Method).Handler(handler).Queries(route.Queries...)
	}

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
		// ReadTimeout:  10 * time.Second,
		// WriteTimeout: 10 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Panic(err)
	}

}
