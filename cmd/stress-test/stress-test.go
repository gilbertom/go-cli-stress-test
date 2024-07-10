package stress_test

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Run is a function that executes the stress test.
func Run(url string, requests, concurrency int) {

	if requests < concurrency {
		concurrency = requests
	}

	requestsPerWorker := requests / concurrency
	extraRequests := requests % concurrency

	var wg sync.WaitGroup
	results := make(chan int, requests)
	client := &http.Client{}

	startTime := time.Now()

	for i := 0; i < concurrency; i++ {
		numRequests := requestsPerWorker
		if i < extraRequests {
			numRequests++
		}
		wg.Add(1)
		go worker(&wg, url, numRequests, results, client)
	}

	wg.Wait()
	close(results)

	elapsedTime := time.Since(startTime)

	// Gerar relatório
	total200 := 0
	statusCodes := make(map[int]int)

	for result := range results {
		if result == 200 {
			total200++
		}
		statusCodes[result]++
	}

	fmt.Printf("Tempo total gasto na execução                : %s\n", elapsedTime)
	fmt.Printf("Quantidade total de requests realizados      : %d\n", requests)
	fmt.Printf("Quantidade de requests com status HTTP 200   : %d\n", total200)
	fmt.Println("Distribuição de outros códigos de status HTTP:")

	if total200 == requests {
		fmt.Println("Nenhuma chamada retornou HTTP Status diferente de 200.")
		return
	}

	for code, count := range statusCodes {
		if code != 200 {
			fmt.Printf("  Status %d: %d\n", code, count)
		}
	}
}

func worker(wg *sync.WaitGroup, url string, numRequests int, results chan int, client *http.Client) {
	defer wg.Done()
	for i := 0; i < numRequests; i++ {
		resp, err := client.Get(url)
		if err != nil {
			results <- 500
			continue
		}
		results <- resp.StatusCode
		resp.Body.Close()
	}
}
