# golang-final-project3-team2

Link Deploy API : https://golang-final-project3-team2-production.up.railway.app/

Repository Team 2 Untuk Final Project 3 Golang (Kampus Merdeka Hacktiv8)

Anggota Kelompok :

- JOVIN LIDAN (GLNG-KS04-023)
  Mengerjakan :
  - Setup Docker dan Init Project
  - Semua API User dan Postman user
  - Semua API Categories dan Postman categories
  - Setup deployment menggunakan railway
  - Helper : GenerateToken, VerifyToken, ValidateRequest, ComparePass, HashPass
  - Helper_test : TestSuccessGenerateToken, TestFailedGenerateToken, TestSuccessComparePass, TestFailedComparePass, TestSuccessHashPass, TestFailedHashPass
- GUSTIO NUSAMBA (GLNG-KS04-025)
  Mengerjakan :
    - Semua API Task dan Postman task

## Cara Install

1. Buka dan jalankan aplikasi docker.
2. Jalankan command `docker compose up --build` untuk menjalankan database postgres di dalam docker container , go dan air auto reload. Tunggu agar docker sudah berjalan dengan baik.
3. Setelah docker container semuanya berjalan dengan baik. Maka port default yang akan dibuka adalah `8080`

_Note : Memerlukan docker terinstall didalam perangkat anda_

_Nama File Postman : `Kanban.postman_collection.json`_

```json
Akun Admin:
Email     : admin@gmail.com
Password  : admin12
```

## List Route
### Users
- **`POST`- Users Register `api/users/register`**, Digunakan untuk membuat user baru.
- **`POST`- Users Login `api/users/login`**, Digunakan untuk melakukan login atau autentikasi user.
- **`PUT`- Users Update `api/users/update-account`**, Digunakan untuk mengubah data user berdasarkan idnya.
- **`DELETE`- Users Delete `api/users/delete-account`**, Digunakan untuk menghapus data user.

### Categories
- **`GET`- Categories Index `api/categories`**, Digunakan untuk mengambil seluruh data categories dari database.
- **`POST`- Categories Store `api/categories`**, Digunakan untuk membuat category baru.
- **`PATCH`- Categories Update `api/categories/:categoryId`**, Digunakan untuk mengubah data category berdasarkan idnya.
- **`DELETE`- Categories Delete `api/categories/:categoryId`**, Digunakan untuk menghapus data category berdasarkan idnya.

### Tasks
- **`GET`- Tasks Index `api/tasks`**, Digunakan untuk mengambil seluruh data tasks dari database.
- **`POST`- Tasks Store `api/tasks`**, Digunakan untuk membuat task baru.
- **`PUT`- Tasks Update `api/tasks/:taskId`**, Digunakan untuk mengubah data task berdasarkan idnya.
- **`PATCH`- Tasks Update Status `api/tasks/update-status/:taskId`**, Digunakan untuk mengubah status task berdasarkan idnya.
- **`PATCH`- Tasks Update Category `api/tasks/update-category/:taskId`**, Digunakan untuk mengubah category task berdasarkan idnya.
- **`DELETE`- Tasks Delete `api/tasks/:taskId`**, Digunakan untuk menghapus data task berdasarkan idnya.
