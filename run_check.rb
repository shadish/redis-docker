require "redis"

if ARGV.length == 0
	throw 'Argument required'
end

input = ARGV[0]
redis = Redis.new(host: "127.0.0.1", port: 6379)
p redis.set('mykey', input)
result = redis.get('mykey')
puts "Got result from redis: #{result}"
puts "Complete."