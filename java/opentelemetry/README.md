# Java OpenTelemetry Sample Application

See [prerequisites](https://paketo.io/docs/howto/java/#prerequisites) of this sample.

## Building

```bash
pack build applications/opentelemetry \
    --buildpack paketo-buildpacks/java \
    --buildpack docker.io/paketobuildpacks/opentelemetry \
    --env BP_OPENTELEMETRY_ENABLED=true \
    --env BP_JVM_VERSION=17
```

## Running

```bash
docker run --rm -p 8080:8080 -e OTEL_JAVAAGENT_ENABLED=true applications/opentelemetry
```

You can configure the OpenTelemetry agent at run-time as described in the [project documentation](https://opentelemetry.io/docs/instrumentation/java/automatic/agent-config/).

```bash
docker run --rm -p 8080:8080 \
    -e OTEL_JAVAAGENT_ENABLED=true \
    -e OTEL_SERVICE_NAME=myapp \
    applications/opentelemetry
```

Alternatively, you can mount the OpenTelemetry configuration as a config tree via a binding. That is useful, for instance, when the configuration is provided as key/value pairs in a Kubernetes Secret object.

```bash
docker run --rm -p 8080:8080 \
  -v $(pwd)/bindings:/bindings \
  -e SERVICE_BINDING_ROOT=/bindings \
  -e OTEL_JAVAAGENT_ENABLED=true \
  applications/opentelemetry
```

## Viewing

```bash
curl -s http://localhost:8080/config | jq .
```

The result will be:

```bash
[
  "OTEL_JAVAAGENT_ENABLED=true",
  "OTEL_SERVICE_NAME=null",
  "OTEL_LOGS_EXPORTER=none",
  "OTEL_METRICS_EXPORTER=none"
]
```

If you configured the `OTEL_SERVICE_NAME` property (either via environment variable or volume binding), the output will be:

```bash
[
  "OTEL_JAVAAGENT_ENABLED=true",
  "OTEL_SERVICE_NAME=myapp",
  "OTEL_LOGS_EXPORTER=none",
  "OTEL_METRICS_EXPORTER=none"
]
```
