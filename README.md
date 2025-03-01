# Golang + Echo Blueprint

Bu loyiha [Golang](w) va [Echo](w) framework'idan foydalangan holda web API yaratish uchun blueprint sifatida xizmat qiladi.

## 📌 Talablar

- [Go](w) 1.18+
- [Echo](w) framework
- [Docker](w) (ixtiyoriy)

## 🚀 O'rnatish

```sh
# Repository'ni klonlash
git clone https://github.com/username/project-name.git
cd project-name

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
docker build -t my-echo-app .
docker run -p 8080:8080 my-echo-app
```

## 📜 Loyiha tuzilishi

```
project-name/
│── main.go
│── go.mod
│── go.sum
├── config/
│   ├── config.go
├── handlers/
│   ├── user_handler.go
├── models/
│   ├── user.go
├── routes/
│   ├── routes.go
├── middleware/
│   ├── auth.go
├── .env
├── Dockerfile
├── README.md
```

## 🔥 API Endpointlar

| Yo'l | Metod | Tavsif |
|------|------|--------|
| `/health` | GET | API holatini tekshirish |
| `/users` | GET | Barcha foydalanuvchilar ro'yxati |
| `/users/:id` | GET | Foydalanuvchi ma'lumotlari |
| `/users` | POST | Yangi foydalanuvchi qo'shish |

## ✅ Qo'shimcha

- `.env` faylidan konfiguratsiyalarni o'qish
- Middleware qo‘llab-quvvatlash
- Docker bilan oson deploy qilish

## 📄 Litsenziya

Ushbu loyiha [MIT](w) litsenziyasi ostida tarqatiladi.

