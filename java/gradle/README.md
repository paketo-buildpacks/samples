# Java Gradle Sample Application

## Building

```bash
pack build applications/gradle
```

## Running

```bash
docker run --tty --publish 8080:8080 applications/gradle
```

## Viewing

```bash
curl -s http://localhost:8080/actuator/health | jq .
```
