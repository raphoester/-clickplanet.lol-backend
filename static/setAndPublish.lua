local currentValue = redis.call("GET", KEYS[1])
if currentValue ~= ARGV[1] then
    redis.call("SET", KEYS[1], ARGV[1])
    local message = "{\"" .. KEYS[1] .. "\":\"" .. ARGV[1] .. "\"}"
    redis.call("PUBLISH", ARGV[2], message)
end
return "K"