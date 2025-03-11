package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, _ := http.NewRequest("GET", "https://google.com", nil)
	client := http.Client{}
	res, _ := client.Do(req)
	body, _ := io.ReadAll(res.Body)
	fmt.Print(string(body))

}
