package internal

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

func InitApi() {
	router := httprouter.New()
	router.GET("/", HandleMain)
	router.GET("/servers", HandleSeconds)

	log.Printf("Server running on http://localhost:%s\n", os.Getenv("PORT"))
	server := &http.Server{
		Addr:              ":" + os.Getenv("PORT"),
		Handler:           router,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

type Msg struct {
	Msg string `json:"msg"`
}

func HandleMain(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&Msg{Msg: "Hello! Check out the repository https://github.com/TonimatasDEV/minestatistics-backend"})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func HandleSeconds(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(PlayerCounts)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
