<div align="center">
<h1>In memory cache</h1>

<p>
My first homework in Go Ninja course
</p>

<p>
• <a href="#about">About</a> •
<a href="#installation">Installation</a>
•<a href="#examples">Examples</a> •
</p>

</div>

## About
Package "In memory cache" that can create cache, set and delete from this cache.

## Installation
To install package enter command: <br>
- go get github.com/xopxe23/in_memory_cache

## Examples
<h4>Code</h4>

```go
package in_memory_cache

type Cache map[string]interface{}

func New() Cache {
	newCache := make(Cache)
	return newCache
}

func (c *Cache) Set(key string, value interface{}) {
	cache := *c
	cache[key] = value
}

func (c *Cache) Delete(key string) {
	cache := *c
	delete(cache, key)
}

func (c *Cache) Get(key string) interface{} {
	cache := *c
	return cache[key]
}
```
<h4>Example</h4>

```go
package main

import (
	"fmt"
	cache "github.com/Xopxe23/in_memory_cache"
)

func main() {
	cache := cache.New() 

	cache.Set("userId", 42)
	userId := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId = cache.Get("userId")
	fmt.Println(userId)
}

```