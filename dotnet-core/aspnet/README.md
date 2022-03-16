# Dotnet Core Sample App using ASPNet

## Building

`pack build dotnet-aspnet-sample --buildpack paketo-buildpacks/dotnet-core`

## Running

`docker run --interactive --tty --env PORT=8080 --publish 8080:8080 dotnet-aspnet-sample`

## Viewing

`curl http://localhost:8080`
