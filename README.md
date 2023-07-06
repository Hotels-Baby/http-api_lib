## A GO library to abstract out common API functions into a single library

### methodlib

```go

package main

import (
	"net/http"
	"yourpath/healthCheck" // replace with the correct import path
	"github.com/Hotels-Baby/http-api_lib/methodlib" 
)

func main() {
	hc := &healthCheck.HealthCheck{
		MethodLib: methodlib.MethodLib{},
	}

	hc.MethodLib.AddMethod("GET")
	hc.MethodLib.AddMethod("POST")

	http.HandleFunc("/", hc.Check)

	http.ListenAndServe(":8080", nil)
}

```

### Example handler implementation (healthCheck)

```go
package healthCheck

import (
	"net/http"
	"github.com/Hotels-Baby/http-api_lib/methodlib" // replace with the correct import path
)

type HealthCheck struct {
	MethodLib methodlib.MethodLib
}

func (hc *HealthCheck) Check(w http.ResponseWriter, r *http.Request) {
	if !hc.MethodLib.IsMethodAllowed(r.Method) {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Handle the request...
}

````
