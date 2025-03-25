# Go Autologin Tool

Đây là công cụ hỗ trợ tự động đăng nhập và xử lý captcha, được viết bằng Go.

## Tính năng

- Tự động lấy thông tin từ trang web (token, cookies)
- Xử lý captcha tự động
- Mô phỏng hoạt động của trình duyệt web
- Hỗ trợ sử dụng proxy

## Cài đặt

```bash
# Clone repository
git clone https://github.com/yourusername/go_atlogin_v1.git
cd go_atlogin_v1

# Cài đặt dependencies
go mod download
```

## Sử dụng

```bash
# Chạy ứng dụng
go run main.go
```

## Cấu trúc dự án

- `main.go`: File chính của ứng dụng
- `client/`: Chứa mã nguồn cho HTTP client
- `config/`: Cấu hình ứng dụng
- `utils/`: Các hàm tiện ích
- `doc.md`: Tài liệu API

## Dependencies

- `github.com/go-resty/resty/v2`: HTTP client library
- `github.com/xuri/excelize/v2`: Thư viện xử lý file Excel

## Giấy phép

MIT License