# Node.js Sample App using no package manager

## Building

`pack build nodejs-sample --buildpack paketo-buildpacks/nodejs`

## Running

`docker run --interactive --tty --init --publish 8080:8080 nodejs-sample`

## Viewing

`curl http://localhost:8080`
