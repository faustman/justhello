# justhello
Trigger Go's background job by Rails Endpoint

## TODO

- [ ] Init API and Worker structure
- [ ] Init Rails API and prepare `/hello` endpoint
- [ ] Docker Compose provision
- [ ] Enqueue the job with PING message
- [ ] Process the job in background with Go
- [ ] Add simple UI for demo
- [ ] Simple performce test script
- [ ] Demo should described in README
- [ ] Send demo project back to guys

## Guide

> TODO

### Redis

launch redis-cli:

```
$ docker-compose exec redis redis-cli

 127.0.0.1:6379> KEYS *
 1) "my_app_namespace:known_jobs"
 2) "my_app_namespace:jobs:send_email"
 127.0.0.1:6379> TYPE my_app_namespace:jobs:send_email
 list
 127.0.0.1:6379> LPOP my_app_namespace:jobs:send_email
 "{\"name\":\"send_email\",\"id\":\"9f5762f3b84ad82130fb2237\",\"t\":1484004277,\"args\":{\"address\":\"test@example.com\",\"customer_id\":4,\"subject\":\"hello world\"}}"
```

