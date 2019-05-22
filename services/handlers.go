package services

import "net/http"



func  GetAccount(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("{\"result\":\"OK\"}"))
}
