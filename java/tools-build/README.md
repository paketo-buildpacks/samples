# Clojure Tools Build Sample Application

See [prerequisites](https://paketo.io/docs/howto/java/#prerequisites) of this sample.

## Building

```bash
pack build applications/clojure-tools-build --builder paketobuildpacks/ubuntu-resolute-builder
```

## Running

```bash
docker run --rm --tty --publish 8080:8080 -e JAVA_TOOL_OPTIONS="-XX:MaxMetaspaceSize=100M" applications/clojure-tools-build
```

## Viewing

```bash
curl -s http://localhost:8080/
```
