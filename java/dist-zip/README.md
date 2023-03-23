# Java Dist-Zip Sample Application

See [prerequisites](https://paketo.io/docs/howto/java/#prerequisites) of this sample.

## Building

```bash
pack build applications/dist-zip --env BP_GRADLE_BUILD_ARGUMENTS="--no-daemon -x test bootDistZip" --env BP_GRADLE_BUILT_ARTIFACT="build/distributions/*.zip" --env BP_JVM_VERSION=17
```

## Running

```bash
docker run --rm --tty --publish 8080:8080 applications/dist-zip
```

## Viewing

```bash
curl -s http://localhost:8080/actuator/health | jq .
```
