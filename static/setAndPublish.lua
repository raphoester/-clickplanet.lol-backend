redis.call("SET", KEYS[1], ARGV[1]);
local message = "{\"" .. KEYS[1] .. "\":\"" .. ARGV[1] .. "\"}";
redis.call("PUBLISH", ARGV[2], message); return "K"
