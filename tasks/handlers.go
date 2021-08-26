package tasks

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

// HandleNotifyEmailTask handler for notify email task.
func HandleNotifyEmailTask(c context.Context, t *asynq.Task) error {
	// Get user ID from given task.
	fmt.Println("Hello")
	var result map[string]interface{}
	json.Unmarshal(t.Payload(), &result)
	fmt.Println(result)

	// id, err := t.Payload.GetInt("user_id")
	// if err != nil {
	// 	return err
	// }

	// Dummy message to the worker's output.
	// fmt.Printf("Send notify Email to User ID %d\n", id)

	return nil
}

// HandleReminderEmailTask for reminder email task.
// func HandleReminderEmailTask(c context.Context, t *asynq.Task) error {
// 	// Get int with the user ID from the given task.
// 	id, err := t.Payload.GetInt("user_id")
// 	if err != nil {
// 		return err
// 	}

// Get string with the sent time from the given task.
// time, err := t.Payload.GetString("sent_in")
// if err != nil {
// 	return err
// }

// Dummy message to the worker's output.
// 	fmt.Printf("Send Reminder Email to User ID %d\n", id)
// 	fmt.Printf("Reason: time is up (%v)\n", time)

// 	return nil
// }

// HandleImageProcessTask for image processing task.
// func HandleImageProcessingTask(c context.Context, t *asynq.Task) error {
// 	// Get user ID from given task.
// 	id, err := t.Payload.GetInt("asset_id")
// 	if err != nil {
// 		return err
// 	}

// Dummy message to the worker's output.
// fmt.Printf("Processing Image with Asset ID %d\n", id)

// 	return nil
// }

// HandleVideoProcessTask for video processing task.
// func HandleVideoProcessingTask(c context.Context, t *asynq.Task) error {
// 	// Get user ID from given task.
// 	id, err := t.Payload.GetInt("asset_id")
// 	if err != nil {
// 		return err
// 	}

// 	// Dummy message to the worker's output.
// 	fmt.Printf("Processing Video with Asset ID %d\n", id)

// 	return nil
// }
