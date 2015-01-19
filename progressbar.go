package goutil

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

func Progressbar() {
	for i := 0; i < 50; i++ {
		time.Sleep(100 * time.Millisecond)
		h := strings.Repeat("=", i) + strings.Repeat(" ", 49-i)
		fmt.Printf("\r[%s]%.0f%%", h, float64(i)/49*100)
		os.Stdout.Sync()
	}
	fmt.Println("\nAll System Go!")
}
