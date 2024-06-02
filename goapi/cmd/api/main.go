package main

import (
	"fmt"
	"net/http"

	"github.com/marcosterpe/goapi/internal/handlers"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")
	fmt.Println(`
 ______     ______      ______    ______    __
/\  ___\   /\  __ \    /\  __ \  /\  == \  /\ \
\ \ \__ \  \ \ \/\ \   \ \  __ \ \ \  ___\ \ \ \
 \ \_____\  \ \_____\   \ \_\ \_\ \ \_\     \ \_\
  \/_____/   \/_____/    \/_/\/_/  \/_/      \/_/`)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Failed to start the server")

	}
}
