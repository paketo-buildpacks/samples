# Node.js Sample App using Yarn

## Building

`pack build yarn-sample --buildpack paketo-buildpacks/nodejs`

## Running

`docker run --interactive --tty --publish 8080:8080 yarn-sample`

## Viewing

`curl http://localhost:8080`
