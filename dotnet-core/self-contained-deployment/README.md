# Dotnet Core Sample App using a Self-Contained Deployment

## Building

`pack build dotnet-scd-sample --buildpack paketo-buildpacks/dotnet-core`

## Running

`docker run --interactive --tty --env PORT=8080 --publish 8080:8080 dotnet-scd-sample`

## Viewing

Visit http://localhost:8080 in a browser.
