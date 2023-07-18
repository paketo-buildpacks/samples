<i> Note: Node.js buildpack is designed to build backend applications that use Node as their server.
To build frontend applications and serve them with NGINX/HTTPD, please use the Web Servers buildpack </i>

# Node.js Sample app using Yarn and a React framework

## Building

`pack build react-sample --buildpack paketo-buildpacks/nodejs --env "BP_NODE_RUN_SCRIPTS=build"`

## Running

`docker run --interactive --tty --init --env PORT=8080 --publish 8080:8080 react-sample`

## Viewing

`curl http://localhost:8080`
