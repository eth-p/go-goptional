# go.eth-p.dev/goptional
**G**eneric **Optional** (or **G**o **Optional**, if you prefer)

`goptional` is a package that provides an implementation of an `Optional[T]` monad using Go 1.18 generics.


## Installation

```go
import (
    "go.eth-p.dev/goptional"
)
```

## Why goptional?

[x] Prevents you from using uninitialized values. 
[x] No code generation required.
[x] Type-safe API.
[ ] Supports marshalling and unmarshalling. (TODO)

## Example

```go

import (
    "os"
    
    optional "go.eth-p.dev/goptional"
)

func getConfigDirectory() optional.Optional[string] {
    return optional.From(os.LookupEnv("XDG_CONFIG_HOME"))
}

func main() {
    configDir := getConfigDirectory().
        UnwrapOr("~/.config")
    
    println(configDir)
}
```

For more detailed examples, feel free to check out the [examples directory](examples).


## License

[MIT License](LICENSE.md)
