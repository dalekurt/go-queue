package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/dalekurt/go-queue/tasks"

	"github.com/hibiken/asynq"
)

func main() {
	// Create a new Redis connection for the client.
	redisConnection := asynq.RedisClientOpt{
		Addr: "localhost:6379", // Redis server address
	}

	// Create a new Asynq client.
	client := asynq.NewClient(redisConnection)
	defer client.Close()

	// Infinite loop to create tasks as Asynq client.
	for {
		// Generate a random user ID.
		userID := rand.Intn(1000) + 10

		// Generate a random asset ID.
		assetID := rand.Intn(1000) + 10

		// Set a delay duration to 2 minutes.
		delay := 2 * time.Minute

		// TODO: Generate a random string for the image task.
		image := "helloworld.jpg"

		// TODO: Generate a random string for the video task.
		video := "helloworld.mp4"

		// Define tasks.
		task1 := tasks.NewNotifyEmailTask(userID)
		task2 := tasks.NewReminderEmailTask(userID, time.Now().Add(delay))
		task3 := tasks.NewImageProcessingTask(assetID, image)
		task4 := tasks.NewVideoProcessingTask(assetID, video)

		// Process the task immediately in critical queue.
		if _, err := client.Enqueue(
			task1,                   // task payload
			asynq.Queue("critical"), // set queue for task
		); err != nil {
			log.Fatal(err)
		}

		// Process the task 2 minutes later in low queue.
		if _, err := client.Enqueue(
			task2,                  // task payload
			asynq.Queue("low"),     // set queue for task
			asynq.ProcessIn(delay), // set time to process task
		); err != nil {
			log.Fatal(err)
		}

		// Process the task immediately in default queue.
		if _, err := client.Enqueue(
			task3,                  // task payload
			asynq.Queue("default"), // set queue for task
		); err != nil {
			log.Fatal(err)
		}

		// Process the task immediately in default queue.
		if _, err := client.Enqueue(
			task4,                   // task payload
			asynq.Queue("critical"), // set queue for task
		); err != nil {
			log.Fatal(err)
		}
	}
}
