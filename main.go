/*
Main package for Ports API
*/
package main

/*
	Contributors:
		Mick Moriarty
		v0.0.1
*/
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cartathecat/datahandler"

	"github.com/gorilla/mux"
)

/*
ErrorResponse ...
*/
type errorResponse struct {
	Code string `json:"code"`
	Info string `json:"info"`
	Msg  string `json:"message"`
}

var port string // Port number

/*
appHandler struct, that defines the function to invoke
*/
type appHandler struct {
	H func(http.ResponseWriter, *http.Request)
}

/*
Server function to invoke the correct handler
*/
func (ah appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	ah.H(w, r)

}

/*
Invoked from within the Go web server code
*/
func errorHandler(w http.ResponseWriter, r *http.Request) {

	errResp := errorResponse{Code: "E1001", Info: "Error", Msg: "Error has occurred while processing the request"}
	responseStatus := http.StatusInternalServerError

	// Set the heder details and write the error or response to the ResponseWriter
	w.Header().Set("content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(responseStatus)

	json.NewEncoder(w).Encode(errResp)

	return
}

/*
Returns JSON Port from GET request
*/
func PortKeyHandler(w http.ResponseWriter, r *http.Request) {

	// Extract the key from the URL .... /port/{key}
	vars := mux.Vars(r)
	key := vars["key"]

	resp := map[string]interface{}{}
	errResp := errorResponse{}
	responseStatus := http.StatusOK

	portID := datahandler.PortID{
		ID: key,
	}

	// Lookup the ID
	port, err := portID.GetPort()
	if err != nil {
		errResp = errorResponse{Code: "E1002", Info: "Port lookup error", Msg: err.Error()}
		responseStatus = http.StatusNotFound
	} else {
		resp[key] = port
	}

	// Set the heder details and write the error or response to the ResponseWriter
	w.Header().Set("content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(responseStatus)

	if responseStatus != http.StatusOK {
		json.NewEncoder(w).Encode(errResp)
	} else {
		json.NewEncoder(w).Encode(resp)
	}
	return
}

/*
List Ports ID function to return a list of PortIDs
*/
func ListPortsHandler(w http.ResponseWriter, r *http.Request) {

	responseStatus := http.StatusOK
	resp := datahandler.ListOfPorts()

	// Set the heder details and write the error or response to the ResponseWriter
	w.Header().Set("content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(responseStatus)

	json.NewEncoder(w).Encode(resp)
	return
}

/*
List All Ports handler to return a list of ALL ports
*/
func ListAllPortsHandler(w http.ResponseWriter, r *http.Request) {

	responseStatus := http.StatusOK
	resp := datahandler.ListAllPorts()

	// Set the heder details and write the error or response to the ResponseWriter
	w.Header().Set("content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(responseStatus)

	json.NewEncoder(w).Encode(resp)

	return
}

/*
Help function to show end-points
*/
func HelpHandler(w http.ResponseWriter, r *http.Request) {

	resp := map[string]interface{}{}
	responseStatus := http.StatusOK

	resp["listports"] = fmt.Sprintf("http://{host}:%s/listports", port)
	resp["listallports"] = fmt.Sprintf("http://{host}:%s/listports/all", port)
	resp["port/{ID}"] = fmt.Sprintf("http://{host}:%s/port/{ID}", port)

	// Set the heder details and write the error or response to the ResponseWriter
	w.Header().Set("content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(responseStatus)

	json.NewEncoder(w).Encode(resp)
	return
}

/*
Generic handler to show 404 errors
*/
func httpNotFoundHandler() http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		errResp := errorResponse{Code: "E1003", Info: "Error", Msg: "Error occurred while processing request"}
		responseStatus := http.StatusNotFound

		// Set the heder details and write the error or response to the ResponseWriter
		w.Header().Set("content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(responseStatus)

		json.NewEncoder(w).Encode(errResp)
		return
	})
}

/*
Craete a new router
*/
func newSubRouter(port string) *mux.Router {

	router := mux.NewRouter()
	subrouter := router.PathPrefix("/").Subrouter()

	// define port endpoints
	subrouter.Handle("/port/{key}", appHandler{PortKeyHandler})
	subrouter.Handle("/listports", appHandler{ListPortsHandler})
	subrouter.Handle("/listports/all", appHandler{ListAllPortsHandler})

	// Help end-point
	subrouter.Handle("/help", appHandler{HelpHandler})

	subrouter.NotFoundHandler = httpNotFoundHandler()

	return subrouter
}

/*
Init function to initialise data
*/
func init() {

	log.Print("Loading data ...")

	err := datahandler.LoadData()
	if err != nil {
		log.Fatalf("Error loading ports.json file, error : %s", err)
	}

	return

}

/*
Main API function
*/
func main() {

	log.Print("Server starting")

	// Get the port number
	port = os.Getenv("API_PORT")
	if port == "" {
		port = "9000"
	}

	subrouter := newSubRouter(port)
	log.Printf("Listening on port : %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, subrouter))

}
