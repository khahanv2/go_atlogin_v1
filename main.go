package main

import (
	"fmt"
	"os"

	"github.com/bongg/autologin/client"
	"github.com/bongg/autologin/config"
	"github.com/bongg/autologin/logger"
	"github.com/bongg/autologin/utils"
)

func main() {
	// Khởi tạo logger - true để bật debug mode
	logger.Init(true)
	
	// Log thông báo khởi động
	logger.Info("Ứng dụng auto login bắt đầu khởi động")
	
	// Tạo cấu hình với giá trị mặc định
	cfg := config.NewConfig("", "")
	logger.Debug("Đã tạo cấu hình mặc định", "baseURL", cfg.BaseURL)
	
	// Tạo client
	cli := client.NewClient(cfg)
	logger.Info("Đã khởi tạo HTTP client")
	
	// Lấy dữ liệu ban đầu (token, cookies)
	logger.Info("Đang lấy thông tin từ trang chủ...")
	err := cli.FetchInitialData()
	if err != nil {
		logger.Error("Lỗi khi lấy dữ liệu ban đầu", err)
		os.Exit(1)
	}
	
	logger.Info("Đã lấy thông tin trang chủ thành công")
	
	// Thông tin chi tiết ở cấp độ debug
	logger.Debug("User-Agent", "value", cli.GetUserAgent())
	logger.Debug("Token", "value", cli.GetToken())
	
	cookieValue := cli.GetCookie()
	cookieType := "BBOSID"
	if utils.ExtractCookie(fmt.Sprintf("IT=%s", cookieValue)) != "" {
		cookieType = "IT"
	}
	
	logger.Debug("Cookie", "type", cookieType, "value", cookieValue)
	logger.Debug("FingerIDX", "value", cli.GetFingerIDX())
	logger.Debug("Cookies", "all", cli.GetAllCookies())
	
	if idyKey := cli.GetIdyKey(); idyKey != "" {
		logger.Debug("IdyKey", "value", idyKey)
	}
	
	// Lấy thông tin Slider Captcha
	logger.Info("Đang lấy slider captcha...")
	captchaData, err := cli.GetSliderCaptcha()
	if err != nil {
		logger.Error("Lỗi khi lấy captcha", err)
	} else {
		logger.Debug("Dữ liệu Captcha", "json", captchaData)
	}
	
	// Hiển thị thông tin CURL nếu cần debug
	logger.Debug("Thông tin cho CURL", 
		"user-agent", cli.GetUserAgent(),
		"token", cli.GetToken(),
		"cookie", fmt.Sprintf("%s=%s", cookieType, cookieValue))
	
	logger.Info("Ứng dụng hoàn thành xử lý")
}