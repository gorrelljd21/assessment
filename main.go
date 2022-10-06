package main

import (
	"math/rand"

	"github.com/gin-gonic/gin"

	"net/http"
)

type questions struct {
	ID       string `json:"id"`
	Question string `json:"question"`
}

func main() {
	r := gin.Default()
	r.GET("/question/:id", getQuestionByID)
	r.GET("/question", getRandomQuestion)

	r.Run("0.0.0.0:8080")
}

func getQuestionByID(c *gin.Context) {
	id := c.Param("id")

	getQuestion := QuestionOfTheDay[id]
	getSpecificQuestion := getQuestion.Question

	c.JSON(http.StatusOK, getSpecificQuestion)
}

func getRandomQuestion(c *gin.Context) {
	questionSlice := []string{}

	for k := range QuestionOfTheDay {
		questionSlice = append(questionSlice, k)
	}

	randNum := rand.Intn(len(questionSlice))
	randKey := questionSlice[randNum]
	randQuestion := QuestionOfTheDay[randKey]

	c.JSON(http.StatusOK, randQuestion)
}

// func getListOfQuestions(c *gin.Context) {
// 	questionsList := []questions{}

// 	for _, q := range QuestionOfTheDay {
// 		questionsList = append(questionsList, q)
// 	}
// }

var QuestionOfTheDay = map[string]questions{
	"88137df0-b7a8-4c91-ace7-a8a27953fb22": {"88137df0-b7a8-4c91-ace7-a8a27953fb22", "What is your favorite book?"},
	"0f1830ec-5a46-437c-8531-1917ae04e908": {"0f1830ec-5a46-437c-8531-1917ae04e908", "What is the fastest vehicle you have ever ridden in?"},
	"3bd8b8f1-4a9e-4724-8ba0-5c27de8a6405": {"3bd8b8f1-4a9e-4724-8ba0-5c27de8a6405", "If you were an animal, what animal would you be?"},
	"8af22d89-b64f-4508-9ef9-569a94847613": {"8af22d89-b64f-4508-9ef9-569a94847613", "What did you want to be when you grew up?"},
	"9480b2e7-2cd0-4ddd-b177-e2f5e5ef8b8f": {"9480b2e7-2cd0-4ddd-b177-e2f5e5ef8b8f", "What do you want to be when you grow up?"},
}
