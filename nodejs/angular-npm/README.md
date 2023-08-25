<i> Note: Node.js buildpack is designed to build backend applications that use Node as their server.
To build frontend applications and serve them with NGINX/HTTPD, please use the Web Servers buildpack </i>

# Node.js Sample app using Npm and a Angular framework

## Building

`pack build angular-sample --buildpack paketo-buildpacks/nodejs --env "BP_NODE_RUN_SCRIPTS=build" --env "NODE_ENV=development"`

## Running

`docker run --interactive --tty --init --env PORT=8080 --publish 8080:8080 angular-sample`

## Viewing

`curl http://localhost:8080`

### Note

When running containerized Angular applications, be sure to specify the host before running `ng serve` in the `package.json`. For examples:

`ng serve --host 0.0.0.0`

We need the additional flag `--env "NODE_ENV=development"` when running `pack build` since we need the `ng` cli provided in the devDependencies.

