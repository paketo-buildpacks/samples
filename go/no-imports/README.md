# Go Sample App using no imports

## Building

`pack build go-sample --buildpack paketo-buildpacks/go`

## Running

`docker run --interactive --tty --env PORT=8080 --publish 8080:8080 go-sample`

## Viewing

`curl http://localhost:8080`
