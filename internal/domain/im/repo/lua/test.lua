-- key1: seq_key,
local value = redis.call("Get", KEYS[1])
if not(value) then
    return -1
end
local data = {}
local reps = ":"
string.gsub(value,'[^'..reps..']+',function (w)
    table.insert(data,w)
end)
local curSeq = tonumber(data[1])
local maxSeq = tonumber(data[2])
if curSeq >= maxSeq then
    return -1
end
curSeq = curSeq + 1
local new = tostring(curSeq) .. ":" .. tostring(maxSeq)
redis.call("Set", KEYS[1], new)
return curSeq
