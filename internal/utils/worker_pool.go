package utils

import (
	"GoTaskHub---A-Collaborative-Task-Management-System/internal/models"
	"fmt"
	"log"
	"os"
	"time"
)

func worker(id int, tasks chan models.UploadTask) {
	for task := range tasks {
		log.Printf("Worker %d processing file: %s", id, task.FileName)
		time.Sleep(2 * time.Second)
		fmt.Println(fmt.Sprintf("./uploads/%s", task.FileName))
		err := os.WriteFile(fmt.Sprintf("uploads/%s", task.FileName), task.Content, os.ModePerm)
		if err != nil {
			log.Printf("Worker %d failed to save file %s: %v", id, task.FileName, err)
		} else {
			log.Printf("Worker %d completed file: %s", id, task.FileName)
		}
	}
}

func StartWorkerPool(uploadQueue chan models.UploadTask) {
	workerCount := 3
	for i := 1; i <= workerCount; i++ {
		go worker(i, uploadQueue)
	}

	go func() {
		for {
			time.Sleep(time.Second * 10)
			if len(uploadQueue) > cap(uploadQueue)/2 {
				workerCount++
				go worker(workerCount, uploadQueue)
				log.Printf("Increased worker count to %d", workerCount)
			}
		}
	}()
}
