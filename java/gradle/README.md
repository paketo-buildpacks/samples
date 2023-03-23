# Java Gradle Sample Application

See [prerequisites](https://paketo.io/docs/howto/java/#prerequisites) of this sample.

## Building

```bash
pack build applications/gradle --env BP_JVM_VERSION=17
```

Alternatively, if you want to attach a `gradle.properties` file to pass additional configuration to Gradle.

```bash
pack build applications/gradle --volume $(pwd)/bindings:/platform/bindings --env BP_JVM_VERSION=17
```

The command above will use the sample `gradle.properties` file from this repo. It may be more useful to copy your local `gradle.properties` file first.

```bash
cp ~/.gradle/gradle.properties java/gradle/bindings/gradle/gradle.properties
pack build applications/gradle --volume $(pwd)/bindings:/platform/bindings  --env BP_JVM_VERSION=17
```

## Running

```bash
docker run --rm --tty --publish 8080:8080 applications/gradle
```

## Viewing

```bash
curl -s http://localhost:8080/actuator/health | jq .
```
