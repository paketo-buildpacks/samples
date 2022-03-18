# Ruby sample app using Unicorn web server

## Building

`pack build unicorn-sample --buildpack paketo-buildpacks/ruby`

## Running

`docker run --interactive --tty --env PORT=8080 --publish 8080:8080 unicorn-sample`

## Viewing

`curl http://localhost:8080`
