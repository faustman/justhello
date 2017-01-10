require 'redis'

class EndpointsController < ApplicationController
  def ping
      @redis = Redis.new
      @worker = Worker::Enqueuer.new(@redis, "JustHello")

      @worker.ping!

      render plain: 'PONG'
  end
end
