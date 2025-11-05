# Node.js Sample App using Yarn

## Building

### Ubuntu

```
pack build yarn-sample --buildpack docker.io/paketobuildpacks/nodejs \
    --builder docker.io/paketobuildpacks/builder-jammy-buildpackless-base
```

### RHEL

```
pack build yarn-sample \
   --extension docker.io/paketobuildpacks/ubi-nodejs-extension \
   --buildpack docker.io/paketobuildpacks/nodejs \
   --builder docker.io/paketobuildpacks/builder-ubi8-buildpackless-base
```

## Running

`docker run --interactive --tty --publish 8080:8080 yarn-sample`

## Viewing

`curl http://localhost:8080`
