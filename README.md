# pub-sub

This is a library that implements the [Publisher-Subscriber Pattern](https://learn.microsoft.com/en-us/azure/architecture/patterns/publisher-subscriber).

## Library code coverage
You can check the code coverage for the library by navigating to the `pubsub` module folder and running the following commands:

```shell
go test -coverprofile=coverage.out ./...
go tool cover -func coverage.out
```