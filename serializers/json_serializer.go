package serializers

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(writer http.ResponseWriter, v interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(v)
}
