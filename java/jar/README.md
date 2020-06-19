# Pre-compiled Java Sample Application

## Building

```bash
pack build applications/jar
```

## Running

```bash
docker run --tty --publish 8080:8080 applications/jar
```

## Viewing

```bash
curl -s http://localhost:8080/actuator/health | jq .
```
