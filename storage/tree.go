package storage

import "time"

type timeTree map[time.Time]*task 

var tree timeTree

func NewTree() timeTree {
    return make(timeTree)
}

func (tree timeTree) Add(time time.Time, t *task) {
    tree[time] = t
}

func (tree timeTree) Find(time time.Time) (t *task, ok bool) {
    t, ok = tree[time]
    return
}
