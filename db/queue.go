package db

import (
	"log"
	"time"

	"sorn/core/models"

	"github.com/wkirk01/AlgoeDB"
)


type QueueItem map[string]interface{}

var queue *AlgoeDB.Database

func initQueue() {
	var err error
	queue, err = AlgoeDB.NewDatabase(&AlgoeDB.DatabaseConfig{Path: "queue.db"})
	if err != nil {
		log.Fatal(err)
	}
		
}

func AddToQueue(queueItem models.QueueItem) error {
	queueItemMap := map[string]interface{}{
		"track_id": queueItem.TrackID,
		"track_name": queueItem.TrackName,
		"created_at": queueItem.CreatedAt,
	}
	err := queue.InsertOne(queueItemMap)
	if err != nil {
		return err
	}
	return nil
}

func GetNextQueueItem() (models.QueueItem, error) {
	
	items := queue.FindMany(map[string]interface{}{})
	if items == nil {
		return models.QueueItem{}, nil
	}

	if len(items) == 0 {
		return models.QueueItem{}, nil
	}

	firstItem := items[0]
	queueItem := models.QueueItem{
		TrackID:    uint(firstItem["track_id"].(int)),
		CreatedAt:  firstItem["created_at"].(time.Time),
	}

	return queueItem, nil

}

func RemoveFromQueue(queueItem models.QueueItem) error {
	err := queue.DeleteOne(map[string]interface{}{
		"track_id": queueItem.TrackID,
	})
	if err != nil {
		return err
	}
	return nil
}

