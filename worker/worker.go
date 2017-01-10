package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gocraft/work"
	"os"
	"os/signal"
)

// Make a redis pool
var redisPool = &redis.Pool{
	MaxActive: 5,
	MaxIdle:   5,
	Wait:      true,
	Dial: func() (redis.Conn, error) {
		return redis.Dial("tcp", ":6379")
	},
}

type Context struct{}

func main() {
	// Make a new pool. Arguments:
	// Context{} is a struct that will be the context for the request.
	// 10 is the max concurrency
	// "JustHello" is the Redis namespace
	// redisPool is a Redis pool
	pool := work.NewWorkerPool(Context{}, 10, "JustHello", redisPool)

	// Add middleware that will be executed for each job
	pool.Middleware((*Context).Log)

	// Map the name of jobs to handler functions
	pool.Job("ping", (*Context).Pong)

	// Start processing jobs
	pool.Start()

	// Wait for a signal to quit:
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan

	// Stop the pool
	pool.Stop()
}

func (c *Context) Log(job *work.Job, next work.NextMiddlewareFunc) error {
	fmt.Println("Starting job: ", job.Name)
	return next()
}

func (c *Context) Pong(job *work.Job) error {
	if err := job.ArgError(); err != nil {
		return err
	}

	fmt.Println("PONG")

	return nil
}
