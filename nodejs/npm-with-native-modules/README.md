# Node.js Sample App using NPM

## Building

### Ubuntu

#### Jammy

```
pack build npm-sample-with-native-module \
    --buildpack docker.io/paketobuildpacks/nodejs \
    --builder docker.io/paketobuildpacks/builder-jammy-buildpackless-base \
    --env BP_NPM_INCLUDE_BUILD_PYTHON=true
```

#### Noble

```
pack build npm-sample-with-native-module \
    --buildpack docker.io/paketobuildpacks/nodejs \
    --builder docker.io/paketobuildpacks/ubuntu-noble-builder-buildpackless \
    --env BP_NPM_INCLUDE_BUILD_PYTHON=true
```

#### Resolute

```
pack build npm-sample-with-native-module \
    --buildpack docker.io/paketobuildpacks/nodejs \
    --builder docker.io/paketobuildpacks/ubuntu-resolute-builder-buildpackless \
    --env BP_NPM_INCLUDE_BUILD_PYTHON=true
```

### RHEL

#### UBI 8

```
pack build npm-sample \
   --extension docker.io/paketobuildpacks/ubi-nodejs-extension \
   --buildpack docker.io/paketobuildpacks/nodejs \
   --builder docker.io/paketobuildpacks/builder-ubi8-buildpackless-base
```

#### UBI 9

```
pack build npm-sample-with-native-module \
   --builder docker.io/paketobuildpacks/ubi-9-builder \
   --env BP_NPM_INCLUDE_BUILD_PYTHON=true
```

#### UBI 10

```
pack build npm-sample-with-native-module \
   --builder docker.io/paketobuildpacks/ubi-10-builder \
   --env BP_NPM_INCLUDE_BUILD_PYTHON=true
```

## Running

`docker run --interactive --tty --init --publish 8080:8080 npm-sample-with-native-module`

## Viewing

`curl http://localhost:8080`
