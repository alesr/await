# Await

Await is a command-line tool written in Go that tests TCP connections by continuously attempting to establish a connection with a specified host and port. It is designed to check if a TCP service is available and reachable.

## Installation

To use Await, you need to have Go installed. Then, you can install it by running the following command:

```shell
go install github.com/alesr/await@latest
```

## Usage

You can run Await by providing the host and port as command-line arguments, along with the duration (in seconds) for how long to keep trying. Here's an example usage:

```shell
await localhost 5432 30
```

This will try to connect to `localhost:5432` for a duration of 30 seconds. If the connection is established, Await will exit with code 0. Otherwise, it will exit with code 1.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
