Rails.application.routes.draw do
    get 'ping', to: 'endpoints#ping'

  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
end
