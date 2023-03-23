# Java Azure Application Insights Sample Application

See [prerequisites](https://paketo.io/docs/howto/java/#prerequisites) of this sample.

## Binding

The buildpack installs the Azure Application Insights Agent, and configures it for usage based on a Service Binding.  This binding consists of a `type` indicating what type of service it is, and an `InstrumentationKey` with the Application Insight Instrumentation Key.

```plain
binding
├── type
└── InstrumentationKey
```

Add your instrumentation key to the binding

```bash
echo "<Instrumentation Key>" > binding/InstrumentationKey
```

## Building

```bash
pack build applications/application-insights --env BP_JVM_VERSION=11 --volume "$(pwd)/binding:/platform/bindings/application-insights"
```

## Running

```bash
docker run --rm --env SERVICE_BINDING_ROOT=/bindings --volume "$(pwd)/binding:/bindings/application" --tty --publish 8080:8080 applications/application-insights
```

## Viewing

```bash
curl -s http://localhost:8080/actuator/health | jq .
```
