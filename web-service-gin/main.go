package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var answer = []string{}

func addNum(c *gin.Context) {
	num, ok := c.GetQuery("num")
	if ok {
		i, _ := strconv.ParseInt(num, 10, 64)
		value := ""

		if i%3 == 0 && i%5 == 0 {
			value = "FizzBuzz"

		} else if i%3 == 0 {
			value = "Fizz"
		} else if i%5 == 0 {
			value = "Buzz"
		} else {
			value = num
		}

		if int(i) > len(answer) {
			newSlice := make([]string, int(i)-len(answer))
			newSlice[len(newSlice)-1] = value
			answer = append(answer, newSlice...)
		} else {
			answer[i-1] = value
		}
	}
}

func getAnswer(c *gin.Context) {
	addNum(c)
	c.IndentedJSON(http.StatusOK, answer)
}

func main() {
	router := gin.Default()
	router.GET("/fizz-buz-server/run", getAnswer)

	router.Run("localhost:8080")
}
