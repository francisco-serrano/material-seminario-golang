package views

type CreateCarRequest struct {
	Brand      string `json:"brand"`
	Model      string `json:"model"`
	Year       int    `json:"year"`
	EngineType string `json:"engine_type"`
}
