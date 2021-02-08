# Clojure Leiningen Sample Application

## Building

```bash
pack build --env JAVA_TOOL_OPTIONS='-XX:MaxMetaspaceSize=100M' applications/leiningen
```

## Running

```bash
docker run --rm --tty --publish 8080:8080 --env JAVA_TOOL_OPTIONS='-XX:MaxMetaspaceSize=100M' applications/leiningen
```

## Viewing

```bash
curl -s http://localhost:8080/
```
