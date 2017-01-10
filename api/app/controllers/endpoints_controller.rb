require 'redis'

class EndpointsController < ApplicationController
  def ping
      worker.ping! request.uuid

      render plain: 'PONG'
  end

  def message
      worker.enqueue message_param

      render plain: 'OK'
  end

  private

  def worker
      @worker ||= Worker::Enqueuer.new(Redis.new, "JustHello")
  end

  def message_param
      params.require(:message)
  end
end
