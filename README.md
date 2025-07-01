# Reverse Proxy in Golang
## Overview

This reverse proxy project is designed to efficiently distribute client requests across multiple backend servers. It uses a flexible load balancing strategy to enhance the performance and reliability of backend services.

## Features

- **Load Balancing:** Supports both **Round Robin** and **Least Connections** strategies to balance traffic.
- **Modular Architecture:** Easily extend or configure components such as server management and load balancing.
- **Fault Tolerance:** Automatically retries and handles failures in routing requests.

## Folder Structure

- **backend1** and **backend2** folders correspond to the locally setup servers for testing purposes.
- **reverse_proxy** folder contains the code of the reverse proxy.
