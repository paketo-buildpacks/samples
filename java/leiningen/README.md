# Clojure Leiningen Sample Application

See [prerequisites](https://paketo.io/docs/howto/java/#prerequisites) of this sample.

## Building

```bash
pack build applications/leiningen  --env BP_JVM_VERSION=11
```

## Running

```bash
docker run --rm --tty --publish 8080:8080 -e JAVA_TOOL_OPTIONS="-XX:MaxMetaspaceSize=100M" applications/leiningen
```

## Viewing

```bash
curl -s http://localhost:8080/
```
