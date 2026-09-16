# Kotlin Sample Application

See [prerequisites](https://paketo.io/docs/howto/java/#prerequisites) of this sample.

## Building

```bash
pack build applications/kotlin
```

### Advanced

You can also select a specific builder and control the JVM version by passing additional arguments:

```bash
pack build applications/kotlin --builder paketobuildpacks/ubuntu-resolute-builder --env BP_JVM_VERSION=17
```

## Running

```bash
docker run --rm --tty --publish 8080:8080 applications/kotlin
```

## Viewing

```bash
curl -s http://localhost:8080/actuator/health | jq .
```
