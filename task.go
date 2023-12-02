package qui

type TaskStatus string

const (
	TaskStatusPending  TaskStatus = "pending"
	TaskStatusComplete TaskStatus = "complete"
	TaskStatusFailed   TaskStatus = "failed"
)

type Task struct {
	ID      string
	Status  TaskStatus
	Payload []byte
}

type GenericTask[T any] struct {
	ID      string
	Status  TaskStatus
	Payload T
}
