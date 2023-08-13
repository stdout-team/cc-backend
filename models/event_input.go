package models

type EventInput struct {
	Title       string      `json:"title" binding:"required"`
	Schedule    [][3]string `json:"schedule" binding:"required"`
	Location    string      `json:"location" binding:"required"`
	Interests   []string    `json:"interests" binding:"required"`
	Description string      `json:"description"`
}
