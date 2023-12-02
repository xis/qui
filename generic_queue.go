package qui

import (
	"context"
	"encoding/json"

	"github.com/rs/xid"
)

type GenericQueue[T any] struct {
	queueName string
	broker    TaskBroker
}

type NewGenericQueueParams struct {
	Name   string
	Broker TaskBroker
}

func NewGenericQueue[T any](params NewGenericQueueParams) *GenericQueue[T] {
	return &GenericQueue[T]{
		queueName: params.Name,
		broker:    params.Broker,
	}
}

func (queue *GenericQueue[T]) GetTask(ctx context.Context) (GenericTask[T], error) {
	task, err := queue.broker.GetTask(ctx, queue.queueName)
	if err != nil {
		return GenericTask[T]{}, err
	}

	var payload T

	err = json.Unmarshal(task.Payload, &payload)
	if err != nil {
		return GenericTask[T]{}, err
	}

	return GenericTask[T]{
		ID:      task.ID,
		Status:  task.Status,
		Payload: payload,
	}, nil
}

type CreateTaskParams[T any] struct {
	Payload  T
	Priority int
}

func (queue *GenericQueue[T]) CreateTask(ctx context.Context, payload T, priority int) (GenericTask[T], error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return GenericTask[T]{}, err
	}

	task := Task{
		ID:      xid.New().String(),
		Status:  TaskStatusPending,
		Payload: payloadBytes,
	}

	_, err = queue.broker.CreateTask(ctx, queue.queueName, task)
	if err != nil {
		return GenericTask[T]{}, err
	}

	return GenericTask[T]{
		ID:      task.ID,
		Status:  task.Status,
		Payload: payload,
	}, nil
}
