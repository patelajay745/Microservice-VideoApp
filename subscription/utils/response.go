package utils

import "go.mongodb.org/mongo-driver/bson"

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

type ResSubscription struct {
	StatusCode int      `json:"statusCode"`
	Data       []bson.M `json:"data"`
	Message    string   `json:"message"`
	Success    bool     `json:"success"`
}
