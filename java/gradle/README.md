# Java Gradle Sample Application

## Building

```bash
pack build applications/gradle
```

Alternatively, if you want to attach a `gradle.properties` file to pass additional configuration to Gradle.

```bash
pack build --volume $(pwd)/bindings:/platform/bindings applications/gradle
```

The command above will use the sample `gradle.properties` file from this repo. It may be more useful to copy your local `gradle.properties` file first.

```bash
cp ~/.gradle/gradle.properties java/gradle/bindings/gradle/gradle.properties
pack build applications/gradle --volume $(pwd)/bindings:/platform/bindings
```

## Running

```bash
docker run --rm --tty --publish 8080:8080 applications/gradle
```

## Viewing

```bash
curl -s http://localhost:8080/actuator/health | jq .
```
