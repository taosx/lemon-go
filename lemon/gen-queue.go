// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package main

import "sync"

type TaskQueue struct {
	items []Task
	lock  sync.RWMutex
}

// 创建队列
func (q *TaskQueue) New() *TaskQueue {
	q.items = []Task{}
	return q
}

// 如队列
func (q *TaskQueue) Append(t Task) {
	q.lock.Lock()
	q.items = append(q.items, t)
	q.lock.Unlock()
}

// 出队列
func (q *TaskQueue) Pop() *Task {
	q.lock.Lock()
	item := q.items[0]
	q.items = q.items[1:len(q.items)]
	q.lock.Unlock()
	return &item
}

// 获取队列的第一个元素，不移除
func (q *TaskQueue) Front() *Task {
	q.lock.Lock()
	item := q.items[0]
	q.lock.Unlock()
	return &item
}

// 判空
func (q *TaskQueue) IsEmpty() bool {
	return len(q.items) == 0
}

// 获取队列的长度
func (q *TaskQueue) Size() int {
	return len(q.items)
}
