## Maybe

[![Go Report Card](https://goreportcard.com/badge/github.com/pmorelli92/maybe)](https://goreportcard.com/report/github.com/pmorelli92/maybe)
[![CI](https://github.com/pmorelli92/maybe/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/pmorelli92/maybe/actions/workflows/ci.yml)
[![Coverage Status](https://coveralls.io/repos/github/pmorelli92/maybe/badge.svg?branch=main)](https://coveralls.io/github/pmorelli92/maybe?branch=main)

Maybe is a library that adds an [Option data type](https://en.wikipedia.org/wiki/Option_type) for Golang. [Related blogpost here.](https://devandchill.com/posts/2021/04/introducing-maybe-package-bring-functional-to-go/)

### What does it offer:

The `Maybe[any]` type exported by this library is immutable and thread safe. The json serialization and de-serialization works in the same way as with the underlying (any) type. Using this library will free you up from using pointers and possible panics.

It also gets rid of the situations where an absence of value means something different from a default (zero) value. For example: a person with salary 100 means he/she has a paid job, 0 means an unpaid internship and null means unemployed. Supporting yourself with `Maybe[int]` eliminates the usage of null replacing it with `HasValue`:

- `salary.Value != 0` has a paid job.
- `salary.Value == 0 && salary.HasValue` has an unpaid internship.
- `salary.HasValue` does not have a job, this is serialized as `null` but you don't have to care about pointers.

### When should I use it:

It can be used for transport layer (as it has json capabilities) but it could also be used on the domain layer.

### Examples:

**Marshal of Maybe[string] without value**

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/pmorelli92/maybe"
)

type Person struct {
	Name maybe.Maybe[string] `json:"name"`
	Age  int                 `json:"age"`
}

func main() {
	p := Person{Age: 28}
	bytes, _ := json.Marshal(p)
	fmt.Println(string(bytes)) // {"name":null,"age":28}
}
```

**Marshal of Maybe[string] with value**

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/pmorelli92/maybe"
)

type Person struct {
	Name maybe.Maybe[string] `json:"name"`
	Age  int                 `json:"age"`
}

func main() {
	p := Person{Age: 28, Name: maybe.Set("Pablo")}
	bytes, _ := json.Marshal(p)
	fmt.Println(string(bytes)) // {"name":"Pablo","age":28}
}
```

**Unmarshal of Maybe[string] without value**

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/pmorelli92/maybe"
)

type Person struct {
	Name maybe.Maybe[string] `json:"name"`
	Age  int                 `json:"age"`
}

func main() {
	var p Person
	_ = json.Unmarshal([]byte(`{"age":28}`), &p)
	fmt.Println(p.Name.HasValue()) // false
}
```


**Unmarshal of Maybe[string] with value**

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/pmorelli92/maybe"
)

type Person struct {
	Name maybe.Maybe[string] `json:"name"`
	Age  int                 `json:"age"`
}

func main() {
	var p Person
	_ = json.Unmarshal([]byte(`{"age":28, "name": "Pablo"}`), &p)
	fmt.Println(p.Name.HasValue()) // true
	fmt.Println(p.Name.Value())    // Pablo
}
```

### Types supported:

`Maybe` is defined to support `[T any]` so it can support all underlying types. Personally I would not suggest using pointers as the underlying type as it will defeat the whole purpose.
