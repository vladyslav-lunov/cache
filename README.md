# Cache

A simple, lightweight in-memory cache library for Go.

## Installation

```bash
go get github.com/vladyslav-lunov/cache
```

## Features

- Simple key-value storage
- Generic value storage using `interface{}`
- Basic cache operations: Set, Get, Delete, Clear

## Usage

### Basic Example

```go
package main

import (
    "fmt"
    "github.com/vladyslav-lunov/cache"
)

func main() {
    // Create a new cache instance
    c := cache.New()

    // Set values
    c.Set("username", "john_doe")
    c.Set("age", 30)
    c.Set("active", true)

    // Get values
    if username, ok := c.Get("username"); ok {
        fmt.Println("Username:", username.(string))
    }

    if age, ok := c.Get("age"); ok {
        fmt.Println("Age:", age.(int))
    }

    // Delete a value
    c.Delete("age")

    // Check if value exists after deletion
    if _, ok := c.Get("age"); !ok {
        fmt.Println("Age was deleted")
    }

    // Clear all cache entries
    c.Clear()
}
```

### Storing Structs

```go
package main

import (
    "fmt"
    "github.com/vladyslav-lunov/cache"
)

type User struct {
    ID       int
    Username string
    Email    string
}

func main() {
    c := cache.New()

    // Store a struct
    user := User{
        ID:       1,
        Username: "alice",
        Email:    "alice@example.com",
    }
    c.Set("user:1", user)

    // Retrieve and type assert
    if val, ok := c.Get("user:1"); ok {
        if user, ok := val.(User); ok {
            fmt.Printf("User: %+v\n", user)
        }
    }
}
```

## API Reference

### `New() *Cache`

Creates and returns a new cache instance.

```go
c := cache.New()
```

### `Set(key string, value interface{})`

Stores a value in the cache with the specified key. If the key already exists, the value is overwritten.

```go
c.Set("key", "value")
```

### `Get(key string) (interface{}, bool)`

Retrieves a value from the cache. Returns the value and a boolean indicating whether the key was found.

```go
value, ok := c.Get("key")
if ok {
    // Use value
}
```

### `Delete(key string)`

Removes a key-value pair from the cache.

```go
c.Delete("key")
```

### `Clear()`

Removes all entries from the cache.

```go
c.Clear()
```

## License

MIT
