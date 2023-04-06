package errors

import (
	"encoding/json"
	"net/http"
)

func JSONError(writer http.ResponseWriter, detail string, code int) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.Header().Set("X-Content-Type-Options", "nosniff")
	writer.WriteHeader(code)
	json.NewEncoder(writer).Encode(HTTPErrorResponse{Detail: detail})
}
