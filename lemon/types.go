package main

type Task struct {
	TaskID     string            `json:"mid"`
	HTTPMethod string            `json:"method"`
	Host       string            `json:"host"`
	Path       string            `json:"path"`
	Header     map[string]string `json:"header"`
	Param      map[string]string `json:"param"`
	Cookie     string            `json:"cookie"`
	Payload    string            `json:"payload"`
}

type Result struct {
	Status       string `json:"status"`
	TaskID       string `json:"task_id"`
	ResponseCode int    `json:"code"`
	Data         string `json:"data"`
	FetchedTime  int64  `json:"time"`
	UserAgent    string `json:"ua"`
}
