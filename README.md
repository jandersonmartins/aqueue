# Aqueue

Simple asynchronous queue task.

## Install

```sh
$ go get github.com/jandersonmartins/aqueue
```

## Usage

```go
import (
	"fmt"

	"github.com/jandersonmartins/aqueue"
)

func main() {
	q := aqueue.NewAqueue(1)

	q.Add(func() {
		fmt.Println("hello")
	})

	q.Wait()
}
```
