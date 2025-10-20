# Node.js Sample App using NPM

## Building

### Ubuntu

```
pack build npm-sample --buildpack docker.io/paketobuildpacks/nodejs \
    --builder docker.io/paketobuildpacks/builder-jammy-buildpackless-base
```

### RHEL

```
pack build npm-sample \
   --extension docker.io/paketobuildpacks/ubi-nodejs-extension \
   --buildpack docker.io/paketobuildpacks/nodejs \
   --builder docker.io/paketobuildpacks/builder-ubi8-buildpackless-base
```

## Running

`docker run --interactive --tty --init --publish 8080:8080 npm-sample`

## Viewing

`curl http://localhost:8080`
