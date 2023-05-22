# Java Native Image Sample Application, using Gradle

## Building

### With `pack`

```bash
pack build applications/native-image \
  --builder paketobuildpacks/builder-jammy-tiny \
  --env BP_NATIVE_IMAGE=true
```

### With the Spring Boot Gradle Plugin

```bash
./gradlew bootBuildImage
```

## Running

```bash
docker run --rm --tty --publish 8080:8080 applications/native-image
```

## Viewing

```bash
curl -s http://localhost:8080/actuator/health | jq .
```
