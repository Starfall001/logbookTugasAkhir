# logbookTugasAkhir
Aplikasi ini dibuat untuk memenuhi matakuliah Teknologi Informasi Kesehatan
##### Hal yang sudah dibuat
- Backend dari sistem berupa
  - CRUD user register
  - CRUD login
  - CRUD logbook
  - CRUD karya
- UI
---
##### Kekurangan dari backend
- User role antara mahasiswa dan dosen bagaimana?
- Belum terintegrasi dengan database kampus
- Belum bisa upload file pada formulir bimbingan
- Belum membuat crud buat komen pada formulir bimbingan
---
##### Kekurangan dari front end
- Masih butuh perbaikan UI
- Masih belum dibuat front end nya

Dokumentasi Api

List dokumentasi
- Bagaimana cara membuat CRUD endpoint
- Bagaimana cara dokumentasi API menggunakan Swagger-UI and OpenApi 3.0.0.


### Project layout
```
|- be/
| |
| |- app            // menjalankan server, database, dan migration
| |- config         // isi dari server, database, dan migration
| |- controllers    // berisi REST corollers (User, logbook, dan karya)
| |- database       // berisi variabel yang menampung library ormnya go yaitu gorm
| |- middleware     // berisi authentikasi token
| |- models         // berisi dokumentasi dari database, request create and update, dan response
| |- routes         // berisi routes
| |- utils          // berisi token generator dan hashing password 
| |- .env           // dokumentasi env
| |- .env.example   // dokumentasi env.example
| |- go.mod         // menampung semua module projek
| |- main.go        // berisi main()
|- fe/
```
---

#### Menjalankan projek
```
$ go run main.go    // posisikan di folder be
```

---
#### User Service

|HTTP Method|URL|Description|
|---|---|---|
|`POST`|http://127.0.0.1:8000/api/v1/user/register | Create new User |
|`POST`|http://127.0.0.1:8000/api/v1/user/login | User Login |
---

#### Karya Service
|HTTP Method|URL|Description|
|---|---|---|
|`POST`|http://127.0.0.1:8000/api/v1/karya | Create new karya |
|`GET`|http://127.0.0.1:8000/api/v1/karya | Get all karya |
|`GET`|http://127.0.0.1:8000/api/v1/karya/{karyaId} | Get karya by id |
|`PATCH`|http://127.0.0.1:8000/api/v1/karya/{karyaId} | Update karya by id |
|`DELETE`|http://127.0.0.1:8000/api/v1/karya/{karyaId} | Delete karya |

---

#### Logbook Service
|HTTP Method|URL|Description|
|---|---|---|
|`POST`|http://127.0.0.1:8000/api/v1/logbook | Create new logbook |
|`GET`|http://127.0.0.1:8000/api/v1/logbook | Get all logbook |
|`GET`|http://127.0.0.1:8000/api/v1/logbook/{logbookId} | Get logbook by id |
|`PATCH`|http://127.0.0.1:8000/api/v1/logbook/{logbookId} | Update logbook by id |
|`DELETE`|http://127.0.0.1:8000/api/v1/logbook/{logbookId} | Delete logbook |
---


