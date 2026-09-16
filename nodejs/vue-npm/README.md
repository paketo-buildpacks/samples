<i> Note: Node.js buildpack is designed to build backend applications that use Node as their server.
To build frontend applications and serve them with NGINX/HTTPD, please use the Web Servers buildpack </i>

# Node.js Sample app using Npm and a Vue framework

## Building

### Ubuntu

#### Jammy

```
pack build vue-sample --buildpack docker.io/paketobuildpacks/nodejs \
    --builder docker.io/paketobuildpacks/builder-jammy-buildpackless-base \
    --env "BP_NODE_RUN_SCRIPTS=build" --env "NODE_ENV=development"
```

#### Noble

```
pack build vue-sample --buildpack docker.io/paketobuildpacks/nodejs \
    --builder docker.io/paketobuildpacks/ubuntu-noble-builder-buildpackless \
    --env "BP_NODE_RUN_SCRIPTS=build" --env "NODE_ENV=development"
```

#### Resolute

```
pack build vue-sample --buildpack docker.io/paketobuildpacks/nodejs \
    --builder docker.io/paketobuildpacks/ubuntu-resolute-builder-buildpackless \
    --env "BP_NODE_RUN_SCRIPTS=build" --env "NODE_ENV=development"
```

### RHEL

#### UBI 8

```
pack build vue-sample\
    --extension docker.io/paketobuildpacks/ubi-nodejs-extension \
    --buildpack docker.io/paketobuildpacks/nodejs \
    --builder docker.io/paketobuildpacks/builder-ubi8-buildpackless-base \
    --env "BP_NODE_RUN_SCRIPTS=build" --env "NODE_ENV=development"
```

#### UBI 9

```
pack build vue-sample \
    --builder docker.io/paketobuildpacks/ubi-9-builder \
    --env "BP_NODE_RUN_SCRIPTS=build" --env "NODE_ENV=development"
```

#### UBI 10

```
pack build vue-sample \
    --builder docker.io/paketobuildpacks/ubi-10-builder \
    --env "BP_NODE_RUN_SCRIPTS=build" --env "NODE_ENV=development"
```

## Running

`docker run --interactive --tty --init --publish 8080:8080 vue-sample`

## Viewing

`curl http://localhost:8080`

### Note

We need the additional flag `--env "NODE_ENV=development"` when running `pack build` since we need the `vue-cli-service` provided in the devDependencies.
