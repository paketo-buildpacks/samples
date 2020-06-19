# Java Maven Sample Application

## Building

```bash
pack build applications/maven
```

## Running

```bash
docker run --tty --publish 8080:8080 applications/maven
```

## Viewing

```bash
curl -s http://localhost:8080/actuator/health | jq .
```
