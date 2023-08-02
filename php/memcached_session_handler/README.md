# PHP Sample App using Memcached Instance Session Handler

This app leverages a PHP session handler for a memcached instance by providing a
[binding](https://paketo.io/docs/howto/configuration/#bindings) of type
`php-memcached-session`. In order to use this app, you will need to set up a memcached
instance and enable basic container to container networking as shown below.

## Building

1. Run a Memcached instance, in this example we are using the default Memcached Docker image:
`docker run -it  -p 6379 memcached:latest`

2. Get the Memcached container bridge network IP address:
`docker inspect <container ID> | jq -r 'map(.NetworkSettings.Networks.bridge.IPAddress) []'`
 You should get something like `172.17.0.2` back.

3. Create a `host` file inside of the app `binding` directory with the value
   set to the IP address from step 2. Also create a `servers` file with the IP
   address in it as well.
    The binding directory should now contain:
    `type`: `php-memcached-session `
    `host`: <IP address from step 2>
    `servers`: <IP address from step 2>

4. Build the app with the service binding:
```
pack build php-memcached-handler-sample \
--env BP_PHP_WEB_DIR=htdocs \
--env SERVICE_BINDING_ROOT=/bindings \
--volume $PWD/binding:/bindings/php-memcached-session \
--buildpack paketo-buildpacks/php \
--builder paketobuildpacks/builder-jammy-full
```

## Running

`docker run --interactive --tty --env PORT=8080 --publish 8080:8080 php-memcached-handler-sample`

## Viewing

You can observe the stored state from the session handler by reaching the app
twice, using a cookie in the request.

1. `curl localhost:8080 --cookie-jar /tmp/cookie`

Observe the counter that displays the number of hits the page has had serving `1`.

2. `curl localhost:8080 --cookie /tmp/cookie`
Observe the counter that displays the number of hits the page has had serving `2`.

Alternatively, view `localhost:8080` in your browser a few times and see the
page count increment.

## Stack Support

The Paketo PHP buildpack requires the Full Jammy Stack. See [stack docs](https://paketo.io/docs/concepts/stacks) for more details
