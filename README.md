# ChaosLib Service API Client

This is a Go client library for interacting with the Chaoslib Service API. It provides convenient access to various service endpoints, including random number generation and other game mechanics.

## Installation

```bash
go get github.com/paekos-enterprises-ltd/chaoslib-service-lib
```

## Usage

First, import the library:

```go
import "github.com/paekos-enterprises-ltd/chaoslib-service-lib"
```

Then you can create clients for different API endpoints.

## Random Number Generation

```go
package main

import (
	"fmt"
	"github.com/paekos-enterprises-ltd/chaoslib-service-lib"
)

func main() {
	client := chaoslib_service_api.NewChaosLibRand("http://localhost:8033")

	// Get a random int64 number
	num, err := client.Int63()
	if err != nil {
		panic(err)
	}
	fmt.Println("Random number:", num)

	// Get a random number in range [0, n)
	numN, err := client.Int63n(100)
	if err != nil {
		panic(err)
	}
	fmt.Println("Random number up to 100:", numN)
}
```

### Crash

```go
package main

import (
	"fmt"
	"github.com/paekos-enterprises-ltd/chaoslib-service-lib"
)

func main() {
	client := chaoslib_service_api.NewCrashProvider("http://localhost:8033")

	crashPoint, err := client.Crash()
	if err != nil {
		panic(err)
	}
	fmt.Println("Crash point:", crashPoint)
}
```

### Weight

```go
package main

import (
	"fmt"
	"github.com/paekos-enterprises-ltd/chaoslib-service-lib"
)

func main() {
	client := chaoslib_service_api.NewWeightProvider("http://localhost:8033")

	// Get a single value
	weight, err := client.Weight(0)
	if err != nil {
		panic(err)
	}
	fmt.Println("Weight value:", weight)

	// Get multiple values
	weights, err := client.Weight(5)
	if err != nil {
		panic(err)
	}
	fmt.Println("Weight values:", weights)
}
```

### Slots

```go
package main

import (
	"fmt"
	"github.com/paekos-enterprises-ltd/chaoslib-service-lib"
)

func main() {
	client := chaoslib_service_api.NewSlotsProvider("http://localhost:8033")

	reels := []int{10, 20, 30}

	// Single spin
	spin, err := client.Slots(reels, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println("Spin result:", spin)

	// Multiple spins
	spins, err := client.Slots(reels, 3)
	if err != nil {
		panic(err)
	}
	fmt.Println("Spin results:", spins)
}
```

### Penalty

```go
package main

import (
	"fmt"
	"github.com/paekos-enterprises-ltd/chaoslib-service-lib"
)

func main() {
	client := chaoslib_service_api.NewPenaltyProvider("http://localhost:8033")

	penalty, err := client.Penalty()
	if err != nil {
		panic(err)
	}
	fmt.Println("Penalty:", penalty)
}
```

If you want, I can also:
•	polish it to sound more “enterprise-grade”
•	add badges / versioning / examples section
•	or convert it into a proper pkg.go.dev-style README