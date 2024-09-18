package utils

import "github.com/patelajay745/Microservice-VideoApp/tweet/models"

type ResError struct {
	Success bool  `json:"success"`
	Error   error `json:"message"`
}

type ResMessage struct {
	Success bool   `jsno:"sucess"`
	Message string `json:"message"`
}

type ResTweets struct {
	StatusCode int            `json:"statusCode"`
	Data       []models.Tweet `json:"data"`
	Message    string         `json:"message"`
	Success    bool           `json:"success"`
}
