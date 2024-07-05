# AWS MemoryDB for Redis Go Application

This project demonstrates how to interact with AWS MemoryDB for Redis using Go.

## Prerequisites

Before running this application, make sure you have the following:

1. Go installed on your machine. You can download it from [here](https://golang.org/dl/).
2. AWS MemoryDB for Redis instance set up and running. Make sure to note down the endpoint.

## Environment Variables

Set the following environment variable to configure your Redis connection:

```bash
export REDIS_ADDR=your-memorydb-endpoint:6379
