# pub-sub

This is a library that implements the [Publisher-Subscriber Pattern](https://learn.microsoft.com/en-us/azure/architecture/patterns/publisher-subscriber).

## Library code coverage
You can check the code coverage for the library by navigating to the `pubsub` module folder and running `go test -cover`.

### Generate coverage report HTML
If you want details about the code coverage, you can run these commands in your terminal to generate a `coverage.html` file. You can open this file to see the coverage report in your browser.

```shell
go test -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```