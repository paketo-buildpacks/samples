# Java WAR Sample Application

## Building

```bash
pack build applications/war
```

## Running

```bash
docker run --rm --tty --publish 8080:8080 applications/war
```

## Viewing

```bash
curl -s http://localhost:8080/actuator/health | jq .
```
