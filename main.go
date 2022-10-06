package main

import (
	"github.com/gin-gonic/gin"
)

type questions struct {
	ID       string `json:"id"`
	Question string `json:"question"`
}

func main() {
	r := gin.Default()

	r.Run("0.0.0.0:8080")
}

var QuestionOfTheDay = map[string]questions{
	"88137df0-b7a8-4c91-ace7-a8a27953fb22": {"88137df0-b7a8-4c91-ace7-a8a27953fb22", "What is your favorite book?"},
	"0f1830ec-5a46-437c-8531-1917ae04e908": {"0f1830ec-5a46-437c-8531-1917ae04e908", "What is the fastest vehicle you have ever ridden in?"},
	"3bd8b8f1-4a9e-4724-8ba0-5c27de8a6405": {"3bd8b8f1-4a9e-4724-8ba0-5c27de8a6405", "If you were an animal, what animal would you be?"},
	"8af22d89-b64f-4508-9ef9-569a94847613": {"8af22d89-b64f-4508-9ef9-569a94847613", "What did you want to be when you grew up?"},
	"9480b2e7-2cd0-4ddd-b177-e2f5e5ef8b8f": {"9480b2e7-2cd0-4ddd-b177-e2f5e5ef8b8f", "What do you want to be when you grow up?"},
}
