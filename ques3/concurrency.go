// Q3. Make 15 concurrent request to https://loripsum.net/api and
// print the data received, wait for maximum 2 seconds for go routines
// to finish the task. If time runs out cancel the request using context.
// (25 Marks)

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

const url = "https://loripsum.net/api"

var wg sync.WaitGroup // New wait group
func Data(ctx context.Context) { // Using sync.Waitgroup to wait for goroutines to finish

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Errorf("Failed to build the request: %s", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("Failed to request url: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("Error request response with status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	log.Println(string(body))

	wg.Done()
}

func main() {

	wg.Add(15)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(2*time.Second))
	defer cancel()
	for i := 1; i < 15; i++ {
		defer wg.Done()
		time.Sleep(time.Second)
		go Data(ctx)

	}
	wg.Wait()

	log.Println("Program finished executing")
}
