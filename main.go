package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	endpoint := "203.194.113.6:9000"
	accessKeyID := "MidFtK0wfiZ6AUjDfZbz"
	secretAccessKey := "KxkgFNq196ok2AKq9U5h2naOUq0Akpi8HyjA4RO3"
	bucketName := "smk-telkom"

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()

	uploadFile(minioClient, ctx, bucketName, "loh-kok-file.txt", "uploaded-file.txt")
	uploadFile(minioClient, ctx, bucketName, "loh-kok-file.txt", "public/uploaded-file.txt")
	listFiles(minioClient, ctx, bucketName)
	downloadFile(minioClient, ctx, bucketName, "uploaded-file.txt", "downloaded-file.txt")

	presignedURL := generatePresignedURL(minioClient, ctx, bucketName, "uploaded-file.txt", 24*time.Hour)
	fmt.Println("Presigned URL:", presignedURL)

	publicURL := generatePublicURL(endpoint, bucketName, "public/uploaded-file.txt")
	fmt.Println("Public URL:", publicURL)

	deleteFile(minioClient, ctx, bucketName, "uploaded-file.txt")
}

func uploadFile(minioClient *minio.Client, ctx context.Context, bucketName, filePath, objectName string) {
	_, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("File %s berhasil diupload sebagai %s\n", filePath, objectName)
}

func listFiles(minioClient *minio.Client, ctx context.Context, bucketName string) {
	objectCh := minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{})
	for object := range objectCh {
		if object.Err != nil {
			log.Fatalln(object.Err)
		}
		fmt.Println(object.Key)
	}
}

func downloadFile(minioClient *minio.Client, ctx context.Context, bucketName, objectName, filePath string) {
	err := minioClient.FGetObject(ctx, bucketName, objectName, filePath, minio.GetObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("File %s berhasil didownload sebagai %s\n", objectName, filePath)
}

func generatePresignedURL(minioClient *minio.Client, ctx context.Context, bucketName, objectName string, expiry time.Duration) string {
	presignedURL, err := minioClient.PresignedGetObject(ctx, bucketName, objectName, expiry, nil)
	if err != nil {
		log.Fatalln(err)
	}
	return presignedURL.String()
}

func generatePublicURL(endpoint, bucketName, objectName string) string {
	return fmt.Sprintf("http://%s/%s/%s", endpoint, bucketName, objectName)
}

func deleteFile(minioClient *minio.Client, ctx context.Context, bucketName, objectName string) {
	err := minioClient.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("File %s berhasil dihapus\n", objectName)
}
