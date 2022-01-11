package Models

type Task struct {
	ID        int
	Name      string
	Status    string
	Priority  string
	CreatedAt string
	CreatedBy string
	DueDate   string
}

func (t *Task) TableName() string {
	return "task"
}
