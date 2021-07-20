# Node.js Sample app using Yarn and a React framework

## Building

`pack build react-sample --buildpack gcr.io/paketo-buildpacks/nodejs --env "BP_NODE_RUN_SCRIPTS=build"`

## Running

`docker run --interactive --tty --init --env PORT=8080 --publish 8080:8080 react-sample`

## Viewing

`curl http://localhost:8080`
