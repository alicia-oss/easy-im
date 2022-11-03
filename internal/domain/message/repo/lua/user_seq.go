package lua

import (
	"easy_im/internal/domain/message/pkg"
	"easy_im/pkg/log"
	"easy_im/pkg/redis"
	"fmt"
	"strconv"
	"strings"
)

var getSeq string

func init() {
	getSeqScript := `local value = redis.call("Get", KEYS[1])
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
						return new`
	getSeq, _ = redis.Client.ScriptLoad(getSeqScript).Result()
}

func GetSeq(key string) (uint64, uint64, error) {
	result, err := redis.Client.EvalSha(getSeq, []string{key}).Result()
	if err != nil {
		log.Error(fmt.Sprintf("GetSeq error;%v", err), pkg.ModuleNameRepoUserSeq)
		return 0, 0, err
	}
	res := strings.Split(result.(string), ":")
	curSeq, err := strconv.ParseUint(res[0], 10, 64)
	maxSeq, err := strconv.ParseUint(res[1], 10, 64)
	if err != nil {
		log.Error(fmt.Sprintf("GetSeq ParseUint error;%v, %v, %v", err, res[0], res[1]), pkg.ModuleNameRepoUserSeq)
		return 0, 0, err
	}
	return curSeq, maxSeq, nil
}
