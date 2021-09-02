# Python sample app using pipenv package manager

## Building

`pack build pipenv-sample --buildpack gcr.io/paketo-buildpacks/python`

## Running

`docker run --interactive --tty --env PORT=8080 --publish 8080:8080 pipenv-sample`

## Viewing

`curl http://localhost:8080`
