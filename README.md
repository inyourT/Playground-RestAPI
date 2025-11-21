# Simple Users API (Go)

Project ini adalah contoh implementasi **REST API sederhana** menggunakan bahasa Go dengan arsitektur terstruktur (handler → service → repository → model).  
API ini menyediakan endpoint untuk mengambil semua user dan membuat user baru menggunakan penyimpanan in-memory (slice).

---

## 🚀 Features

- GET `/users` – mengambil semua user
- POST `/users` – membuat user baru
- Root `/` – mengecek apakah API berjalan
- Arsitektur terpisah: handler, service, repository, model, dan server
- Penyimpanan sementara menggunakan slice (tanpa database)
