require 'test_helper'

class EndpointsControllerTest < ActionDispatch::IntegrationTest
  test "should get ping" do
    get ping_url

    assert_response :success
    assert_equal "PONG", @response.body
  end

end
