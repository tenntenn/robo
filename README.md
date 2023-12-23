# robo

## How to use

```go
package main

import (
	"os"

	"github.com/tenntenn/robo"
)

func main() { os.Exit(robo.Main(run)) }

func run(r *robo.Robo) {
	r.R()
	r.R()

	r.S()
	r.S()

	r.U()
	r.U()
}
```

## Copyright

[いらすとや](https://www.irasutoya.com/) has the copyright of images under the img directory.
