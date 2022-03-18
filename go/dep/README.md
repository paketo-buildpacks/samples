# Go Sample App using Dep

## Building

`pack build dep-sample --buildpack paketo-buildpacks/go`

## Running

`docker run --interactive --tty --env PORT=8080 --publish 8080:8080 dep-sample`

## Viewing

`curl http://localhost:8080`
