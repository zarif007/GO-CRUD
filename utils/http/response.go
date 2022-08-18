package http

import (
	"net/http"
)

type response struct {
	Status int         `json:"status"`
	Result interface{} `json:"result"`
}

func newResponse() *response {

}

func (resp *response) bytes() []byte {

}

func (resp *response) string() string {

}

func (resp *response) sendResponse(w http.ResponseWriter, r *http.Request) {

}

//200
func StatusNoContent() {

}

//400
func StatusBadRequest() {

}

//404
func StatusNoFound() {

}

//405
func StatusMethodNotAllowed() {

}
