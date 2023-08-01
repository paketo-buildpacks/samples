# PHP Sample App With Extensions Loaded Via Custom `.ini` Snippet

This app loads PHP extensions via a custom PHP `.ini` snippet located inside
the `.php.ini.d` directory. It runs with the built-in PHP server.

## Building

`pack build php-extension-sample --buildpack paketo-buildpacks/php --builder paketobuildpacks/builder-jammy-full`

## Running

`docker run --interactive --tty --env PORT=8080 --publish 8080:8080 php-extension-sample`

## Viewing

`curl http://localhost:8080`

## Stack Support

The Paketo PHP buildpack requires the Full Jammy Stack. See [stack docs](https://paketo.io/docs/concepts/stacks) for more details
