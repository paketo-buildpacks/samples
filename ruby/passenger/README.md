# Ruby sample app using Passenger web server

## Building

`pack build passenger-sample --buildpack gcr.io/paketo-community/ruby`

## Running

`docker run --interactive --tty --env PORT=8080 --publish 8080:8080 passenger-sample`

## Viewing

`curl http://localhost:8080`
