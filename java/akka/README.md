# Scala Akka Sample Application

## Building

```bash
pack build applications/akka
```

## Running

```bash
docker run --rm --tty --publish 8080:8080 applications/akka
```

## Viewing

The log output shows the interation of the actors:

```plain
[2020-06-19 14:48:16,054] [INFO] [akka.event.slf4j.Slf4jLogger] [AkkaQuickStart-akka.actor.default-dispatcher-3] [] - Slf4jLogger started
[2020-06-19 14:48:16,123] [INFO] [com.example.Greeter$] [AkkaQuickStart-akka.actor.default-dispatcher-6] [akka://AkkaQuickStart/user/greeter] - Hello Charles!
[2020-06-19 14:48:16,124] [INFO] [com.example.GreeterBot$] [AkkaQuickStart-akka.actor.default-dispatcher-6] [akka://AkkaQuickStart/user/Charles] - Greeting 1 for Charles
[2020-06-19 14:48:16,126] [INFO] [com.example.Greeter$] [AkkaQuickStart-akka.actor.default-dispatcher-6] [akka://AkkaQuickStart/user/greeter] - Hello Charles!
[2020-06-19 14:48:16,126] [INFO] [com.example.GreeterBot$] [AkkaQuickStart-akka.actor.default-dispatcher-6] [akka://AkkaQuickStart/user/Charles] - Greeting 2 for Charles
[2020-06-19 14:48:16,127] [INFO] [com.example.Greeter$] [AkkaQuickStart-akka.actor.default-dispatcher-3] [akka://AkkaQuickStart/user/greeter] - Hello Charles!
[2020-06-19 14:48:16,127] [INFO] [com.example.GreeterBot$] [AkkaQuickStart-akka.actor.default-dispatcher-6] [akka://AkkaQuickStart/user/Charles] - Greeting 3 for Charles
```
