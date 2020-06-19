# Java Azure Application Insights Sample Application

## Binding

The buildpack installs the Azure Application Insights Agent, and configures it for usage based on a Service Binding.  This binding consists of a `kind` indicating what type of service it is, and an `InstrumentationKey` with the Application Insight Instrumentation Key.

```plain
binding
├── metadata
│   ├── kind
│   └── provider
└── secret
    └── InstrumentationKey
```

Add your instrumentation key to the binding

```bash
echo "<Instrumentation Key>" > binding/secret/InstrumentationKey
```

## Building

```bash
pack build applications/application-insights --volume "$(pwd)/binding:/bindings/application-insights"
```

## Running

```bash
docker run --tty --publish 8080:8080 applications/application-insights
```

## Viewing

```bash
curl -s http://localhost:8080/actuator/health | jq .
```
