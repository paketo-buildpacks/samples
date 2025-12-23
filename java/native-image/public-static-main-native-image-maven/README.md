# Java Native Image Sample Application, Basic

This is a basic Java app (i.e. public static void main) built using Native Image & Buildpacks.

## Building

### With `pack`

```bash
pack build applications/native-image \
  --builder paketobuildpacks/builder-jammy-tiny \
  --env BP_NATIVE_IMAGE=true \
  --env BP_JVM_VERSION=25
```

## Running

```bash
docker run --rm --tty --publish 8080:8080 applications/native-image
```

## Viewing

```bash
docker logs applications/native-image
```
