# Node.js Sample App using NPM

## Building

`pack build npm-sample --buildpack gcr.io/paketo-buildpacks/nodejs`

## Running

`docker run --interactive --tty --init --publish 8080:8080 npm-sample`

## Viewing

`curl http://localhost:8080`
