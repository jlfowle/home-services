package endpoints

type PingRequest struct{}

type PingResponse struct {
	Output string `json:"output"`
}

type HealthRequest struct{}

type HealthResponse struct {
	Code int    `json:"status"`
	Err  string `json:"err,omitempty"`
}
