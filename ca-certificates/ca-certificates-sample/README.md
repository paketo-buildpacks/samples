# CA Certificates Sample Application

## Binding

The Paketo CA Certificates Buildpack adds CA Certificates to the system truststore and build and or runtime. Certificates must be provided with a binding of `type` `ca-certificates`. Each certificate in the binding should contain exactly one PEM encoded CA certificate.

## Preparing the Binding
Certificates must be provided to the buildpack or running application with a binding of `type` `ca-certificates`.

This sample contains a skeleton binding
```plain
binding
└── type
```

Given a file `<your-ca.pem>` containing a single PEM encoded CA certificate needed to verify a TLS connection to an https URL `<url>`, add the CA certificate to the binding.
```bash
cp <your-ca.pem> binding/your-ca.pem
```

## Verifying Connections at Runtime
This sample contains a simple Golang application that will make a `HEAD` request to a provided URL.

Build the sample application
```bash
pack build applications/ca-certificates \
    --buildpack paketo-buildpacks/go
```
Run the sample application **without** the binding, passing the URL as a positional argument, to observe the expected error (should print `ERROR: Head "<url>": x509: certificate signed by unknown authority`):
```bash
docker run --rm \
  applications/ca-certificates <url>
```

Rerun the sample application **with** the binding, to verify the connection (should print `SUCCESS!`).

```bash
docker run --rm \
  --env SERVICE_BINDING_ROOT=/bindings \
  --volume "$(pwd)/binding:/bindings/ca-certificates" \
  applications/ca-certificates <url>
```


## Verifying Connections at Build-Time
This sample contains a simple buildpack that will use curl to make a `HEAD` request to provided `$BP_TEST_URL` at build-time.

Build the sample application with the sample buildpack and set `$BP_TEST_URL` to your URL in the build environment.
```bash
pack build applications/ca-certificates \
    --builder paketobuildpacks/builder-jammy-full:latest \
    --buildpack ./buildpack \
    --buildpack paketo-buildpacks/go \
    --volume "$(pwd)/binding:/platform/bindings/ca-certificates" \
    --env BP_TEST_URL=<url>
```

The builder `paketobuildpacks/builder-jammy-full:latest` is required in the above command because the full stack provides `curl`.
