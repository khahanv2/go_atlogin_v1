package utils

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
)

// GenerateRandomBytes generates random bytes of specified length
func GenerateRandomBytes(length int) ([]byte, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// GenerateRandomHex generates a random hex string of specified byte length
func GenerateRandomHex(length int) (string, error) {
	bytes, err := GenerateRandomBytes(length)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// GenerateRandomBase64 generates a random base64 string of specified byte length
func GenerateRandomBase64(length int) (string, error) {
	bytes, err := GenerateRandomBytes(length)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}

// ExtractCookie extracts cookie value from cookie string
func ExtractCookie(cookieStr string) string {
	// Extract BBOSID cookie
	re := regexp.MustCompile(`BBOSID=([^;]+)`)
	matches := re.FindStringSubmatch(cookieStr)
	if len(matches) > 1 {
		return matches[1]
	}

	// Extract IT cookie
	reIT := regexp.MustCompile(`IT=([^;]+)`)
	matchesIT := reIT.FindStringSubmatch(cookieStr)
	if len(matchesIT) > 1 {
		return matchesIT[1]
	}

	return ""
}

// FormatCookieHeader formats cookies for HTTP header
func FormatCookieHeader(cookies map[string]string) string {
	var result []string
	for name, value := range cookies {
		result = append(result, fmt.Sprintf("%s=%s", name, value))
	}
	return strings.Join(result, "; ")
}

// ParseCookieHeader parses cookie header into map
func ParseCookieHeader(cookieStr string) map[string]string {
	result := make(map[string]string)
	parts := strings.Split(cookieStr, ";")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		keyVal := strings.SplitN(part, "=", 2)
		if len(keyVal) == 2 {
			result[keyVal[0]] = keyVal[1]
		}
	}
	return result
}

// ExtractField extracts field from HTML content
func ExtractField(content, fieldName string) string {
	pattern := fmt.Sprintf(`name=[\"']%s[\"']\s+value=[\"']([^\"']+)[\"']`, regexp.QuoteMeta(fieldName))
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(content)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

// ExtractJSONValue extracts a value from JSON-like string
func ExtractJSONValue(content, key string) string {
	pattern := fmt.Sprintf(`\"%s\"\s*:\s*\"([^\"]+)\"`, regexp.QuoteMeta(key))
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(content)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}