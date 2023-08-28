FROM debian:stable-20230522

RUN apt-get update && apt install -y lsb-release curl gpg

RUN curl -fsSL https://packages.redis.io/gpg |  gpg --dearmor -o /usr/share/keyrings/redis-archive-keyring.gpg

RUN echo "deb [signed-by=/usr/share/keyrings/redis-archive-keyring.gpg] https://packages.redis.io/deb $(lsb_release -cs) main" |  tee /etc/apt/sources.list.d/redis.list

RUN apt-get update && apt-get install -y redis

COPY redis.conf /usr/local/etc/redis/redis.conf

CMD ["redis-server","/usr/local/etc/redis/redis.conf"]