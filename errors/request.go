package errors

import "net/http"

const INTERNAL_SERVER_ERROR = `{"message": "Internal Server Error"}`

func MethodNotSupported(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte(`{"message": "Method not supported"}`))
}

func ContentTypeNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusRequestEntityTooLarge)
	w.Write([]byte(`{"message":"Payload too large"}`))
}
