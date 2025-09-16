package errors

import (
	"net"
	"net/url"
	"strings"
)

func GetDetailedError(err error) string {
	if netErr, ok := err.(net.Error); ok {
		if netErr.Timeout() {
			return "timeout"
		}
		if netErr.Temporary() {
			return "temporary network error"
		}
	}
	
	if dnsErr, ok := err.(*net.DNSError); ok {
		if dnsErr.IsNotFound {
			return "domain not found"
		}
		if dnsErr.IsTemporary {
			return "temporary DNS error"
		}
		return "DNS error"
	}
	
	if urlErr, ok := err.(*url.Error); ok {
		if urlErr.Timeout() {
			return "timeout"
		}
		if urlErr.Temporary() {
			return "temporary error"
		}
		if strings.Contains(urlErr.Error(), "connection refused") {
			return "connection refused"
		}
		if strings.Contains(urlErr.Error(), "no such host") {
			return "host not found"
		}
		if strings.Contains(urlErr.Error(), "network is unreachable") {
			return "network unreachable"
		}
	}
	
	errorStr := err.Error()
	if strings.Contains(errorStr, "connection reset") {
		return "connection reset"
	}
	if strings.Contains(errorStr, "connection refused") {
		return "connection refused"
	}
	if strings.Contains(errorStr, "no such host") {
		return "host not found"
	}
	if strings.Contains(errorStr, "timeout") {
		return "timeout"
	}
	
	return "network error"
}
