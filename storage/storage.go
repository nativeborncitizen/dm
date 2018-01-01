package storage

type task struct {
    desc string
}

func New(desc string) *task {
    return &task{desc}
}

func (t task) GetDesc() string {
    return t.desc
}
