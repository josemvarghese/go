package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://buybackbazaar.com/")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	// fmt.Println(resp)
	// bs := make([]byte, 9999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	io.Copy(os.Stdout, resp.Body)
}
