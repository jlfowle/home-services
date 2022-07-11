package endpoints

type GetRequest struct{}

type GetResponse struct {
	Err string `json:"err,omitempty"`
}

type ServiceStatusRequest struct{}

type ServiceStatusResponse struct {
	Code int    `json:"status"`
	Err  string `json:"err,omitempty"`
}
