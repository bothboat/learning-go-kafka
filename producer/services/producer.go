package services

import (
	"encoding/json"
	"events"
	"reflect"

	"github.com/IBM/sarama"
)

type EventProducer interface {
	Produce(event events.Event) error
}

type eventProducer struct {
	producer sarama.SyncProducer
}

func NewEventProducer(producer sarama.SyncProducer) EventProducer {
	return &eventProducer{producer}
}

func (obj eventProducer) Produce(event events.Event) error {
	topics := reflect.TypeOf(event).Name()

	value, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := sarama.ProducerMessage{
		Topic: topics,
		Value: sarama.ByteEncoder(value),
	}

	_, _, err = obj.producer.SendMessage(&msg)
	if err != nil {
		return err
	}

	return nil
}
