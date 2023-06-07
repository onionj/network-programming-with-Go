package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "http://host:port/page")
		os.Exit(1)
	}

	url, err := url.Parse(os.Args[1])
	checkError(err)

	client := &http.Client{Timeout: time.Minute}

	request, err := http.NewRequest("GET", url.String(), nil)
	checkError(err)

	request.Header.Add("Accept-Charset", "UTF-8;q=1, ISO-8859-1;q-0")

	response, err := client.Do(request)
	checkError(err)

	if response.StatusCode != 200 {
		fmt.Println(response.Status)
		os.Exit(1)
	}

	chSet := getCharset(response)
	fmt.Printf("got charset %s\n", chSet)

	if chSet != "UTF-8" && chSet != "ISO-8859-1" {
		fmt.Println("Cannot handle", chSet)
		os.Exit(4)
	}

	var buf [512]byte
	reader := response.Body
	fmt.Println("got body")

	file, err := os.Create(url.Host + ".html")
	checkError(err)
	defer file.Close()

	fmt.Println("open file:", file.Name())

	all_bytes := 0

	for {
		n, err := reader.Read(buf[0:])
		if err != nil {
			break
		}

		file.Write(buf[:n])
		all_bytes += n
	}

	fmt.Println(all_bytes, "bytes were written in the file")

	os.Exit(0)
}

func getCharset(response *http.Response) string {
	contentType := response.Header.Get("Content-Type")

	if contentType == "" {
		// guess
		return "UTF-8"
	}
	idx := strings.Index(contentType, "charset=")
	if idx == -1 {
		// guess
		return "UTF-8"
	}
	return strings.Trim(contentType[idx+len("charset="):], " ")
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
