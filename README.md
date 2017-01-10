# justhello
Trigger Go's background job by Rails Endpoint

## TODO

- [X] Init API and Worker structure
- [X] Init Rails API and prepare `/ping` endpoint
- [X] Docker Compose provision
- [X] Enqueue the job with PING message
- [X] Process the job in background with Go
- [X] Add simple UI for demo (bash scripts)
- [ ] Simple performce test script
- [X] Demo should described in README (Guide)
- [ ] Send demo project back to guys

## Guide

### Prerequirements

For launching you need:

- Unix like shell
- Docker
- Docker Compose

### Provisioning

For get all this stuff up and running run:

```
cd <~/path/to/justhello>
$ ./run.sh
```

This will download all required containers, build the images with apps,
and run in foreground next services:

- Redis server
- API (api/ folder)
- Worker (worker/ folder)

### Testing

After provisioning just open the separate shell and trigger the http://localhsot:3000 enpoints:

ping the worker

```
$ ./ping.sh
```


send message to the worker

```
$ ./hello.sh world
```

The Worker will print results into the stdout.

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

