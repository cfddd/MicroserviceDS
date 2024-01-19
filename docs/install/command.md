## mysql

```shell
docker build -t my-mysql .
docker run -d -p 3306:3306 --name my-mysql -v ~/mysql/data:/var/lib/mysql my-mysql

```
## redis
```shell
docker pull redis 

docker run --restart=always \
-p 6379:6379 \
--name my-redis \
-v ~/redis/redis.conf:/etc/redis/redis.conf \
-v ~/redis/data:/data \
-d redis redis-server /etc/redis/redis.conf

```
## rabbitMQ
```shell

docker pull rabbitmq:3-management

docker run -d \
--name my-rabbitmq \
-p 5672:5672 -p 15672:15672 \
-e RABBITMQ_DEFAULT_USER=guest \
-e RABBITMQ_DEFAULT_PASS=guest \
-e RABBITMQ_DEFAULT_VHOST=/ rabbitmq:3-management
```