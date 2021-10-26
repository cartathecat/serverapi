package main

import (
	"strconv"
	"testing"
	//	"fmt"
	//"log"
	"net/http"
	"net/http/httptest"
	//"testing"
	//	"os"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/help", HelpHandler).Methods("GET")
	router.HandleFunc("/listports", ListPortsHandler).Methods("GET")
	router.HandleFunc("/listports/all", ListAllPortsHandler).Methods("GET")
	router.HandleFunc("/port/{key}", PortKeyHandler).Methods("GET")

	return router
}

func Test_HelpHandler(t *testing.T) {

	request, _ := http.NewRequest("GET", "/help", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	t.Log("Response :" + strconv.Itoa(response.Code))
	assert.Equal(t, http.StatusOK, response.Code, "OK response is expected")

}

func Test_ListPortsHandler(t *testing.T) {

	request, _ := http.NewRequest("GET", "/listports", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	t.Log("Response :" + strconv.Itoa(response.Code))
	assert.Equal(t, http.StatusOK, response.Code, "OK response is expected")

}

func Test_ListAllPortsHandler(t *testing.T) {

	request, _ := http.NewRequest("GET", "/listports/all", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	t.Log("Response :" + strconv.Itoa(response.Code))
	assert.Equal(t, http.StatusOK, response.Code, "OK response is expected")

}

func Test_PortKeyHandlerWithValidKey(t *testing.T) {

	request, _ := http.NewRequest("GET", "/port/USADQ", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	t.Log("Response :" + strconv.Itoa(response.Code))
	assert.Equal(t, http.StatusOK, response.Code, "OK response is expected")

}

func Test_PortKeyHandlerWithInvalidKey(t *testing.T) {

	request, _ := http.NewRequest("GET", "/port/AAAA", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	t.Log("Response :" + strconv.Itoa(response.Code))
	assert.Equal(t, http.StatusNotFound, response.Code, "OK response is expected")

}
