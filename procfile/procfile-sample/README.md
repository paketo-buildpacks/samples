# Procfile Static Webserver Sample Application

## Building

This is a simple go web server that is compiled at run time. It requires specific buildpacks, which are specified below. Since we are specifying the buildpacks at build time any builder can be used, but the ubuntu-noble-builder-buildpackless builder is used in this example.

### Buildpacks
* The [go-dist](https://github.com/paketo-buildpacks/go-dist) buildpack provides golang, but it needs another buildpack to require go in order to participate in the build.
* The [build-plan](https://github.com/paketo-community/build-plan) buildpack reads plan.toml and requires go at launch (runtime).
* The [procfile](https://github.com/paketo-buildpacks/procfile) buildpack sets the container entrypoint which calls `go run ...`.

```bash
pack build applications/procfile \
    --builder docker.io/paketobuildpacks/ubuntu-noble-builder-buildpackless \
    --buildpack docker.io/paketobuildpacks/go-dist \
    --buildpack docker.io/paketocommunity/build-plan \
    --buildpack docker.io/paketobuildpacks/procfile
```

## Running

```bash
docker run --tty --publish 8080:8080 applications/procfile
```

## Viewing

```bash
curl -s http://localhost:8080/hello-world.txt
```
