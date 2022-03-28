# Python sample app using poetry package manager

## Building

`pack build poetry-sample --buildpack paketo-buildpacks/python`

## Running

`docker run --interactive --tty --env PORT=8080 --publish 8080:8080 poetry-sample`

## Viewing

`curl http://localhost:8080`
