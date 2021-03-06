package server

const (
	Running = "running"
	Stopped = "stopped"
)

type State struct {
	MockID string `json:"mock_id"`
	URL    string `json:"url"`
	Status string `json:"status"`
}

type Controller struct {
	Shutdown func()
}
