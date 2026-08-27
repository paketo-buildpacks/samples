# Java Maven Sample Application

See [prerequisites](https://paketo.io/docs/howto/java/#prerequisites) of this sample.

## Building

```bash
pack build applications/maven
```

Alternatively, if you want to attach a Maven `settings.xml` file to pass additional configuration to Maven.

```bash
pack build applications/maven --volume $(pwd)/bindings:/platform/bindings
```

The command above will use the sample `settings.xml` file from this repo. It may be more useful to copy your local `settings.xml` first.

```bash
cp ~/.m2/settings.xml java/maven/bindings/maven/settings.xml
pack build applications/maven --volume $(pwd)/bindings:/platform/bindings
```

### Advanced

You can also select a specific builder by passing the `--builder` flag:

```bash
pack build applications/maven --builder paketobuildpacks/ubuntu-resolute-builder
```

## Running

```bash
docker run --rm --tty --publish 8080:8080 applications/maven
```

## Viewing

```bash
curl -s http://localhost:8080/actuator/health | jq .
```
