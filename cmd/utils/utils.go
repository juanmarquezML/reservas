package utils

import (
	"encoding/json"
	"learngo/api_only_go_task/pkg/model"
	"net/http"

	"github.com/go-playground/validator"
)

type MsgStatus struct {
	RespType string
	Mensaje  string
}
type Validador interface {
	Validation()
}

var usersResponses = map[string]MsgStatus{
	"error_writting":   {"Error", "Error writing response"},
	"email_obligatory": {"Error", "The email parameter is obligatory"},
	"error_marshaling": {"Error", "Error marshaling response"},
	"error_body":       {"Error", "Error parsing body content"},
	"error_post":       {"Error", "Error to Created Record"},
	"post_succefully":  {"Ok", "The record have been created successfully"},
}

func Validation(str interface{}) error {
	validate := validator.New()
	return validate.Struct(str)
}

func ShowMessage(w http.ResponseWriter, keyMensaje string) {
	var msg model.Resp
	var err error
	resp := make([]byte, 0)
	data, ok := usersResponses[keyMensaje]

	if ok {
		msg = model.Resp{data.RespType, data.Mensaje}
	} else {
		msg = model.Resp{"Error", keyMensaje}
	}
	resp, err = json.Marshal(msg)
	if err != nil {
		http.Error(w, "error marshaling response", http.StatusInternalServerError)
		return
	}
	_, err = w.Write(resp)
	if err != nil {
		http.Error(w, "error writing response", http.StatusInternalServerError)
		return

		/*
			http.Error(w, "error to created user", http.StatusBadRequest)
			return
		*/
	}

}
