# IOT Smart Agriculture Backend

Hệ thống Backend cho dự án Nông nghiệp Thông minh (Smart Agriculture), hỗ trợ quản lý thiết bị IoT, thu thập dữ liệu cảm biến và cung cấp API cho Web.

## 🚀 Giới thiệu
Dự án được xây dựng dựa trên ngôn ngữ **Go** với kiến trúc **Clean Architecture**, tập trung vào hiệu suất, khả năng mở rộng và tính bảo mật. Hệ thống hỗ trợ hai luồng truy cập chính:
- **Web API**: Dành cho người dùng quản lý, giám sát qua JWT Authentication.
- **Device API**: Dành cho các thiết bị IoT gửi/nhận dữ liệu qua API Key.

## 📋 Yêu cầu môi trường
- **Go**: v1.25 trở lên.
- **PostgreSQL**: v14 trở lên.
- **Công cụ bổ sung**: `swag` (để cập nhật tài liệu API).

## 🛠 Cài đặt & Chạy ứng dụng

### 1. Clone repository
```bash
git clone <repository_url>
cd IOT-Smart-Agriculture/backend
```

### 2. Cài đặt các package phụ thuộc
```bash
go mod tidy
```

### 3. Cấu hình biến môi trường
Tạo file `.env` từ file mẫu:
```bash
cp .env.example .env
```
Sau đó cập nhật thông tin kết nối Database và JWT Secret trong file `.env`.

### 4. Chạy ứng dụng
Hệ thống sẽ tự động chạy migration để khởi tạo các bảng cần thiết khi khởi động lần đầu.
```bash
go run cmd/server/main.go
```
Server sẽ mặc định chạy tại: `http://localhost:8080`

## 📂 Cấu trúc dự án
Dự án tuân thủ cấu trúc Clean Architecture:

```text
├── cmd/server/        # Điểm khởi đầu của ứng dụng (main.go)
├── docs/              # Swagger documentation (tự động tạo bởi swag)
├── internal/          # Code nghiệp vụ chính
│   ├── config/        # Quản lý cấu hình (biến môi trường)
│   ├── database/      # Kết nối cơ sở dữ liệu (Postgres)
│   ├── dto/           # Data Transfer Objects (Request/Response)
│   ├── handlers/      # Tầng giao diện API (nhận request, trả response)
│   ├── middlewares/   # Xử lý trung gian (Auth, Log, v.v.)
│   ├── models/        # Định nghĩa các thực thể (GORM/SQL models)
│   ├── repositories/  # Tầng tương tác trực tiếp với Database
│   ├── router/        # Định nghĩa các route API
│   └── services/      # Tầng xử lý logic nghiệp vụ chính
├── migration/         # Các file SQL migration
├── utils/             # Các hàm tiện ích (JWT, Response wrapper, v.v.)
└── .env               # File cấu hình môi trường (không commit)
```

## 📝 Thử nghiệm API

### 1. Swagger UI
Tài liệu API được tích hợp sẵn qua Swagger. Sau khi chạy server, bạn có thể truy cập tại:
`http://localhost:8080/swagger/index.html`

### 2. Xác thực (Authentication)
Dự án hỗ trợ 2 phương thức xác thực:
- **JWT (Bearer Auth)**:
  - Sử dụng cho các đầu API `/web/...`.
  - Header: `Authorization: Bearer <your_jwt_token>`
  - Đăng nhập tại `/web/auth/login` để lấy token.
- **API Key**:
  - Sử dụng cho các đầu API `/device/...`.
  - Header: `IOT-API-Key: <your_device_api_key>`

### 3. Cập nhật tài liệu Swagger
Nếu bạn thay đổi code hoặc chú thích (comments) API, hãy chạy lệnh sau để cập nhật:
```bash
swag init -g cmd/server/main.go
```

## 🛡 Tính năng nổi bật
- **Graceful Shutdown**: Đảm bảo đóng kết nối an toàn khi server dừng.
- **Standardized Response**: Phản hồi API nhất quán theo định dạng chung.
- **Partial Update**: Hỗ trợ cập nhật từng trường dữ liệu linh hoạt (PATCH).
- **Automated Migration**: Tự động quản lý schema database.
