<i> Note: Node.js buildpack is designed to build backend applications that use Node as their server.
To build frontend applications and serve them with NGINX/HTTPD, please use the Web Servers buildpack </i>

# Node.js Sample app using Npm and a Vue framework

## Building

`pack build vue-sample --buildpack paketo-buildpacks/nodejs --env "BP_NODE_RUN_SCRIPTS=build" --env "NODE_ENV=development"`

## Running

`docker run --interactive --tty --init --publish 8080:8080 vue-sample`

## Viewing

`curl http://localhost:8080`

### Note

We need the additional flag `--env "NODE_ENV=development"` when running `pack build` since we need the `vue-cli-service` provided in the devDependencies.
