# Node.js Sample App using no package manager

## Building

### Ubuntu

```
pack build nodejs-sample --buildpack docker.io/paketobuildpacks/nodejs \
    --builder docker.io/paketobuildpacks/builder-jammy-buildpackless-base
```

### RHEL

```
pack build nodejs-sample \
   --extension docker.io/paketobuildpacks/ubi-nodejs-extension \
   --buildpack docker.io/paketobuildpacks/nodejs \
   --builder docker.io/paketobuildpacks/builder-ubi8-buildpackless-base
```

## Running

`docker run --interactive --tty --init --publish 8080:8080 nodejs-sample`

## Viewing

`curl http://localhost:8080`
