package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWriter struct {
}

func main() {
	resp, _ := http.Get("http://google.com/")

	fmt.Println(resp.Status)

	// bs := make([]byte, 99999)
	// _, err := resp.Body.Read(bs)
	// if err != nil {
	// 	fmt.Println("Somthing happened", err)
	// }

	//io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (lw logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Wrote this many bytes to terminal", len(bs))
	return len(bs), nil
}
