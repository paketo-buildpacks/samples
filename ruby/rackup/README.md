# Ruby sample app using Rackup

## Building

`pack build rackup-sample --buildpack paketo-buildpacks/ruby`

## Running

`docker run --interactive --tty --env PORT=8080 --publish 8080:8080 rackup-sample`

## Viewing

`curl http://localhost:8080`
