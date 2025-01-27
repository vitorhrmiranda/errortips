# TIPs for issuing errors in GO

## 1. Make Exported Errors Public:
Exported errors allow users of your package to check for specific error types using errors.Is or errors.As.
```go
var (
  ErrNotFoundFile = errors.New("file not found")
  ErrPermission   = errors.New("permission denied")
)
```

## 2. Add Context to Errors
Context helps users debug and understand the circumstances under which the error occurred.
Wrapping errors allows users to unwrap and inspect the root cause.
```go
if err != nil {
  return fmt.Errorf("read file in main: %w", err)
}
// or
if err != nil {
  return errors.Join(ErrNotFoundFile, err)
}
```

## 3. Leverage Custom Error Types
Custom error types provide richer error information and are helpful for advanced error handling.
```go
type ValidationError struct {
  Field string
  Msg   string
}

func (e ValidationError) Error() string {
  return fmt.Sprintf("validation failed for field '%s': %s", e.Field, e.Msg)
}
```

## 4. Use Sentinel Errors for Comparisons
Sentinel errors are useful for constant error comparisons.
```go
if errors.Is(err, mypackage.ErrNotFound) {
  // handle not found case
}
// or
var perr *fs.PathError
if errors.As(err, &perr) {
	// handle validation error
}
```
## 5. Avoid Overloading Error Details
Errors should convey essential information without becoming verbose or confusing. Include the most relevant details, and if the context is complex, log additional details elsewhere.

## 6. Be Consistent Across the Package
Consistency makes your package easier to use and maintain.
Define a clear strategy for error creation and propagation, and stick to it throughout the package.

## 7. Document the Errors, Avoid Using Panics and Provide Helper Functions for Error Checking.
Panics are harder to recover from and generally should be reserved for truly exceptional situations. Return errors instead of panicking, and let the caller decide how to handle them.

Clear documentation helps package users understand the errors they may encounter and how to handle them. Use GoDoc comments to describe exported errors and their usage.

Helper functions improve the ergonomics of error handling. Expose functions to check for specific errors or enrich error contexts.
