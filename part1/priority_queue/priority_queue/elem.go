package priority_queue

import (
	"fmt"
)

type Elem struct {
	Score int
	Data  interface{}
}

func (e *Elem) String() string {
	return fmt.Sprintf("Score: %d, Data: %v", e.Score, e.Data)
}
