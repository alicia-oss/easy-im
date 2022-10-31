package lua

import (
	"easy_im/internal/domain/im/model"
	"easy_im/pkg/redis"
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	t.Run("lua", func(t *testing.T) {

		sha, err := redis.Client.ScriptLoad(`
				-- key1: seq_key
local value = redis.call("Get", KEYS[1])
if not(value) then
    return 0
end
local data = {}
local reps = ":"
string.gsub(value,'[^'..reps..']+',function (w)
    table.insert(data,w)
end)
local curSeq = tonumber(data[1])
local maxSeq = tonumber(data[2])
if curSeq >= maxSeq then
    return 0
end
curSeq = curSeq + 1
local new = tostring(curSeq) .. ":" .. tostring(maxSeq)
redis.call("Set", KEYS[1], new)
return curSeq
					`).Result()
		if err != nil {
			fmt.Println(err)
			return
		}
		res := redis.Client.EvalSha(sha, []string{model.BuildSeqKey(10)})
		fmt.Println(res)

	})
}
