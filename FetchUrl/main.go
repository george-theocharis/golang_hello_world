package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	var buffer bytes.Buffer
	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, "http://") {
			buffer.WriteString("http://")
		}

		buffer.WriteString(url)
		fmt.Printf("%s", buffer.String())
		resp, err := http.Get(buffer.String())
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s", resp.Status)
		b, err := io.Copy(os.Stdout, resp.Body) //ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%d", b)
	}
}
