# Node.js Sample App using NPM

## Building

### Ubuntu

#### Jammy

```
pack build npm-sample --buildpack docker.io/paketobuildpacks/nodejs \
    --builder docker.io/paketobuildpacks/builder-jammy-buildpackless-base
```

#### Noble

```
pack build npm-sample --buildpack docker.io/paketobuildpacks/nodejs \
    --builder docker.io/paketobuildpacks/ubuntu-noble-builder-buildpackless
```

#### Resolute

```
pack build npm-sample --buildpack docker.io/paketobuildpacks/nodejs \
    --builder docker.io/paketobuildpacks/ubuntu-resolute-builder-buildpackless
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
pack build npm-sample \
   --builder docker.io/paketobuildpacks/ubi-9-builder
```

#### UBI 10

```
pack build npm-sample \
   --builder docker.io/paketobuildpacks/ubi-10-builder
```

## Running

`docker run --interactive --tty --init --publish 8080:8080 npm-sample`

## Viewing

`curl http://localhost:8080`
