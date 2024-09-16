```markdown
# Like Microservice

This microservice is responsible for handling "like" actions on videos, comments, and tweets. It is written in Go using the Gin framework and MongoDB as the database.

## Table of Contents

- [Installation](#installation)
- [Configuration](#configuration)
- [Routes](#routes)
- [Models](#models)
- [Middleware](#middleware)
- [Controllers](#controllers)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. Clone the repository:
2. Install dependencies:
   go mod tidy
3. Set up your MongoDB instance and update the connection URI in the code.

4. Run the service:
   go run main.go

## Configuration

Ensure you have a MongoDB instance running and update the MongoDB URI in the `.env` file:

## Routes

The following routes are available in this microservice:

- **Toggle Like on Video**: `POST /api/v1/likes/toggle/v/:videoId`
- **Toggle Like on Comment**: `POST /api/v1/likes/toggle/c/:commentId`
- **Toggle Like on Tweet**: `POST /api/v1/likes/toggle/t/:tweetId`
- **Get Liked Videos**: `GET /api/v1/likes/videos`
```

## Models

The `Like` struct represents a like action in the database:

````go
type Like struct {
    ID        primitive.ObjectID  `bson:"_id" json:"_id"`
    Video     *primitive.ObjectID `bson:"video,omitempty" json:"video,omitempty"`
    Tweet     *primitive.ObjectID `bson:"tweet,omitempty" json:"tweet,omitempty"`
    Comment   *primitive.ObjectID `bson:"comment,omitempty" json:"comment,omitempty"`
    LikedBy   primitive.ObjectID  `bson:"likedBy" json:"likedBy"`
    CreatedAt time.Time           `bson:"createdAt" json:"createdAt"`
    UpdatedAt time.Time           `bson:"updatedAt" json:"updatedAt"`
}


## Middleware

The `VerifyJWT` middleware is used to authenticate requests. Ensure you have implemented this middleware to validate JWT tokens.

## Controllers

The controllers handle the business logic for each route. Here are the main functions:

- `ToggleVideoLike(c *gin.Context, client *mongo.Client)`
- `ToggleCommentLike(c *gin.Context, client *mongo.Client)`
- `ToggleTweetLike(c *gin.Context, client *mongo.Client)`
- `GetLikedVideos(c *gin.Context, client *mongo.Client)`

## Usage

To use this microservice, make HTTP requests to the defined routes. Ensure you include a valid JWT token in the request headers for authentication.

### Example Request

```sh
curl -X POST http://localhost:8001/api/v1/likes/toggle/v/12345 \
-H "Authorization: Bearer <your_jwt_token>"
````

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

Free to use :)
