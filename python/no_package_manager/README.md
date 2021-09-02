# Python sample app using no package manager

## Building

`pack build no-package-manager-sample --buildpack gcr.io/paketo-buildpacks/python`

## Running

`docker run --interactive --tty --env PORT=8080 --publish 8080:8080 no-package-manager-sample`

## Viewing

`curl http://localhost:8080`
