package controllers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ExampleHandler(successChanceThreshold, fastChanceThreshold, fastResponseThreshold int) func(*gin.Context) {
	return func(c *gin.Context) {
		// simulate request taking time
		latency := getRandomLatencyMs(1, fastResponseThreshold)
		slowChance := generateRandomInt(1, 100)
		if slowChance > fastChanceThreshold {
			latency = getRandomLatencyMs(fastResponseThreshold+1, 1000)
		}
		time.Sleep(latency)

		// simulate a random status code
		c.JSON(getRandomStatusCode(successChanceThreshold), gin.H{
			"response": "pong",
		})
	}
}

// getRandomLatencyMs generates a random request latency between min and max in millisecond
func getRandomLatencyMs(min, max int) time.Duration {
	return time.Duration(generateRandomInt(min, max) * int(time.Millisecond))
}

// getRandomStatusCode generates a random status code that's 200 successChancePercentageThreshold% of the time
func getRandomStatusCode(successChancePercentageThreshold int) int {
	codes := []int{
		http.StatusBadRequest,
		http.StatusGatewayTimeout,
	}

	chance200 := generateRandomInt(1, 100)
	if chance200 <= successChancePercentageThreshold {
		return http.StatusOK
	}

	return codes[generateRandomInt(0, len(codes)-1)]
}

// generateRandomInt generates a random integer between min and max
// note that in go 1.20, there is no need to call rand.seed,
// but in older versions it should be called
func generateRandomInt(min int, max int) int {
	return rand.Intn(max-min+1) + min
}
