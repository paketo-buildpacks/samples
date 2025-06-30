# Quarkus Native Sample Application

## Building

### With `pack`

```bash
pack build applications/quarkus-native \
  --builder paketobuildpacks/builder-jammy-tiny \
  --env BP_NATIVE_IMAGE=true \
  --env BP_MAVEN_ADDITIONAL_BUILD_ARGUMENTS="-Dquarkus.package.type=native-sources" \
  --env BP_MAVEN_BUILT_ARTIFACT="target/native-sources" \
  --env BP_NATIVE_IMAGE_BUILD_ARGUMENTS_FILE="native-sources/native-image.args" \
  --env BP_NATIVE_IMAGE_BUILT_ARTIFACT="native-sources/*-runner.jar" \
  --env BP_JVM_VERSION=21
```

## Running

```bash
docker run --rm --tty --publish 8080:8080 applications/quarkus-native
```

## Viewing

```bash
curl -s http://localhost:8080/hello
```

or

```bash
curl -s http://localhost:8080/hello/greeting/$(whoami)
```
