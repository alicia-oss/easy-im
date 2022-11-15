package pkg

import (
	"easy_im/pkg/log"
	"fmt"
	"testing"
	"time"
)

func TestMsgRetryQueue(t *testing.T) {
	t.Run("t", func(t *testing.T) {
		queue := NewMsgRetryQueue()
		go queue.Run()
		go func() {
			for {
				s, ok := <-queue.Listen()
				if !ok {
					log.Info("stop", "test")
					return
				}
				log.Info(s, "test")
			}
		}()
		time.Sleep(3 * time.Second)
		queue.SubmitJob(&Job{
			Info:     "1",
			Duration: 2 * time.Second,
		})
		queue.SubmitJob(&Job{
			Info:     "2",
			Duration: 5 * time.Second,
		})
		queue.SubmitJob(&Job{
			Info:     "3",
			Duration: 2 * time.Second,
		})
		queue.SubmitJob(&Job{
			Info:     "4",
			Duration: 4 * time.Second,
		})
		time.Sleep(10 * time.Second)
		queue.SubmitJob(&Job{
			Info:     "6",
			Duration: 3 * time.Second,
		})
		time.Sleep(6 * time.Second)
		queue.Close()
		time.Sleep(20 * time.Second)
	})
	t.Run("ss", func(t *testing.T) {
		timer := time.NewTimer(24 * time.Hour)
		go func() {
			a := <-timer.C
			fmt.Println(a)
		}()
		timer.Reset(5 * time.Second)
		time.Sleep(30 * time.Second)
	})
}
