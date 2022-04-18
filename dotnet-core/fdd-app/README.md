# Dotnet Core Sample App using a Framework-Dependent Deployment

## Building

`pack build dotnet-fdd-sample --buildpack paketo-buildpacks/dotnet-core`

## Running

`docker run --interactive --tty --env PORT=8080 --publish 8080:8080 dotnet-fdd-sample`

## Viewing

Visit http://localhost:8080 in a browser.
