require 'json'

module Worker
    class Enqueuer
        def initialize(redis, options = {})
            @redis = redis
            @scope = options
        end

        def ping!(request_id = nil)
            @job = Job.new('ping', requestId: request_id)

            @redis.lpush(self.job_key('ping'), @job)
        end

        def job_key(job_name)
            "#{@scope}:jobs:#{job_name}"
        end
    end

    class Job
        def initialize(name, args = {})
            @message = {
                id: SecureRandom.hex(12),
                name: name,
                t: Time.new.to_i,
                args: args
            }
        end

        def to_s
            JSON.generate(@message)
        end
    end
end
