
# 🏢 Simple ERP Backend System

Sebuah sistem ERP (Enterprise Resource Planning) sederhana yang dibangun menggunakan bahasa Go (Golang). Project ini dirancang untuk memanajemen modul-modul umum seperti pengguna, produk, inventaris, dan transaksi.

## 🚀 Fitur Utama

- ✅ Autentikasi dan Otorisasi JWT
- 📦 Manajemen Produk dan Stok
- 👥 Modul Pengguna dan Role
- 📊 Modul Transaksi dan Laporan
- 💾 Database: MySQL
- ⚡ Caching: Redis (opsional)
- 🧩 Clean Architecture & Modular Design
- 🐳 Dockerized

## 📂 Struktur Proyek

```
simple-erp/
├── cmd/                    # Entry point aplikasi
├── config/                 # Konfigurasi aplikasi (.env, database, dll)
├── internal/
│   ├── domain/             # Interface dan entity (DDD)
│   ├── service/            # Logika bisnis
│   ├── repository/         # Implementasi data access
│   └── handler/            # HTTP handler (controller)
│   └── dto/                # Request/Response model (API-safe)
├── pkg/                    # Library atau utilitas umum
├── migrations/             # Skrip migrasi database
├── Dockerfile              # Docker image build
├── docker-compose.yml      # Layanan Docker (db, redis, app)
├── go.mod / go.sum         # Modul Go
├── .env.example            # Contoh file konfigurasi
└── README.md
```

## 📦 Teknologi yang Digunakan

- **Golang**
- **Fiber** sebagai web framework
- **MySQL** untuk database relasional
- **Redis** untuk caching/token (opsional)
- **JWT** untuk autentikasi
- **GORM** untuk ORM/Query Builder
- **Docker** untuk containerisasi

## 📄 Cara Menjalankan (Local)

```bash
# 1. Clone repositori
git clone https://github.com/johanagus/simple-erp.git
cd simple-erp

# 2. Copy konfigurasi
cp .env.example .env

# 3. Jalankan Docker
docker-compose up --build
```

## Default Login

Saat pertama kali aplikasi dijalankan, data user dummy akan otomatis terbuat melalui proses seeder.  
Gunakan akun berikut untuk login ke sistem:

- **Email:** admin@simple.erp
- **Password:** admin123

Anda dapat mengubah data ini di file seeder sesuai kebutuhan.

## 🧪 Endpoint API

```http

POST /api/v1/auth/signin
POST /api/v1/auth/signout
POST /api/v1/auth/token/refresh

GET /api/v1/users
GET /api/v1/user/{id}
UPDATE /api/v1/user/{id}
DELETE /api/v1/user/{id}
POST /api/v1/user

GET /api/v1/products
GET /api/v1/product/{id}
UPDATE /api/v1/product/{id}
DELETE /api/v1/product/{id}
POST /api/v1/product

```

Gunakan link Postman berikut untuk eksplorasi endpoint. 
https://www.postman.com/johanagus/workspace/simple-pos/folder/9335228-1083d586-aa3d-4cdb-beb7-08461ed4ad2d?action=share&creator=9335228&ctx=documentation&active-environment=9335228-6f10e50b-a00a-452d-a526-083e27f22e95


## 📈 Rencana Pengembangan

- [x] Modul Autentikasi
- [x] Modul Users
- [ ] Modul Produk & Inventaris
- [ ] Modul Transaksi
- [ ] Laporan PDF/Excel
- [ ] Integrasi ke frontend React/NextJS

## 👨‍💻 Kontribusi

Pull Request terbuka untuk siapa saja. Fork dan buat branch baru!

---

## 📮 Kontak

Dibuat dengan ❤️ oleh **Johan Agus Saputro**  
[LinkedIn](https://www.linkedin.com/in/johan-agus/) • [GitHub](https://github.com/johanagus)
