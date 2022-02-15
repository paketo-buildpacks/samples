# Ruby sample app using Rails Asset precompilation 

## Building

`pack build rails-sample --buildpack gcr.io/paketo-buildpacks/ruby`

## Running

`docker run --interactive --tty --env SECRET_KEY_BASE="some-secret" --publish 9292:9292 rails-sample`

## Viewing

`curl http://localhost:9292`
