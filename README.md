
# Cesar

Cesar is a simple REST API server that converts a range of numbers to Roman numerals.

## Features

- Converts a range of numbers to Roman numerals.
- Supports both integer and Roman numeral representations.

## Installation

To install Cesar, you need to have Go installed on your system. Then, you can clone the repository and build the binary:

```bash
git clone https://github.com/gi4nks/cesar.git
cd cesar
go build -o bin/cesar cmd/cesar/main.go
```

It is possible also to use make to build using the Makefile.

```bash
git clone https://github.com/gi4nks/cesar.git
cd cesar
make build
```

## Usage

### Running the Server

To start the Cesar server, run the following command:

```bash
./bin/cesar serve
```

By default, the server will run on port 8080.

### Making Requests

You can make requests to the server using HTTP GET requests. Here are some examples:

- Convert a range of numbers to Roman numerals:
  ```
  GET /convert/{start}/{end}
  ```
  Replace `{start}` and `{end}` with the desired range of numbers.

- Example:
  ```
  GET /convert/1/10
  ```

- Convert a single number to Roman numerals:
  ```
  GET /convert/{start}
  ```
  Replace `{start}` with the desired numbers.

- Example:
  ```
  GET /convert/1
  ```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.