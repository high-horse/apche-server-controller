# Apache Control Program

This Go program provides a simple command-line interface to control the Apache server on Ubuntu systems.

## Prerequisites

- Go programming language installed on your system
- Ubuntu operating system
- Apache2 installed
- Sudo privileges

## Installation

1. Clone this repository or download the `apache-control.go` file.
2. Open a terminal and navigate to the directory containing the Go file.
3. Compile the program:
   ```
   go build main.go
   ```

## Usage

Run the program with sudo and one of the following commands:

- Start Apache:
  ```
  sudo ./go-apache start
  ```

- Stop Apache:
  ```
  sudo ./go-apache stop
  ```

- Restart Apache:
  ```
  sudo ./go-apache restart
  ```

- Show help:
  ```
  ./go-apache help
  ```

## What it does

This program uses `systemctl` to control the Apache service:
- `start`: Starts the Apache server
- `stop`: Stops the Apache server
- `restart`: Restarts the Apache server
- `help`: Displays usage information

## Note

This program requires sudo privileges to control the Apache service. Make sure you have the necessary permissions before running it.

## Troubleshooting

If you encounter any errors, check that:
- Apache2 is installed on your system
- You have sudo privileges
- The Apache service name is `apache2` (this may vary on some systems)

For any other issues, please refer to the Apache documentation or seek assistance from the Ubuntu community forums.