# Concurrent JSON Sum Calculator

This project is a Go application that reads a JSON file containing pairs of integers, processes the file concurrently using goroutines, and calculates the total sum of the integers. The application supports graceful shutdown to ensure all goroutines are properly terminated.

## Features

- Concurrent processing of JSON data using producer-consumer pattern
- Graceful shutdown handling with `SIGINT` and `SIGTERM` signals
- Clean code architecture with separated concerns
