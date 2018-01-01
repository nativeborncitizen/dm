package storage

type task struct {
    desc string
}

func NewTask(desc string) *task {
    return &task{desc}
}

func (t task) GetDesc() string {
    return t.desc
}
