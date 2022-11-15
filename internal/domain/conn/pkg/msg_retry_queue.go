package pkg

import (
	"easy_im/pkg/log"
	"easy_im/pkg/timing_queue"
	"time"
)

type Job struct {
	Info     string
	Duration time.Duration
}

func NewMsgRetryQueue() *MsgRetryQueue {
	return &MsgRetryQueue{
		tQueue:    timing_queue.NewTimeQueue(),
		timer:     time.NewTimer(24 * time.Hour),
		nextTime:  time.Now(),
		closeChan: make(chan struct{}),
		addChan:   make(chan *Job, 5),
		taskChan:  make(chan string, 5),
	}
}

type MsgRetryQueue struct {
	tQueue    *timing_queue.TimingQueue
	timer     *time.Timer
	nextTime  time.Time
	closeChan chan struct{}
	addChan   chan *Job
	taskChan  chan string
}

func (q *MsgRetryQueue) Run() {
	for {
		select {
		case <-q.closeChan:
			log.Info("MsgRetryQueue stop......", "domain_message_retry_queue")
			close(q.addChan)
			close(q.taskChan)
			return
		case job := <-q.addChan:
			q.addJob(job)
		case <-q.timer.C:
			q.produceTask()
		}
	}

}

func (q *MsgRetryQueue) Close() {
	close(q.closeChan)
}

func (q *MsgRetryQueue) Listen() <-chan string {
	return q.taskChan
}

func (q *MsgRetryQueue) SubmitJob(job *Job) {
	q.addChan <- job
}

func (q *MsgRetryQueue) addJob(job *Job) {
	doTime := time.Now().Add(job.Duration)
	_, v := q.tQueue.Peek()
	if doTime.Before(q.nextTime) || v == nil {
		q.timer.Reset(job.Duration)
		q.nextTime = doTime
	}
	q.tQueue.Push(doTime, job.Info)
}

func (q *MsgRetryQueue) produceTask() {

	_, v := q.tQueue.Pop()
	value := v.(string)
	q.taskChan <- value
	t, v := q.tQueue.Peek()
	if v == nil {
		return
	}
	q.timer = time.NewTimer(t.Sub(time.Now()))
}
