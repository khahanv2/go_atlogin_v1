# Tài liệu API của ứng dụng

Đây là hướng dẫn tóm tắt làm việc với ứng dụng qua JavaScript trong trình duyệt.

## Lấy cookie và RequestVerificationToken từ trang web

Các cookies và RequestVerificationToken được sử dụng trong mọi request quan trọng.

```javascript
// Lấy RequestVerificationToken
let token = document.querySelector('input[name="__RequestVerificationToken"]').value;

// Lấy cookie
let cookie = document.cookie;
```

## Yêu cầu đăng nhập (Login Request)

Yêu cầu đăng nhập được gửi qua phương thức POST với form data.

```javascript
let formData = new FormData();
formData.append('__RequestVerificationToken', token); // Token từ trang
formData.append('Username', 'your_username');
formData.append('Password', 'your_password');

let response = await fetch('/Account/Login', {
    method: 'POST',
    body: formData,
    headers: {
        'RequestVerificationToken': token
    }
});
```

## Xử lý Captcha

Trong quá trình đăng nhập, hệ thống có thể yêu cầu giải captcha. Thông tin captcha được trả về dưới dạng JSON.

```javascript
let captchaResponse = await fetch('/Captcha/GetCaptcha', {
    method: 'GET',
    headers: {
        'RequestVerificationToken': token
    }
});

let captchaData = await captchaResponse.json();
// captchaData sẽ chứa thông tin về captcha cần giải
```

## Gửi xác thực Captcha

Sau khi giải captcha, bạn cần gửi kết quả để xác thực.

```javascript
let verifyCaptchaData = new FormData();
verifyCaptchaData.append('__RequestVerificationToken', token);
verifyCaptchaData.append('CaptchaId', captchaData.captchaId);
verifyCaptchaData.append('Solution', 'captcha_solution'); // Giải pháp cho captcha

let verifyCaptchaResponse = await fetch('/Captcha/VerifyCaptcha', {
    method: 'POST',
    body: verifyCaptchaData,
    headers: {
        'RequestVerificationToken': token
    }
});
```

## Lấy thông tin người dùng

Sau khi đăng nhập thành công, bạn có thể lấy thông tin người dùng.

```javascript
let userInfoResponse = await fetch('/User/GetInfo', {
    method: 'GET',
    headers: {
        'RequestVerificationToken': token
    }
});

let userInfo = await userInfoResponse.json();
// userInfo sẽ chứa thông tin người dùng
```

## Đăng xuất

Để đăng xuất, bạn cần gửi một yêu cầu POST.

```javascript
let logoutResponse = await fetch('/Account/Logout', {
    method: 'POST',
    headers: {
        'RequestVerificationToken': token
    }
});
```