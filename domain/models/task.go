package models

type Task struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Status    int    `json:"status"`
	Version   int    `json:"version"`
	CreatedAt string `json:"created_at"`
}

type TaskQueryParam struct {
}
