# gRPC Calculator

This is a simple calculator application that uses gRPC (Google Remote Procedure Call) to perform basic arithmetic operations. It's designed as a learning project to help you understand the basics of gRPC.

Add some more complex operations!

## Getting Started

These instructions will guide you through the process of setting up and running the gRPC Calculator project on your local machine.

### Prerequisites

- [Go](https://golang.org/) (version 1.16 or later)
- [protoc](https://github.com/protocolbuffers/protobuf) (Protocol Buffers compiler)
- [gRPC-Go](https://github.com/grpc/grpc-go) library

### Setup

1. Clone the repository:

   ```
   git clone https://github.com/your-username/grpc-calculator.git
   ```

2. Change to the project directory:

   ```
   cd grpc-calculator
   ```

3. Generate the gRPC code:

   ```
   make generate
   ```

   This command generates the necessary Go code for the gRPC server and client.

4. Build the server:

   ```
   make build-server
   ```

5. Build the client:

   ```
   make build-client
   ```

### Running the Application

1. Start the server:

   ```
   ./calculator-server
   ```

   The server will start listening for gRPC requests on `localhost:2706`.

2. Run the client:

   ```
   ./calculator-client
   ```

   The client will connect to the server and perform some basic arithmetic operations.

## Usage

The gRPC Calculator supports the following operations:

- Addition
- Subtraction
- Multiplication
- Division
- Sum

You can interact with the calculator by running the client and providing the necessary input.
