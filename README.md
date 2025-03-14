# MinIO Golang Example

## Langkah-Langkah Membuat Kode
1. **Inisialisasi Proyek**
   ```sh
   go mod init minio-example
   go get github.com/minio/minio-go/v7
   ```
2. **Import Library yang Dibutuhkan**
   - `context`
   - `fmt`
   - `log`
   - `time`
   - `github.com/minio/minio-go/v7`
   - `github.com/minio/minio-go/v7/pkg/credentials`
3. **Buat Koneksi ke MinIO**
   - Gunakan `minio.New()` untuk membuat instance MinIO Client.
4. **Implementasi Fungsi Utama**
   - Upload file (`uploadFile`)
   - List file dalam bucket (`listFiles`)
   - Download file (`downloadFile`)
   - Generate presigned URL (`generatePresignedURL`)
   - Generate public URL (`generatePublicURL`)
   - Hapus file (`deleteFile`)
5. **Jalankan Program**
   ```sh
   go run main.go
   ```

## Penjelasan Kode
Kode ini menghubungkan aplikasi dengan MinIO untuk mengelola file dalam bucket. Fitur utama yang disediakan meliputi upload, download, melihat daftar file, menghapus file, serta pembuatan URL akses (presigned dan public). Semua operasi dilakukan menggunakan MinIO SDK untuk Golang.

