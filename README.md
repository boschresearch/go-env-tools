# Common env extraction tools

## Installing
```bash
go env -w GOPRIVATE=github.boschdevcloud.com/bcai-internal/
go get github.boschdevcloud.com/bcai-internal/go-env-tools
```

## Usage
```golang
import (
   "github.boschdevcloud.com/bcai-internal/go-env-tools"
)
func main() {
   envtools.GetEnvOrPanic("MY_ENV")
}
```
