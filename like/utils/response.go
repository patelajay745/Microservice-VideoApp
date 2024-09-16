package utils

import "github.com/patelajay745/Microservice-VideoApp/like/models"

type ResMessage struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// ResError //
// response struct
type ResError struct {
	Success bool  `json:"success"`
	Error   error `json:"message"`
}

type ResLikes struct {
	StatusCode int           `json:"statusCode"`
	Data       []models.Like `json:"data"`
	Message    string        `json:"message"`
	Success    bool          `json:"success"`
}
