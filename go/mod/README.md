# Go Sample App using Mod

## Building

`pack build mod-sample --buildpack paketo-buildpacks/go`

## Running

`docker run --interactive --tty --env PORT=8080 --publish 8080:8080 mod-sample`

## Viewing

`curl http://localhost:8080`
