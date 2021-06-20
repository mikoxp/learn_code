from redis import Redis

redis_connection = Redis(decode_responses=True)

list_key = "example-list"

redis_connection.rpush(list_key, 1, 2, 3, 4, 5)
redis_connection.lpush(list_key, 7,8)

print(redis_connection.lrange(list_key, 0, -1))