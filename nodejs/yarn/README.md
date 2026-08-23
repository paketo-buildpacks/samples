# Node.js Sample App using Yarn

## Building

### Ubuntu

#### Jammy

```
pack build yarn-sample --buildpack docker.io/paketobuildpacks/nodejs \
    --builder docker.io/paketobuildpacks/builder-jammy-buildpackless-base
```

#### Noble

```
pack build yarn-sample --buildpack docker.io/paketobuildpacks/nodejs \
    --builder docker.io/paketobuildpacks/ubuntu-noble-builder-buildpackless
```

#### Resolute

```
pack build yarn-sample --buildpack docker.io/paketobuildpacks/nodejs \
    --builder docker.io/paketobuildpacks/ubuntu-resolute-builder-buildpackless
```

### RHEL

#### UBI 8

```
pack build yarn-sample \
   --extension docker.io/paketobuildpacks/ubi-nodejs-extension \
   --buildpack docker.io/paketobuildpacks/nodejs \
   --builder docker.io/paketobuildpacks/builder-ubi8-buildpackless-base
```

#### UBI 9

```
pack build yarn-sample \
   --builder docker.io/paketobuildpacks/ubi-9-builder
```

#### UBI 10

```
pack build yarn-sample \
   --builder docker.io/paketobuildpacks/ubi-10-builder
```

## Running

`docker run --interactive --tty --publish 8080:8080 yarn-sample`

## Viewing

`curl http://localhost:8080`
