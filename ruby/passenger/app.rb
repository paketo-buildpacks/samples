require 'sinatra/base'

class ExampleApp < Sinatra::Base
  get '/' do
    erb :index
  end
end
