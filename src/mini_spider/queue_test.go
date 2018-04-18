/* queue_test.go - test for queue.go  */
/*
modification history
--------------------
2017/07/23, by Xiongmin LIN, create
*/
/*
DESCRIPTION
*/
package mini_spider

import (
	"testing"
)

func TestQueue(t *testing.T) {
	var queue Queue
	queue.Init()

	queue.Add(&CrawlTask{"http://www.baidu.com", 2, nil})
	queue.Add(&CrawlTask{"http://www.sina.com", 1, nil})
	
	queue.SetMaxLen(2)

	err := queue.Add(&CrawlTask{"http://www.test.com", 1, nil})

	if err == nil {
		t.Error("queue is full, qeueu.Add should occur an error")
	}

	task := queue.Pop()
	if task == nil || task.Url != "http://www.baidu.com" {
        t.Errorf("www.baidu.com should be poped first")
    }

	qLen := queue.Len()
	if qLen != 1 {
		t.Error("queue.Len() should be 1, now it's %d", qLen)
	}
	
	n := queue.GetUnfinished()
	if n != 2 {
		t.Error("queue.GetUnfinished() should be 2, now it's %d", n)
	}
	
	queue.FinishOneTask()
	n = queue.GetUnfinished()
	if n != 1 {
		t.Error("queue.GetUnfinished() should be 1, now it's %d", n)
	}

	task = queue.Pop()
	if task == nil || task.Url != "http://www.sina.com" || task.Depth != 1 {
		t.Error("queue should pop http://www.sina.com with depth 1")
	}

	task = queue.Pop()
	if task != nil {
		t.Error("queue should pop nil")
	}

	queue.FinishOneTask()
	n = queue.GetUnfinished()
	if n != 0 {
		t.Error("all tasks in queue had finished, GetUnfinished() should return 0")
	}

}