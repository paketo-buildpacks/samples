# Python sample app using poetry package manager and executable script

## Building

`pack build poetry-run-sample --buildpack paketo-buildpacks/python`

## Running

`docker run --interactive --tty --env PORT=8080 --publish 8080:8080 poetry-run-sample`

## Viewing

`curl http://localhost:8080`
