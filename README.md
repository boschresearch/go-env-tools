# Common env extraction tools

## Installing
```bash
go get github.com/boschresearch/go-env-tools
```

## Usage
```golang
import (
   "github.com/boschresearch/go-env-tools"
)
func main() {
   envtools.GetEnvOrPanic("MY_ENV")
}
```
