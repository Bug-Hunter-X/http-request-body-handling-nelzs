func handleRequest(w http.ResponseWriter, r *http.Request) {
    // ... other code ...
    defer func() {
        if err := r.Body.Close(); err != nil {
            log.Printf("Error closing request body: %v", err)
            // Consider more sophisticated error handling here, like returning an error to the caller.
        }
    }()
    // ... other code ...
}