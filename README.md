# Ignoring Errors in `http.Request.Body.Close()` in Go

This repository demonstrates a common error in Go when handling HTTP requests: ignoring the potential error returned by `http.Request.Body.Close()`.  Failing to handle this error can lead to resource leaks and unexpected behavior.

The `bug.go` file shows the problematic code, while `bugSolution.go` provides a corrected version.

## Problem

The `http.Request.Body.Close()` function can return an error indicating a failure to close the request body.  Ignoring this error can lead to: 

* **Resource leaks:** The underlying resources associated with the request body may not be released, leading to performance issues or application crashes.
* **Hidden bugs:** The error might indicate a problem with the request itself, which could remain undetected.

## Solution

Always check for errors returned by `http.Request.Body.Close()` and handle them appropriately. This typically involves logging the error or taking other corrective action based on the error type.