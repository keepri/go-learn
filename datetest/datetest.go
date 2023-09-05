package datetest

import (
	"fmt"
	"time"

	tinytime "github.com/wagslane/go-tinytime"
)

func Run() {
	tt := tinytime.New(1585750374)

	tt = tt.Add(time.Hour * 48)
	fmt.Println(tt)
}
