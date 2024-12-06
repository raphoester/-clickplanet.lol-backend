-- KEYS[1] -> key to update
-- ARGV[1] -> new value
-- ARGV[2] -> stream name

local currentValue = redis.call("GET", KEYS[1])
if currentValue ~= ARGV[1] then
    redis.call("SET", KEYS[1], ARGV[1])
    redis.call("XADD", ARGV[2], "*", "t", KEYS[1], "n", ARGV[1], "o", (currentValue or ""))
end
return "K"
