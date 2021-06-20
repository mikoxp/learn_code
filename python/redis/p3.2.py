from redis import Redis

redis_connection = Redis(decode_responses=True)

list_key = "example-list"
# waiting to message
while True:
    print(redis_connection.brpop(list_key))