package utils

import (
	"net/http"
	"encoding/json"
)

func HandleJsonResponse(w http.ResponseWriter, response interface{}){
	jsonData,err := json.Marshal(response);
	FatalWithResponse(err,w);
	w.WriteHeader(http.StatusOK);
	w.Header().Set("Content-Type","application/json")
	w.Write(jsonData)



}

func FatalWithResponse(err error,w http.ResponseWriter){
	Fatal(err);
	if(err != nil){
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
}




