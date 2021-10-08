# Clojure Deps Sample Application

## Building

```bash
pack build applications/clojure-deps
```

## Running

```bash
docker run --rm --tty --publish 8080:8080 -e JAVA_TOOL_OPTIONS="-XX:MaxMetaspaceSize=100M" applications/clojure-deps
```

## Viewing

```bash
curl -s http://localhost:8080/
```
