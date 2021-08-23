package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"strings"
)

var proxies []string

func Import() (int, error) {
	proxyFile, err := os.OpenFile("proxies.txt", os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, fmt.Errorf("proxy.Import: %s \n", err)
	}

	defer proxyFile.Close()
	scanner := bufio.NewScanner(proxyFile)
	for scanner.Scan() {
		proxies = append(proxies, scanner.Text())
	}
	if len(proxies) < 1 {
		return 0, fmt.Errorf("proxy.Import: No proxies loaded\n")
	}
	return len(proxies), nil
}

func Get() (*url.URL, error) {
	if len(proxies) < 1 {
		return nil, errors.New("Pool is empty")
	}
	proxySplit := strings.Split(proxies[rand.Intn(len(proxies))], ":")
	switch len(proxySplit) {
	case 2:
		proxyUrl, _ := url.Parse(fmt.Sprintf("http://%s:%s", proxySplit[0], proxySplit[1]))
		return proxyUrl, nil
	case 4:
		proxyUrl, _ := url.Parse(fmt.Sprintf("http://%s:%s@%s:%s", proxySplit[2], proxySplit[3], proxySplit[0], proxySplit[1]))
		return proxyUrl, nil
	}
	return nil, nil
}
