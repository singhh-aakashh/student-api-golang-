package response

import (
	"encoding/json"
	"net/http"

)

type Response struct{
	Status string `json:"status"`
	Error string	`json:"error"`
}

func WriteJson(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return 
	}
	
}

func GeneralError(err error) Response{
 return Response{
	Status: "Error",
	Error: err.Error(),
 }
}
