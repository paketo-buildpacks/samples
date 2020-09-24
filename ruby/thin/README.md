# Ruby sample app using Thin web server

## Building

`pack build thin-sample --buildpack gcr.io/paketo-community/ruby`

## Running

`docker run --interactive --tty --env PORT=8080 --publish 8080:8080 thin-sample`

## Viewing

`curl http://localhost:8080`
