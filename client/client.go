package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/bongg/autologin/config"
	"github.com/go-resty/resty/v2"
)

// Client represents an HTTP client for the application
type Client struct {
	cfg        *config.Config
	httpClient *resty.Client
	userAgent  string
	token      string
	cookie     string
	allCookies string
	fingerIDX  string
	idyKey     string
}

// NewClient creates a new client instance
func NewClient(cfg *config.Config) *Client {
	client := &Client{
		cfg:        cfg,
		httpClient: resty.New().SetTimeout(30 * time.Second),
		userAgent:  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36",
		fingerIDX:  "adg345asdf98723nkasdf8723", // Giả lập fingerprint
	}
	return client
}

// FetchInitialData gets initial data from the homepage
func (c *Client) FetchInitialData() error {
	// Thực hiện request đến trang chủ
	resp, err := c.httpClient.R().
		SetHeader("User-Agent", c.userAgent).
		Get(c.cfg.BaseURL)

	if err != nil {
		return fmt.Errorf("error requesting homepage: %w", err)
	}

	// Kiểm tra status code
	if resp.StatusCode() != 200 {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	// Lấy cookies
	cookies := resp.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "BBOSID" || cookie.Name == "IT" {
			c.cookie = cookie.Value
		}
		c.allCookies += fmt.Sprintf("%s=%s; ", cookie.Name, cookie.Value)
	}

	// Lấy RequestVerificationToken từ nội dung HTML
	body := string(resp.Body())
	tokenStart := strings.Index(body, "__RequestVerificationToken")
	if tokenStart != -1 {
		tokenStart = strings.Index(body[tokenStart:], "value=") + tokenStart + 7
		tokenEnd := strings.Index(body[tokenStart:], "\"") + tokenStart
		if tokenEnd > tokenStart {
			c.token = body[tokenStart:tokenEnd]
		}
	}

	return nil
}

// GetSliderCaptcha gets slider captcha information
func (c *Client) GetSliderCaptcha() (string, error) {
	// Gửi request lấy captcha
	resp, err := c.httpClient.R().
		SetHeader("User-Agent", c.userAgent).
		SetHeader("RequestVerificationToken", c.token).
		SetCookie(&http.Cookie{Name: "BBOSID", Value: c.cookie}).
		Get(fmt.Sprintf("%s/Captcha/GetCaptcha", c.cfg.BaseURL))

	if err != nil {
		return "", fmt.Errorf("error requesting captcha: %w", err)
	}

	// Kiểm tra status code
	if resp.StatusCode() != 200 {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	return string(resp.Body()), nil
}

// Các getter methods
func (c *Client) GetUserAgent() string {
	return c.userAgent
}

func (c *Client) GetToken() string {
	return c.token
}

func (c *Client) GetCookie() string {
	return c.cookie
}

func (c *Client) GetAllCookies() string {
	return c.allCookies
}

func (c *Client) GetFingerIDX() string {
	return c.fingerIDX
}

func (c *Client) GetIdyKey() string {
	return c.idyKey
}