# Java Dist-Zip Sample Application

## Building

```bash
pack build applications/dist-zip --env BP_GRADLE_BUILD_ARGUMENTS="--no-daemon -x test bootDistZip" --env BP_GRADLE_BUILT_ARTIFACT="build/distributions/*.zip"
```

## Running

```bash
docker run --rm --tty --publish 8080:8080 applications/dist-zip
```

## Viewing

```bash
curl -s http://localhost:8080/actuator/health | jq .
```
