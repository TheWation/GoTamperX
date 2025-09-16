package httpclient

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/TheWation/GoTamperX/internal/errors"
)

type Result struct {
	Method      string
	StatusCode  int
	ContentLen  int
	Error       error
}

type Client struct {
	timeout      time.Duration
	proxy        string
	randomAgent  bool
	httpClient   *http.Client
	userAgents   []string
	headers      map[string]string
}

func NewClient(timeout int, proxy string, randomAgent bool, headers []string) *Client {
	client := &Client{
		timeout:     time.Duration(timeout) * time.Second,
		proxy:       proxy,
		randomAgent: randomAgent,
		headers:     make(map[string]string),
		userAgents: []string{
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
			"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:89.0) Gecko/20100101 Firefox/89.0",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:89.0) Gecko/20100101 Firefox/89.0",
			"Mozilla/5.0 (X11; Linux x86_64; rv:89.0) Gecko/20100101 Firefox/89.0",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1.1 Safari/605.1.15",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Edge/91.0.864.59",
			"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36 Edg/91.0.864.59",
		},
	}

	// Parse custom headers
	for _, header := range headers {
		if header != "" {
			parts := strings.SplitN(header, ":", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				client.headers[key] = value
			}
		}
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	if proxy != "" {
		proxyURL, err := url.Parse(proxy)
		if err == nil {
			transport.Proxy = http.ProxyURL(proxyURL)
		}
	}

	client.httpClient = &http.Client{
		Timeout:   client.timeout,
		Transport: transport,
	}

	return client
}

func (c *Client) TestMethods(targetURL string, methods []string, concurrency int) []Result {
	if !c.isValidURL(targetURL) {
		return []Result{{
			Method: "VALIDATION",
			Error:  fmt.Errorf("invalid URL format"),
		}}
	}

	results := make(chan Result, len(methods))
	var wg sync.WaitGroup

	semaphore := make(chan struct{}, concurrency)

	for _, method := range methods {
		wg.Add(1)
		go func(method string) {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() { <-semaphore }()
			
			result := c.makeRequest(method, targetURL)
			results <- result
		}(method)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var allResults []Result
	for result := range results {
		allResults = append(allResults, result)
	}

	return allResults
}

func (c *Client) makeRequest(method, url string) Result {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return Result{
			Method: method,
			Error:  fmt.Errorf("failed to create request: %v", err),
		}
	}

	if c.randomAgent {
		req.Header.Set("User-Agent", c.getRandomUserAgent())
	} else {
		req.Header.Set("User-Agent", "TamperX/2.0")
	}
	req.Header.Set("Accept", "*/*")

	// Add custom headers
	for key, value := range c.headers {
		req.Header.Set(key, value)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		errorMsg := errors.GetDetailedError(err)
		return Result{
			Method: method,
			Error:  fmt.Errorf("%s", errorMsg),
		}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{
			Method:     method,
			StatusCode: resp.StatusCode,
			ContentLen: 0,
			Error:      fmt.Errorf("failed to read response: %v", err),
		}
	}

	return Result{
		Method:     method,
		StatusCode: resp.StatusCode,
		ContentLen: len(body),
		Error:      nil,
	}
}

func (c *Client) isValidURL(urlStr string) bool {
	u, err := url.Parse(urlStr)
	if err != nil {
		return false
	}
	
	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}
	
	if u.Host == "" {
		return false
	}
	
	urlPattern := regexp.MustCompile(`^https?://(?:[a-z0-9-]+\.)+[a-z]{2,}/?`)
	return urlPattern.MatchString(urlStr)
}

func (c *Client) getRandomUserAgent() string {
	if len(c.userAgents) == 0 {
		return "TamperX/2.0"
	}
	
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(c.userAgents))))
	if err != nil {
		return "TamperX/2.0"
	}
	
	return c.userAgents[n.Int64()]
}
