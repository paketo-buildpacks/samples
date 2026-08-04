# Python sample app using uv package manager

## Building

`pack build uv-sample --env "BP_ENABLE_PACKAGE_MANAGERS=true" --buildpack paketo-buildpacks/python`

## Running

`docker run --interactive --tty --env PORT=8080 --publish 8080:8080 uv-sample`

## Viewing

`curl http://localhost:8080`
