package main

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gocraft/work"
	"log"
	"os"
	"os/signal"
)

// Make a redis pool
var redisPool = &redis.Pool{
	MaxActive: 5,
	MaxIdle:   5,
	Wait:      true,
	Dial: func() (redis.Conn, error) {
		return redis.DialURL(os.Getenv("REDIS_URL"))
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
	pool.Job("message", (*Context).ProcessMessage)

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
	log.Println("Starting job: ", job.Name)
	return next()
}

func (c *Context) Pong(job *work.Job) error {
	requestId := job.ArgString("requestId")

	// not sure what is that for..
	if err := job.ArgError(); err != nil {
		return err
	}

	log.Println("PONG requestId:", requestId)

	return nil
}

func (c *Context) ProcessMessage(job *work.Job) error {
	message := job.ArgString("message")

	log.Println("GOT MESSAGE:", message)

	return nil
}
