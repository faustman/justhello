require "test_helper"

require "redis"
require "json"

SCOPE = "JustHelloTest"

class Worker::EnqueuerTest < ActiveSupport::TestCase
  setup do
    @redis = Redis.new

    @worker = Worker::Enqueuer.new(@redis, SCOPE)
  end

  test ".job_key" do
    job_name = "someJobName"

    assert_equal "#{SCOPE}:jobs:#{job_name}", @worker.job_key(job_name)
  end

  test "should enqueue ping message to the Redis List" do

    @worker.ping!

    key = @worker.job_key("ping")

    last_message = JSON.parse(@redis.lpop(key))

    assert_not_nil last_message, "Message should exists"

    assert_not_nil last_message["id"]
    assert_not_nil last_message["name"]
    assert_not_nil last_message["t"]
    assert_not_nil last_message["args"]

    assert_equal "ping", last_message["name"]

  end
end
