# Java Native Image Sample Application, Basic

This is a basic Java app (i.e. public static void main) built using Native Image & Buildpacks.

## Building

### With `pack`

```bash
pack build applications/native-image \
  --builder paketobuildpacks/ubuntu-resolute-builder \
  --run-image paketobuildpacks/ubuntu-resolute-run-tiny \
  --env BP_NATIVE_IMAGE=true \
  --env BP_JVM_VERSION=25
```

> **Note:** `--run-image paketobuildpacks/ubuntu-resolute-run-tiny` is required here because the Ubuntu Resolute builder's default run image is the full `run` variant, not `run-tiny`. The tiny run image is the recommended base for GraalVM native-image binaries, so this sample explicitly opts into it.

## Running

```bash
docker run --rm --tty --publish 8080:8080 applications/native-image
```

## Viewing

```bash
docker logs applications/native-image
```
