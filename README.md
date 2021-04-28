## Maybe

[![Go Report Card](https://goreportcard.com/badge/github.com/pmorelli92/maybe)](https://goreportcard.com/report/github.com/pmorelli92/maybe)

Maybe is a library that adds an [Option data type](https://en.wikipedia.org/wiki/Option_type) for some native Go types.

### What does it offer:

The types exported by this library are immutable and thread safe. The json serialization and deserialization works in the same way as with the native types. Using this library will free you up from using pointers and possible panics.

It also gets rid of the situations where an absence of value means something different from a default (zero) value. For example: a person with salary 100 means he/she has a paid job, 0 means an unpaid internship and null means unemployed. Supporting yourself with Option eliminates the usage of null replacing it with `HasValue`.


### When should I use it:

It can be used for transport layer (as it has json capabilities) but it could also be used on the domain layer.

### Examples:

**Marshal of String Option without value**

```go
package main

import (
	"fmt"
	"encoding/json"
	"github.com/pmorelli92/maybe"
)

type Person struct {
	Name maybe.String `json:"name"`
	Age  int          `json:"age"`
}

func main() {
	p := Person{Age: 28}
	bytes, _ := json.Marshal(p)
	fmt.Println(string(bytes)) // {"name":"null","age":28}
}
```

**Marshal of String Option with value**

```go
package main

import (
    "fmt"
    "encoding/json"
    "github.com/pmorelli92/maybe"
)

type Person struct {
    Name maybe.String `json:"name"`
    Age  int          `json:"age"`
}

func main() {
    p := Person{Age: 28, Name: maybe.SetString("Pablo")}
    bytes, _ := json.Marshal(p)
    fmt.Println(string(bytes)) // {"name":"Pablo","age":28}
}
```

**Unmarshal of String Option without value**

```go
package main

import (
    "fmt"
    "encoding/json"
    "github.com/pmorelli92/maybe"
)

type Person struct {
    Name maybe.String `json:"name"`
    Age  int          `json:"age"`
}

func main() {
    var p Person
    _ = json.Unmarshal([]byte(`{"age":28}`), &p)
    fmt.Println(p.Name.HasValue()) // false
}
```


**Unmarshal of String Option with value**

```go
package main

import (
    "fmt"
    "encoding/json"
    "github.com/pmorelli92/maybe"
)

type Person struct {
    Name maybe.String `json:"name"`
    Age  int          `json:"age"`
}

func main() {
    var p Person
    _ = json.Unmarshal([]byte(`{"age":28, "name": "Pablo"}`), &p)
    fmt.Println(p.Name.HasValue()) // true
    fmt.Println(p.Name.Value()) // Pablo
}
```

### Types supported:

- bool
- string
- float
- int
- time

If this library is not supporting certain type, feel free to do a pull request or add an issue asking for it.

### Generics

Go does not support generics as of now, but the draft was recently approved. When they become available on Go 1.18 this library will be updated and only a generic struct will remain.
The library will look like this: [go2playgrounds](https://go2goplay.golang.org/p/YBqR5GX7N6m).
