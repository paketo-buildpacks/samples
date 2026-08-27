# Java Native Image Sample Application

## Building

### With `pack`

#### Spring Boot 3.1.x and later based app

```bash
pack build applications/native-image \
  --builder paketobuildpacks/ubuntu-resolute-builder \
  --run-image paketobuildpacks/ubuntu-resolute-run-tiny \
  --env BP_MAVEN_ACTIVE_PROFILES=native \
  --env BP_JVM_VERSION=25
```

> **Note:** `--run-image paketobuildpacks/ubuntu-resolute-run-tiny` is required here (and mirrored by `runImage` in `pom.xml`) because the Ubuntu Resolute builder's default run image is the full `run` variant, not `run-tiny`. The tiny run image is the recommended base for GraalVM native-image binaries, so this sample explicitly opts into it.

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
