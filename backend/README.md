# IOT Smart Agriculture Backend

Hệ thống backend cho dự án Nông nghiệp Thông minh (Smart Agriculture), được xây dựng bằng ngôn ngữ **Go** và framework **Gin**.

## Công nghệ sử dụng

- **Ngôn ngữ:** Go (Golang) v1.25.0
- **Web Framework:** [Gin Gonic](https://github.com/gin-gonic/gin)
- **Cơ sở dữ liệu:** PostgreSQL
- **Thư viện DB:** `pgx/v5`
- **Migration:** `golang-migrate`
- **Quản lý biến môi trường:** `godotenv`

## Cấu trúc thư mục

- `cmd/server/`: Điểm khởi đầu của ứng dụng (`main.go`).
- `internal/`: Chứa mã nguồn logic nội bộ (handlers, services, repositories, models, dto, middlewares).
- `migration/`: Các file SQL migration để khởi tạo và cập nhật cấu trúc DB.
- `utils/`: Các tiện ích hỗ trợ như Dependency Injection.

## Hướng dẫn cài đặt

### 1. Yêu cầu hệ thống
- Go v1.25.0 trở lên.
- PostgreSQL.

### 2. Các bước cài đặt
1. Truy cập vào thư mục backend:
   ```bash
   cd backend
   ```
2. Cài đặt các dependencies:
   ```bash
   go mod download
   ```
3. Cấu hình biến môi trường:
   - Sao chép file `.env.example` thành `.env`:
     ```bash
     cp .env.example .env
     ```
   - Cập nhật các giá trị trong `.env`:
     - `PORT`: Cổng chạy server (VD: 8080).
     - `DATABASE_URL`: Đường dẫn kết nối Postgres (VD: `postgres://user:password@localhost:5432/dbname?sslmode=disable`).

### 3. Khởi chạy ứng dụng
Chạy lệnh sau để khởi động server:
```bash
go run cmd/server/main.go
```
*Lưu ý: Ứng dụng sẽ tự động chạy các file migration khi khởi động lần đầu.*

## Tài liệu 2 API chính

Hệ thống cung cấp 2 luồng API chính: một cho các thiết bị IoT gửi dữ liệu và một cho ứng dụng Web lấy dữ liệu hiển thị.

### 1. API Gửi dữ liệu cảm biến (Dành cho thiết bị IoT)

Dùng để các thiết bị cảm biến gửi dữ liệu định kỳ về server.

- **Endpoint:** `POST /IOT-api/device/sensor-data`
- **Xác thực:** Yêu cầu Header `IOT-API-Key`.
- **Định dạng Body (JSON):**
  ```json
  {
    "rain_level": 15.5,
    "light": 350.0,
    "soil_moisture": 45.2,
    "ph": 6.5
  }
  ```
- **Phản hồi thành công:** `201 Created`
  ```json
  {
    "message": "sensor data saved",
    "created_at": "2023-10-27T10:00:00Z"
  }
  ```

### 2. API Lấy dữ liệu cảm biến (Dành cho Web Frontend)

Dùng để lấy lịch sử dữ liệu của một thiết bị cụ thể.

- **Endpoint:** `GET /IOT-api/web/devices/:deviceID/sensor-data`
- **Tham số:**
  - `deviceID` (Path): UUID của thiết bị.
  - `number` (Query - Tùy chọn): Số lượng bản ghi muốn lấy (mặc định là 1). VD: `?number=10`.
- **Phản hồi thành công:** `200 OK`
  ```json
  {
    "message": "get data successful",
    "data": [
      {
        "rain_level": 15.5,
        "light": 350.0,
        "soil_moisture": 45.2,
        "ph": 6.5,
        "created_at": "2023-10-27T10:00:00Z"
      }
    ]
  }
  ```

## Cách sử dụng

1. **Đăng ký thiết bị:** Bạn cần có thông tin thiết bị trong bảng `devices` cùng với một `api_key` hợp lệ.
2. **Gửi dữ liệu:** Sử dụng Postman hoặc thư viện HTTP trên vi điều khiển (ESP32/Arduino) để gửi POST request kèm `IOT-API-Key`.
3. **Xem dữ liệu:** Truy cập thông qua endpoint web để lấy dữ liệu mới nhất phục vụ việc vẽ biểu đồ hoặc hiển thị lên Dashboard.
