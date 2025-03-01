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

## 📜 Loyiha tuzilishi

```
project-name/
├── Makefile
├── api/
│   ├── controllers/
│   │   ├── auth.go
│   │   ├── controllers.go
│   │   └── root.go
│   ├── middlewares/
│   │   └── middlewares.go
│   └── routes/
│       ├── auth.go
│       ├── root.go
│       └── routes.go
├── assets/
├── bin/
│   ├── api
│   └── app
├── bootstrap/
│   ├── app.go
│   └── env.go
├── cmd/
│   └── main.go
├── domain/
│   └── auth.go
├── go.mod
├── go.sum
├── internal/
├── pkg/
├── repository/
│   └── repository.go
├── usecase/
│   └── usecase.go
├── .env
├── Dockerfile
├── README.md
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

