<div align="center">
   <h1>
      Quiz 3 Sanbercode golang batch-58 Agus Prasetyo :boom:
   </h1>
</div>

## List Package / Library
- `config` : Digunakan untuk menampung data environtment / `.env` yang digunakan untuk menghubungkan project ke database
- `controllers` : Digunakan untuk mengontrol dan menangani permintaan yang diberikan oleh user, contoh : mengontrol untuk permintaan authentifikasi, permintaan menambahkan, mengedit dan menghapus data dan lain-lain
- `database` : Digunakan untuk menambah tabel migration
- `middleware` : Digunakan untuk penjembatan data dan sebagai pengamanan sesuai dengan mekanisme yang di tentukan
- `model` : Digunakan untuk menampung data kolom sesuai dengan tabel yang dibuat
- `repository` : Digunakan untuk jembatan antara `model` dan `controller`
- `services` : Digunakan untuk menyediakan layanan / service untuk mendukung fitur yang dibuat, untuk case sekarang digunakan sebagai layanan otentifikasi pengguna

## Dummy Authentifikasi
Untuk dapat menjalankan dummy project, dapat menggunakan user dummy sebagai berikut : <br>
```
Admin
username : admin
password : admin

User 1
username : user1
password: user1
```

## Endpoint API
### Login
#### HTTP Request
```json
   POST http://localhost:8080/login
```
#### Parameters
| Parameters    |               | Description  |
| ------------- |:-------------:| ------------- |

#### Result
| Parameters    |  Description  |
| ------------- |:--------------|

```json
   {
      "id": 1,
      "username": "admin",
      "password": "jGl25bVBBBW96Qi9Te4V37Fnqchz/Eu4qB9vKrRIqRg=",
      "created_at": "2024-08-04T00:00:00Z",
      "created_by": null,
      "modified_at": null,
      "modified_by": null
   }
```

### Category

### Book
