# Paketo Buildpacks Sample Applications

A collection of sample applications that can be built using Paketo Buildpacks.

## Prerequisites

1. Clone this repository: `git clone https://github.com/paketo-buildpacks/samples`
1. [Pack](https://buildpacks.io/docs/install-pack/)

## Adding New Samples
* If the app is a part of an existing language family:
  * Add app to the appropriate language family in its own subdirectory.
  * Add a test context to the *_test.go file in the language family directory.
* If the app is a part of a new language family:
  * Create a new directory for the language family.
  * Create a new test file <language_family_name>/*_test.go containing a new
    test suite.
  * Be mindful of which builders the app is compatible with and set up test
    suites accordingly.
  * Run `./scripts/generate-test-workflow.sh -l <language_family_name>` to
    generate a Github Actions workflow that runs the tests.
* Update README.md.

## Samples

### Java
* [AspectJ](/java/aspectj)
* [Azure Application Insights](/java/application-insights)
* [Gradle DistZip](/java/dist-zip)
* [Gradle](/java/gradle)
* [Kotlin](/java/kotlin)
* [Leiningen (Clojure)](/java/leiningen)
* [Maven](/java/maven)
* [Native Image](/java/native-image)
* [Pre-compiled JAR](/java/jar)
* [Scala Akka](/java/akka)
* [WAR](/java/war)

### Node.js
* [No Package Manager](/nodejs/no-package-manager)
* [NPM](/nodejs/npm)
* [Yarn](/nodejs/yarn)

### Dotnet Core
* [Runtime-only](/dotnet-core/runtime)
* [ASPNet](/dotnet-core/aspnet)

### Go
* [No Imports](/go/no-imports)
* [Mod](/go/mod)

### PHP
* [Built-in Webserver](/php/builtin-server)
* [NGINX](/php/nginx)
* [Apache HTTPD](/php/httpd)
* [Composer](/php/composer)
* [Composer with Extensions](/php/composer_with_extensions)
* [Extensions via custon .ini](/php/app_with_extensions)
* [Memcached Session Handler](/php/memcached_session_handler)
* [Redis Session Handler](/php/redis_session_handler)

### Ruby
* [Passenger](/ruby/passenger)
* [Puma](/ruby/puma)
* [Rackup](/ruby/rackup)
* [Rake](/ruby/rake)
* [Thin](/ruby/thin)
* [Unicorn](/ruby/unicorn)

### Python
* [No Package Manager](/python/no_package_manager)
* [Pip](/python/pip)
* [Pipenv](/python/pipenv)
* [Miniconda](/python/conda)
* [Poetry (dependency management only)](/python/poetry)
* [Poetry (dependency management and run commands)](/python/poetry-run)

### Procfile
* [Static Webserver](/procfile)

### Web Servers
* [HTTPD](/web-servers/httpd-sample)
* [NGINX](/web-servers/nginx-sample)
* [Zero Configuration](/web-servers/no-config-file-sample)
* [JavaScript Frontend](/web-servers/javascript-frontend-sample)

## Testing Samples
To run integration tests that `pack build` each of the sample apps, use
`scripts/smoke.sh`. See `scripts/smoke.sh -h` for usage information.

For example, to run tests for the Go and .NET Core samples with the Paketo tiny
and base builders, run:
```
./smoke.sh --builder paketobuildpacks/builder:tiny \
           --builder paketobuildpacks/builder:base \
           --suite go \
           --suite dotnet-core
```
