# PHP Sample App With Extensions Loaded Via Custom `.ini` Snippet

This app loads PHP extensions via a custom PHP `.ini` snippet located inside
the `.php.ini.d` directory. It runs with the built-in PHP server.

## Building

`pack build php-extension-sample --buildpack paketo-buildpacks/php`

## Running

`docker run --interactive --tty --env PORT=8080 --publish 8080:8080 php-extension-sample`

## Viewing

`curl http://localhost:8080`
