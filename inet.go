package autil

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var InetUseragent = "" //User-Agent string used by Inet functions

// InetGet downloads a file using the HTTP/HTTPS protocol.
// The file is created/overwritten using the given filename.
func InetGet(url, filename string) error {

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", InetUseragent)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// InetRead reads the content of a file using the HTTP/HTTPS protocol.
// The whole filesize will be read into RAM.
func InetRead(url string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", InetUseragent)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// InetReadLimited reads the content of a file using the HTTP/HTTPS protocol.
// The download stops after a given limit (in bytes).
func InetReadLimited(url string, limit int64) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", InetUseragent)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, limit))
	if err != nil {
		return "", err
	}

	return string(body), nil
}
