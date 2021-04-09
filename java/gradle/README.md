# Java Gradle Sample Application

## Building

```bash
pack build applications/gradle
```

Alternatively, if you want to attach a `gradle.properties` file to pass additional configuration to Gradle.

```bash
pack build --volume $(pwd)/bindings:/platform/bindings applications/gradle
```

## Running

```bash
docker run --rm --tty --publish 8080:8080 applications/gradle
```

## Viewing

```bash
curl -s http://localhost:8080/actuator/health | jq .
```
