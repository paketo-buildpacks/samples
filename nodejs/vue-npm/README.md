# Node.js Sample app using Npm and a Vue framework

## Building

`pack build vue-sample --buildpack gcr.io/paketo-buildpacks/nodejs --env "BP_NODE_RUN_SCRIPTS=build" --env "NODE_ENV=development"`

## Running

`docker run --interactive --tty --init --publish 8080:8080 vue-sample`

## Viewing

`curl http://localhost:8080`

### Note

We need the additional flag `--env "NODE_ENV=development"` when running `pack build` since we need the `vue-cli-service` provided in the devDependencies.
