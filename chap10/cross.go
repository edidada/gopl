package chap10

import
{"fmt"
"runtime"}

func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}
