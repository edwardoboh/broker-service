package main

import (
	"encoding/json"
	"net/http"
)

type JsonResponse struct{
	Error bool `json:"error"`
	Message string `json:"message"`
	Data any `json:"data,omitEmpty"`
}

func welcome(w http.ResponseWriter, r *http.Request){
	payload := JsonResponse{Error: false, Message: "Handler Hit", Data: []string{"atatus", "message", "done"}}
	
	// response holds the formated json byte slice to be sent to the response write
	//
	// here is a more detailed explanation of the content of this variable
	response, error := json.Marshal(payload)
	if error != nil {
		http.Error(w, "An error occured while parsing data", http.StatusInternalServerError)
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(response)
}