# Node.js Sample App using no package manager

## Building

### Ubuntu

#### Jammy

```
pack build nodejs-sample --buildpack docker.io/paketobuildpacks/nodejs \
    --builder docker.io/paketobuildpacks/builder-jammy-buildpackless-base
```

#### Noble

```
pack build nodejs-sample --buildpack docker.io/paketobuildpacks/nodejs \
    --builder docker.io/paketobuildpacks/ubuntu-noble-builder-buildpackless
```

#### Resolute

```
pack build nodejs-sample --buildpack docker.io/paketobuildpacks/nodejs \
    --builder docker.io/paketobuildpacks/ubuntu-resolute-builder-buildpackless
```

### RHEL

#### UBI 8

```
pack build nodejs-sample \
   --extension docker.io/paketobuildpacks/ubi-nodejs-extension \
   --buildpack docker.io/paketobuildpacks/nodejs \
   --builder docker.io/paketobuildpacks/builder-ubi8-buildpackless-base
```

#### UBI 9

```
pack build nodejs-sample \
   --builder docker.io/paketobuildpacks/ubi-9-builder
```

#### UBI 10

```
pack build nodejs-sample \
   --builder docker.io/paketobuildpacks/ubi-10-builder
```

## Running

`docker run --interactive --tty --init --publish 8080:8080 nodejs-sample`

## Viewing

`curl http://localhost:8080`
