# Python sample app using conda package manager

## Building

`pack build conda-sample --buildpack gcr.io/paketo-buildpacks/python`

## Running

`docker run --interactive --tty --env PORT=8080 --publish 8080:8080 conda-sample`

## Viewing

`curl http://localhost:8080`
