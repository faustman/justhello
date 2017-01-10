package main

import (
	"github.com/gocraft/work"
	"testing"
)

func TestPongJob(t *testing.T) {
	job := &work.Job{
		Name:       "ping",
		ID:         "1231241",
		EnqueuedAt: 1231241,
	}

	context := &Context{}

	context.Pong(job)

	// TODO: check the output message
}
