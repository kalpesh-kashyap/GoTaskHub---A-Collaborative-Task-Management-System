package utils

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

func DownloadPool(taskQueue chan string, outputDir string) {
	workerCount := 3
	var wg sync.WaitGroup
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go dwonlaodWorker(i, taskQueue, outputDir, &wg)
	}

	go func() {
		if len(taskQueue) > cap(taskQueue)/5 {
			workerCount++
			go dwonlaodWorker(workerCount, taskQueue, outputDir, &wg)
			log.Printf("Increased worker count to %d", workerCount)
		}
	}()
	go func() {
		wg.Wait()
		log.Println("All workers have exited")
	}()
}

func dwonlaodWorker(id int, taskQueue chan string, outputDir string, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range taskQueue {
		log.Printf("Worker %d processing URL: %s", id, task)
		downloadAndSaveProcess(task, outputDir)
	}
}

func downloadAndSaveProcess(task string, outputDir string) []string {
	downloadedFile := make(chan []byte, 1)
	results := make(chan string, 2)

	var wg sync.WaitGroup

	wg.Add(2)
	go downloadFileFromURL(task, downloadedFile, results, &wg)
	go saveDownloadedFile(task, downloadedFile, results, &wg, outputDir)

	go func() {
		wg.Wait()
		close(results)
		close(downloadedFile)
	}()
	var allResults []string
	for result := range results {
		allResults = append(allResults, result)
	}
	return allResults
}

func downloadFileFromURL(url string, downloadedFile chan []byte, results chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	res, err := http.Get(url)
	if err != nil {
		results <- "failed to download " + url
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		results <- "faile to read " + url
		return
	}

	downloadedFile <- body
	results <- "successfuly downloaded " + url
}

func saveDownloadedFile(url string, downloadedFile chan []byte, results chan string, wg *sync.WaitGroup, outputDir string) {
	defer wg.Done()
	fileContent := <-downloadedFile
	if len(fileContent) == 0 {
		results <- "faile to save file " + url
		return
	}
	fileName := filepath.Base(url)
	outputPath := filepath.Join(outputDir, fileName)

	out, err := os.Create(outputPath)
	if err != nil {
		results <- "faile to create file for " + fileName
	}

	defer out.Close()

	_, err = out.Write(fileContent)

	if err != nil {
		results <- "failed to write to file for " + fileName
		return
	}

	results <- "successfully saved file " + fileName
}
