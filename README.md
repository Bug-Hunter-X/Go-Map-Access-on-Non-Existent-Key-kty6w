# Go Map Access on Non-Existent Key

This example demonstrates a common but subtle issue in Go when accessing keys in maps that do not exist.  Attempting to access a non-existent key in a Go map will not result in an error, but instead will return the zero value for the map's value type. This can easily lead to unexpected behavior and bugs, especially when working with numerical types where a zero value might be a valid data point.  The solution presents a safer approach.

## Bug

The `bug.go` file showcases the problem.  Accessing the non-existent key "b" returns 0, which may be misinterpreted as a valid value rather than an indication that the key is missing.