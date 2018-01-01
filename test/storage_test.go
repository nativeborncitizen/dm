package storage

import (
    "testing"
    "time"
    "github.com/nativeborncitizen/dm/storage"
)

func TestNewTask(t *testing.T) {
    a := storage.NewTask("AAA")
    if a.GetDesc() != "AAA" {
        t.Error("Expected AAA, got", a.GetDesc())
    }
}

func TestNewTree(t *testing.T) {
    n := time.Now()
    tree := storage.NewTree()
    tree.Add(n, storage.NewTask("AAA"))
    res, ok := tree.Find(n)
    if !ok {
        t.Error("Expected task")    
    }
    if res.GetDesc() != "AAA" {
        t.Error("Expected AAA, got", res.GetDesc())    
    }
}
