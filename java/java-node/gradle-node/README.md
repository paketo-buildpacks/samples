# Java Gradle Sample Application

See [prerequisites](https://paketo.io/docs/howto/java/#prerequisites) of this sample.

## Building

```bash
pack build applications/gradle-node --env BP_JAVA_INSTALL_NODE=true --env BP_NODE_PROJECT_PATH=frontend
```

### Advanced

You can also select a specific builder and control the JVM version by passing additional arguments:

```bash
pack build applications/gradle-node --builder paketobuildpacks/ubuntu-resolute-builder --env BP_JVM_VERSION=21 --env BP_JAVA_INSTALL_NODE=true --env BP_NODE_PROJECT_PATH=frontend
```

## Running

```bash
docker run --rm --tty --publish 8080:8080 applications/gradle-node
```

## Viewing

Open your web browser at: http://localhost:8080/; you should see a page generated using JavaScript code getting data from a Spring Boot Flux Rest Controller
