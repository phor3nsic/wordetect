package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var wg sync.WaitGroup

	for scanner.Scan() {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			if isWordPress(url) {
				fmt.Println(url)
			}
		}(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Erro ao ler a entrada:", err)
	}

	wg.Wait()
}

func isWordPress(url string) bool {
	endpoints := []string{"/wp-content", "/blog/wp-content"}
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	for _, endpoint := range endpoints {
		finalURL, err := checkRedirect(url+endpoint, client)
		if err == nil && strings.Contains(finalURL, endpoint) {
			return true
		}
	}
	return false
}

func checkRedirect(url string, client *http.Client) (string, error) {
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Verifica se houve redirecionamento
	if resp.StatusCode == http.StatusMovedPermanently || resp.StatusCode == http.StatusFound {
		return resp.Header.Get("Location"), nil
	}

	return "", fmt.Errorf("nenhum redirecionamento encontrado")
}
