package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UssdRequest struct {
	// The unique session id
	SessionId   string `form:"sessionId"`
	ServiceCode string `form:"serviceCode"`
	PhoneNumber string `form:"phoneNumber"`
	Text        string `form:"text"`
}

func main() {
	r := gin.Default()

	r.GET("/get", func(context *gin.Context) {
		var request UssdRequest

		err := context.ShouldBindQuery(&request)

		if err != nil {
			fmt.Print(err)
		}

		fmt.Println(request.SessionId, " Session Id")

		context.JSON(http.StatusOK, gin.H{"message": "Good"})
	})

	r.Run()
}

/**
What I want to do is;
-	Select a set of predefined questions array and store in data structure
-	Create abstract layer to fetch questions
-	Use redis container and bind to local host
-	Connect to redis from our golang app
-	Implement endpoint to receive text and session from user, and send first screen.
*/
