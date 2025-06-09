package entrypoints

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func ReadJson(w http.ResponseWriter, r *http.Request, data any) error {
	maxBites := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBites))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

func WriteJSONError(w http.ResponseWriter, status int, message string, metadata map[string]any) error {
	w.Header().Set("Content-Type", "application/json")

	data := ErrorAPIResponse{
		Error:    message,
		Metadata: metadata,
	}

	return WriteJSON(w, status, &data)
}

func JsonResponse[T any](w http.ResponseWriter, status int, data T) error {
	return WriteJSON(w, status, &data)
}
