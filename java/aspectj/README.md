# Java AspectJ Sample Application

## Building

```bash
pack build applications/aspectj
```

## Running

```bash
docker run --rm --tty --publish 8080:8080 applications/aspectj
```

## Viewing

```bash
curl -s http://localhost:8080/actuator/health | jq .
```
