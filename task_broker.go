package qui

import (
	"context"
)

type TaskBroker interface {
	GetTask(ctx context.Context, queueName string) (Task, error)
	CreateTask(ctx context.Context, queueName string, task Task) (Task, error)
	CompleteTask(ctx context.Context, queueName string, id string) error
	FailTask(ctx context.Context, queueName string, id string) error
}
