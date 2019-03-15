package lib

import (
	"os"
	in "GoLecture/puzzlers/article3/q4/lib/internal"
)

func Hello(name string) {
	in.Hello(os.Stdout, name)
}
