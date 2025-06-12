package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type User struct {
	ID       uint32 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	r := gin.New()
	r.GET("/test/get-users", func(c *gin.Context) {
		// System information logging
		log.Printf("Number of CPUs: %d", runtime.NumCPU())
		log.Printf("GOMAXPROCS: %d", runtime.GOMAXPROCS(0))

		numCPU := runtime.NumCPU()
		users := make([]User, 10_000)

		var wg sync.WaitGroup
		chunk := len(users) / numCPU

		var totalTime int64

		// High-resolution start time
		start := time.Now()

		for i := 0; i < numCPU; i++ {
			wg.Add(1)
			go func(threadID int) {
				defer wg.Done()

				iterStart := time.Now()

				end := (threadID + 1) * chunk
				if end > len(users) {
					end = len(users)
				}
				start := threadID * chunk

				for j := start; j < end; j++ {
					strJ := strconv.Itoa(j)
					users[j] = User{
						ID:       uint32(j),
						Username: "user" + strJ,
						Email:    "user" + strJ + "@gmail.com",
						Password: "password" + strJ,
					}
				}

				atomic.AddInt64(&totalTime, time.Since(iterStart).Microseconds())
			}(i)
		}
		wg.Wait()

		overallTime := time.Since(start)
		log.Printf("Total execution time: %v microseconds", overallTime.Microseconds())
		log.Printf("Cumulative goroutine time: %v microseconds", atomic.LoadInt64(&totalTime))

		c.JSON(200, users)
	})

	err := r.Run(":8060")
	if err != nil {
		return
	}
}
