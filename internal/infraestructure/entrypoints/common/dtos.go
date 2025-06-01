package entrypoints

type SuccessAPIResponse[T any] struct {
	Data T `json:"data"`
}

type ErrorAPIResponse struct {
	Error    string         `json:"error"`
	Metadata map[string]any `json:"metadata,omitempty"`
}
