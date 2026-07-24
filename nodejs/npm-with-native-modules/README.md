# Node.js Sample App using NPM

## Building

### Ubuntu

```
pack build npm-sample-with-native-module \
    --buildpack docker.io/paketobuildpacks/nodejs \
    --builder docker.io/paketobuildpacks/builder-jammy-buildpackless-base \
    --env BP_NPM_INCLUDE_BUILD_PYTHON=true
```

### RHEL

```
pack build npm-sample \
   --extension docker.io/paketobuildpacks/ubi-nodejs-extension \
   --buildpack docker.io/paketobuildpacks/nodejs \
   --builder docker.io/paketobuildpacks/builder-ubi8-buildpackless-base
```

## Running

`docker run --interactive --tty --init --publish 8080:8080 npm-sample-with-native-module`

## Viewing

`curl http://localhost:8080`
