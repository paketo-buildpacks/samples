# Python sample app using pip package manager

## Building

`pack build pip-sample --buildpack gcr.io/paketo-buildpacks/python`

## Running

`docker run --interactive --tty --env PORT=8080 --publish 8080:8080 pip-sample`

## Viewing

`curl http://localhost:8080`
