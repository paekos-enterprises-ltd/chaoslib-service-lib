# Chaoslib Service API Client

Это клиентская библиотека на Go для взаимодействия с Chaoslib Service API. Она предоставляет удобный доступ к различным эндпоинтам сервиса, включая генерацию случайных чисел и другие игровые механики.

## Установка

```bash
go get github.com/user/chaoslib-service-lib
```

## Использование

Сначала необходимо импортировать библиотеку:

```go
import "github.com/user/chaoslib-service-lib"
```

Затем можно создавать клиенты для различных эндпоинтов API.

### Генерация случайных чисел

```go
package main

import (
	"fmt"
	"github.com/user/chaoslib-service-lib"
)

func main() {
	client := chaoslib_service_api.NewChaosLibRand("http://localhost:8080")
	
	// Получить случайное число int64
	num, err := client.Int63()
	if err != nil {
		panic(err)
	}
	fmt.Println("Случайное число:", num)

	// Получить случайное число в диапазоне [0, n)
	numN, err := client.Int63n(100)
	if err != nil {
		panic(err)
	}
	fmt.Println("Случайное число до 100:", numN)
}
```

### Crash

```go
package main

import (
	"fmt"
	"github.com/user/chaoslib-service-lib"
)

func main() {
	client := chaoslib_service_api.NewCrashProvider("http://localhost:8080")
	
	crashPoint, err := client.Crash()
	if err != nil {
		panic(err)
	}
	fmt.Println("Точка краша:", crashPoint)
}
```

### Weight

```go
package main

import (
	"fmt"
	"github.com/user/chaoslib-service-lib"
)

func main() {
	client := chaoslib_service_api.NewWeightProvider("http://localhost:8080")
	
	// Получить одно значение
	weight, err := client.Weight(0)
	if err != nil {
		panic(err)
	}
	fmt.Println("Значение веса:", weight)

	// Получить несколько значений
	weights, err := client.Weight(5)
	if err != nil {
		panic(err)
	}
	fmt.Println("Значения весов:", weights)
}
```

### Slots

```go
package main

import (
	"fmt"
	"github.com/user/chaoslib-service-lib"
)

func main() {
	client := chaoslib_service_api.NewSlotsProvider("http://localhost:8080")
	
	reels := []int{10, 20, 30}

	// Один спин
	spin, err := client.Slots(reels, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println("Результат спина:", spin)

	// Несколько спинов
	spins, err := client.Slots(reels, 3)
	if err != nil {
		panic(err)
	}
	fmt.Println("Результаты спинов:", spins)
}
```

### Penalty

```go
package main

import (
	"fmt"
	"github.com/user/chaoslib-service-lib"
)

func main() {
	client := chaoslib_service_api.NewPenaltyProvider("http://localhost:8080")
	
	penalty, err := client.Penalty()
	if err != nil {
		panic(err)
	}
	fmt.Println("Штраф:", penalty)
}
```
