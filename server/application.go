package server

import (
	"github.com/gorilla/mux"
	"github.com/lukemgriffith/captainhook"
	"log"
	"net/http"
)

//TODO: Document
func New(es captainhook.EndpointService) http.Handler {

	log := NewLog("CaptainHook ")
	mux := mux.NewRouter()
	fs := http.FileServer(http.Dir("static"))
	AppServer := &AppServer{mux, log}

	hookEng := captainhook.NewHookEngine(log, &es)

	log.Println("Starting Application Server.")

	mux.Handle("/", fs)
	mux.HandleFunc("/webhook/{id}", hookEng.Hook)

	ec := RestController{log, NewEndpointController(es)}

	mux.HandleFunc("/endpoint", ec.ServeHTTP)
	mux.HandleFunc("/endpoint/{name}", ec.ServeHTTP)

	return AppServer

}

//TODO: Document
type AppServer struct {
	mux *mux.Router
	log *log.Logger
}

//TODO: Document
func (a *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
