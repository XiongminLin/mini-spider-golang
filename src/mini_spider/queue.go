/* queue.go - FIFO */
/*
modification history
--------------------
2017/7/17, by linxiongmin, create
*/
/*
DESCRIPTION
*/
package mini_spider

import (
	"container/list"
	"errors"
	"sync"
)

/* max queue length */
const (
    MAX_QUEUE_LEN = 65535
)

/* queue */
type Queue struct {
	lock   sync.Mutex
	cond   *sync.Cond
	tasks  *list.List
	maxLen int         // max queue length
	unfinished int	   // number of unfinished tasks
}

/* Initialize the queue */
func (q *Queue) Init() {
	q.cond = sync.NewCond(&q.lock)
	q.tasks = list.New()
	q.maxLen = MAX_QUEUE_LEN
}

/* Add to the queue */
func (q *Queue) Add(task *CrawlTask) error {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	var err error

	if q.tasks.Len() >= q.maxLen {
		err = errors.New("Queue is full")
	} else {
		q.tasks.PushBack(task)
		q.unfinished += 1
		q.cond.Signal()
		err = nil
	}
	
	return err
}

/* pop a task from the queue */
func (q *Queue) Pop() *CrawlTask {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	for q.tasks.Len() == 0 {
		q.cond.Wait()
	}

	task := q.tasks.Front()
	q.tasks.Remove(task)

	return task.Value.(*CrawlTask)
}

/* Get length of the queue */
func (q *Queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()

	var len int
	len = q.tasks.Len()
	
	return len
}

/* Finish one task */
func (q *Queue) FinishOneTask() {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.unfinished -= 1
}

/* get count of unfinished tasks */
func (q *Queue) GetUnfinished() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	
	ret := q.unfinished
	
	return ret
}

/* set max queue length */
func (q *Queue) SetMaxLen(maxLen int) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.maxLen = maxLen
}
