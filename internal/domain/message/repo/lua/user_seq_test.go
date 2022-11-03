package lua

import (
	"easy_im/internal/domain/message/model"
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("lua", func(t *testing.T) {
		seq, max, err := GetSeq(model.BuildUserSeqKey(10))
		if err != nil {
			return
		}
		fmt.Println(seq, max)

	})
}
