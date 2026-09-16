<i> Note: Node.js buildpack is designed to build backend applications that use Node as their server.
To build frontend applications and serve them with NGINX/HTTPD, please use the Web Servers buildpack </i>

# Node.js Sample app using Yarn and a React framework

## Building

### Ubuntu

#### Jammy

```
pack build react-sample --buildpack docker.io/paketobuildpacks/nodejs \
    --builder docker.io/paketobuildpacks/builder-jammy-buildpackless-base \
    --env "BP_NODE_RUN_SCRIPTS=build"
```

#### Noble

```
pack build react-sample --buildpack docker.io/paketobuildpacks/nodejs \
    --builder docker.io/paketobuildpacks/ubuntu-noble-builder-buildpackless \
    --env "BP_NODE_RUN_SCRIPTS=build"
```

#### Resolute

```
pack build react-sample --buildpack docker.io/paketobuildpacks/nodejs \
    --builder docker.io/paketobuildpacks/ubuntu-resolute-builder-buildpackless \
    --env "BP_NODE_RUN_SCRIPTS=build"
```

### RHEL

#### UBI 8

```
pack build react-sample \
    --extension docker.io/paketobuildpacks/ubi-nodejs-extension \
    --buildpack docker.io/paketobuildpacks/nodejs \
    --builder docker.io/paketobuildpacks/builder-ubi8-buildpackless-base \
    --env "BP_NODE_RUN_SCRIPTS=build"
```

#### UBI 9

```
pack build react-sample \
    --builder docker.io/paketobuildpacks/ubi-9-builder \
    --env "BP_NODE_RUN_SCRIPTS=build"
```

#### UBI 10

```
pack build react-sample \
    --builder docker.io/paketobuildpacks/ubi-10-builder \
    --env "BP_NODE_RUN_SCRIPTS=build"
```

## Running

`docker run --interactive --tty --init --env PORT=8080 --publish 8080:8080 react-sample`

## Viewing

`curl http://localhost:8080`
