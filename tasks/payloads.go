package tasks

import (
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

const (
	// TypeNotifyEmail is a name of the task type
	// for sending a notify email.
	TypeNotifyEmail = "email:notify"

	// TypeReminderEmail is a name of the task type
	// for sending a reminder email.
	// TypeReminderEmail = "email:reminder"

	// TypeImageProcessing is a name of the task type
	// for processing an image.
	// TypeImageProcessing = "image:processing"

	// TypeVideoProcessing is a name of the task type
	// for processing an video1.
	// TypeVideoProcessing = "video:processing"
)

// NewNotifyEmailTask task payload for a new notify email.
func NewNotifyEmailTask(id int) *asynq.Task {
	// Specify task payload.
	payload := map[string]interface{}{
		"user_id": id, // set user ID
	}

	jsonString, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)

	}

	fmt.Println(jsonString)

	// Return a new task with given type and payload.
	return asynq.NewTask(TypeNotifyEmail, jsonString)
}

// NewReminderEmailTask task payload for a reminder email.
// func NewReminderEmailTask(id int, ts time.Time) *asynq.Task {
// 	// Specify task payload.
// 	payload := map[string]interface{}{
// 		"user_id": id,          // set user ID
// 		"sent_in": ts.String(), // set time to sending
// 	}

// 	// Return a new task with given type and payload.
// 	return asynq.NewTask(TypeReminderEmail, payload)
// }

// // NewImageProcessingTask task payload for processing an image.
// func NewImageProcessingTask(id int, image string) *asynq.Task {
// 	// Specify task payload.
// 	payload := map[string]interface{}{
// 		"asset_id": id, // set asset ID
// 		"image":    image,
// 	}

// 	// Return a new task with given type and payload.
// 	return asynq.NewTask(TypeImageProcessing, payload)
// }

// // NewVideoProcessingTask task payload for processing an video.
// func NewVideoProcessingTask(id int, video string) *asynq.Task {
// 	// Specify task payload.
// 	payload := map[string]interface{}{
// 		"asset_id": id, // set asset ID
// 		"video":    video,
// 	}

// 	// Return a new task with given type and payload.
// 	return asynq.NewTask(TypeVideoProcessing, payload)
// }
