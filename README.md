# Paketo Buildpacks Sample Applications

A collection of sample applications that can be built using Paketo Buildpacks.

## Prerequisites

1. Clone this repository: `git clone https://github.com/paketo-buildpacks/samples`
1. [Pack](https://buildpacks.io/docs/install-pack/)

## Adding New Samples
* Add app to the appropriate language family
* Add a smoke test in the `tests` directory
  * If the app is a part of an existing language family: add a test context to
    the existing test file.
  * If the app is a part of a new language family: add a new test file, and add
    the suite to `tests/init_test.go`. Be mindful of which builders the app is
    compatible with.
* Update README.md

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
* [NPM](/nodejs/npm)
* [Yarn](/nodejs/yarn)

### Dotnet Core
* [Runtime-only](/dotnet-core/runtime)
* [ASPNet](/dotnet-core/aspnet)

### Go
* [Mod](/go/mod)
* [Dep](/go/dep)

### PHP
* [Built-in Webserver](/php/webserver)
* [NGINX](/php/nginx)
* [Apache HTTPD](/php/httpd)

### Ruby
* [Passenger](/ruby/passenger)
* [Puma](/ruby/puma)
* [Rackup](/ruby/rackup)
* [Rake](/ruby/rake)
* [Thin](/ruby/thin)
* [Unicorn](/ruby/unicorn)

### Procfile
* [Static Webserver](/procfile)
