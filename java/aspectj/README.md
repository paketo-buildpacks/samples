# Java AspectJ Sample Application

See [prerequisites](https://paketo.io/docs/howto/java/#prerequisites) of this sample.

## Building

```bash
pack build applications/aspectj --env BP_JVM_VERSION=11
```

## Running

```bash
docker run --rm --tty --publish 8080:8080 applications/aspectj
```

## Viewing

```bash
curl -s http://localhost:8080/actuator/health | jq .
```
