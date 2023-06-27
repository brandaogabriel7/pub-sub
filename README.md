# pub-sub

This is a library that implements the [Publisher-Subscriber Pattern](https://learn.microsoft.com/en-us/azure/architecture/patterns/publisher-subscriber). It allows you to publish messages to queues and subscribe to receive messages from those queues. It also allows you to pass a storage service to store the messages for each queue. It already comes with an implementation of storage service that uses the OS file system.

The `pubsub` folder contains the library itself and the `pubsub-sample` folder has a simple console application to showcase the use of the library.

## Library code coverage
You can check the code coverage for the library by navigating to the `pubsub` module folder and running the following commands:

```shell
go test -coverprofile=coverage.out ./...
go tool cover -func coverage.out

# or, if you want to see more details of which lines are not covered, you can do the following and generate an html file
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```