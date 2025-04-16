package bredis

const (
	getSingleData = `
		local out = {}
		local result = {redis.call("getex", KEYS[1], "ex", ARGV[1])}
		table.insert(result, 1, ARGV[2])	
		table.insert(out, result)
		return out
	`

	getZSetData = `
		local out = {}
		local results = redis.call('ZRANGE', KEYS[1], 0, -1)
		table.insert(results, 1, ARGV[2])	
		table.insert(out, results)
		return out
	`

	setZSetData = `
		local count = redis.call('ZCARD', KEYS[1])
		local score = count + 1
		for i = 2, #ARGV do
			redis.call('ZADD', KEYS[1], score, ARGV[i])
			score = score + 1
		end

		return score - 1
	`
)
