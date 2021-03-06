package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lukemgriffith/captainhook"
	"github.com/lukemgriffith/captainhook/util"
	"log"
	"net/http"
)

//TODO: Document
type EndpointController struct {
	service captainhook.EndpointService
	log     *log.Logger
}

//TODO: Document
func NewEndpointController(es captainhook.EndpointService) *EndpointController {
	log := util.NewLog("EndpointController ")
	return &EndpointController{es, log}
}

// Get recieved a single instance of Endpoint
func (e *EndpointController) Get(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	if name, ok := vars["name"]; ok {
		obj, err := e.service.Endpoint(name)

		if err != nil {
			e.log.Print("Cannot find endpoint", name, "no content.")
			w.WriteHeader(http.StatusNoContent)
			return
		}

		json, err := json.Marshal(obj)

		if err != nil {
			e.log.Print("json error", err)
			e.log.Print("Unable to marshal endpoint", name, "to json.")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	} else {

		obj, err := e.service.Endpoints()

		if err != nil {
			e.log.Print("Unable to get all endpoints.")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		json, err := json.Marshal(obj)
		if err != nil {
			e.log.Print("Unable to marshal endpoints to json.")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(json)
	}

}

func (e *EndpointController) Post(w http.ResponseWriter, r *http.Request)   {
	w.WriteHeader(http.StatusMethodNotAllowed)
}
func (e *EndpointController) Patch(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(http.StatusMethodNotAllowed)
}
func (e *EndpointController) Delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}
