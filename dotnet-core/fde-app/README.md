# Dotnet Core Sample App using a Framework-Dependent Executable

## Building

`pack build dotnet-fde-sample --buildpack paketo-buildpacks/dotnet-core`

## Running

`docker run --interactive --tty --env PORT=8080 --publish 8080:8080 dotnet-fde-sample`

## Viewing

Visit http://localhost:8080 in a browser.
