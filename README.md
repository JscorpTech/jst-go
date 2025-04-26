# Golang + Echo Blueprint

Bu loyiha [Golang](w) va [Echo](w) framework'idan foydalangan holda web API yaratish uchun blueprint sifatida xizmat qiladi.

## 📌 Talablar

- [Go](w) 1.18+
- [Echo](w) framework
- [Docker](w) (ixtiyoriy)

## 🚀 O'rnatish

```sh
# Repository'ni klonlash
git clone https://github.com/JscorpTech/jst-go.git
cd jst-go

# Modullarni yuklash
go mod tidy
```

## 🏗 Ishga tushirish

### Lokal muhitda

```sh
go run main.go
```

### Docker orqali

```sh
docker build -t jst-go .
docker run -p 8080:8080 jst-go
```

## 🔥 API Endpointlar

| Yo'l         | Metod | Tavsif                           |
| ------------ | ----- | -------------------------------- |
| `/health`    | GET   | API holatini tekshirish          |
| `/users`     | GET   | Barcha foydalanuvchilar ro'yxati |
| `/users/:id` | GET   | Foydalanuvchi ma'lumotlari       |
| `/users`     | POST  | Yangi foydalanuvchi qo'shish     |

## ✅ Qo'shimcha

- `.env` faylidan konfiguratsiyalarni o'qish
- Middleware qo‘llab-quvvatlash
- Docker bilan oson deploy qilish

## 📄 Litsenziya

Ushbu loyiha [MIT](w) litsenziyasi ostida tarqatiladi.


Here is an organized summary of the provided code snippets, categorized by their functionalities:

### 1. **Utility Functions**
   - **HTTP Responses**: The `utils/http.go` file provides functions to create success and error responses using a DTO
(Domain Transfer Object) struct. These are useful for standardizing API responses.

   - **JWT Handling**: The `utils/jwt.go` file includes functions for generating and parsing JSON Web Tokens. It uses Viper
for configuration, requiring an environment variable or config file for the secret key.

   - **Password Management**: The `utils/password.go` file offers functions to hash passwords and verify them using bcrypt.
These are essential for secure user authentication.

### 2. **Services**
   - **SMS Service**: The `services/sms.go` file defines an SMS service with a basic implementation that prints messages to
the console. It can be extended to integrate with real SMS providers.

### 3. **Package Structure and Considerations**
   - Each utility and service is in its own package, promoting modularity.
   - Ensure all dependencies (e.g., Viper, JWT, bcrypt) are properly imported and configured.
   - Implement error handling middleware to manage errors returned by these functions.
   - Organize the project structure with proper directories for commands, internal packages, etc.

### 4. **Testing and Best Practices**
   - Develop comprehensive unit tests for each function, especially those dealing with security aspects like JWT and
password hashing.
   - Follow Go best practices: consistent naming, documentation comments, struct definitions, and formatting.

This setup provides a robust foundation for building a Go application with essential utilities and services. Proper
integration of these components will ensure a secure, efficient, and maintainable codebase.