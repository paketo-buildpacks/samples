# Java Native Image Sample Application

## Building

### With `pack`

```bash
pack build applications/native-image \
  --builder paketobuildpacks/builder-jammy-tiny \
  --env BP_NATIVE_IMAGE=true \
  --env BP_MAVEN_ACTIVE_PROFILES="native"
```

### With the Spring Boot Maven Plugin

```bash
./mvnw -Dmaven.test.skip=true spring-boot:build-image -Pnative
```

## Running

```bash
docker run --rm --tty --publish 8080:8080 applications/native-image
```

## Viewing

```bash
curl -s http://localhost:8080/actuator/health | jq .
```
