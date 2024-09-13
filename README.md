# Language Performance HTTP Test

This project benchmarks the performance of simple HTTP servers implemented in Go, Python (Flask), and Node.js. It uses tools like `wrk` to compare key performance metrics such as requests per second (RPS), latency, and data transfer rates.

## Table of Contents
- [Introduction](#introduction)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
  - [Go HTTP Server](#go-http-server)
  - [Python HTTP Server](#python-http-server)
  - [Node.js HTTP Server](#nodejs-http-server)
- [Running HTTP Servers](#running-http-servers)
- [Benchmarking with `wrk`](#benchmarking-with-wrk)
- [Results](#results)
- [License](#license)

## Introduction

This project compares the performance of HTTP servers in Go, Python, and Node.js under concurrent load. Each language has a minimal HTTP server implementation that responds to a simple request. The performance of each server is measured by benchmarking tools (`wrk`) to evaluate:
- **Requests per second (RPS)**
- **Latency**
- **Transfer per second**

## Prerequisites

Before you can install and run the project, make sure you have the following tools installed:

- **Go**: [Download Go](https://golang.org/doc/install)
- **Python** (version 3.6+): [Download Python](https://www.python.org/downloads/)
- **Node.js** (version 12+): [Download Node.js](https://nodejs.org/en/download/)
- **wrk** for benchmarking
  - **wrk**: `brew install wrk` (on Mac) or `sudo apt-get install wrk` (on Linux)

## Installation

### Go HTTP Server

1. Navigate to the Go server directory:
    ```bash
    cd go/server
    ```

2. Run the Go server:
    ```bash
    go run main.go
    ```

3. The Go server will be running on `http://localhost:8080`.

### Python HTTP Server

1. Navigate to the Python server directory:
    ```bash
    cd python/server
    ```

2. Install the required Python dependencies:
    ```bash
    pip install flask
    ```

3. Run the Python server:
    ```bash
    python3 server.py
    ```

4. The Python server will be running on `http://localhost:8080`.

### Node.js HTTP Server

1. Navigate to the Node.js server directory:
    ```bash
    cd nodejs/server
    ```

2. Install the required Node.js dependencies:
    ```bash
    npm install
    ```

3. Run the Node.js server:
    ```bash
    node server.js
    ```

4. The Node.js server will be running on `http://localhost:8080`.

## Running HTTP Servers

You need to run the Go, Python, and Node.js servers separately in different terminal windows or tabs. Each server will listen on port `8080`.

## Benchmarking with `wrk`

Once the servers are up and running, you can benchmark them using the `wrk` tool. For example, to benchmark with 16 threads, 100 concurrent connections, and a duration of 30 seconds:

```bash
wrk -t16 -c100 -d30s http://localhost:8080
```

### Benchmark Results

**Node.js:**

```plaintext
Running 30s test @ http://localhost:8080
  16 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.43ms    1.41ms  37.11ms   97.38%
    Req/Sec     2.58k   270.83     4.50k    93.31%
  1232506 requests in 30.04s, 224.50MB read
Requests/sec:  41031.59
Transfer/sec:      7.47MB
```

**Python:**

```plaintext
Running 30s test @ http://localhost:8080
  16 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    36.30ms    8.15ms  109.80ms   92.36%
    Req/Sec   160.27     34.29   217.00     85.50%
  16420 requests in 30.09s, 3.05MB read
  Socket errors: connect 44, read 0, write 0, timeout 0
Requests/sec:    545.72
Transfer/sec:    103.92KB
```

**Go:**

```plaintext
Running 30s test @ http://localhost:8080
  16 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.26ms    1.60ms  47.19ms   94.26%
    Req/Sec     5.73k     1.34k   24.23k    80.75%
  2737073 requests in 30.03s, 355.00MB read
Requests/sec:  91141.77
Transfer/sec:     11.82MB
```

## Results

After benchmarking, you can compare the following metrics:
- **Requests per second (RPS)**: Higher is better.
- **Latency**: Lower is better.
- **Transfer per second**: Higher is better, indicating efficient handling of data.

Based on the benchmark results:
- **Go** performed the best in terms of requests per second and transfer per second.
- **Node.js** had lower latency compared to Python but was not as fast as Go.
- **Python** had the highest latency and the lowest requests per second among the three.

## License

This project is licensed under the MIT License.
