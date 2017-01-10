# JustHello API

Based on very customized rails app. API mode

## Structure

```
- Redis Initialization (ENV: REDIS_URL for configure the Redis connection)
- Model WorkerEngueuer holds Worker Enqueue domain
```

## DDD

```
- Application JustHello
    - Context Worker
        - Domain Job
            - SubDomain Pong
            - SubDomain ProcessMessage

    - Context API
        - Domain Endpoints
        - Domain WorkerEnqueuer
            - SubDomain Ping
            - SubDomain EnqueueMessage
```
