package main

import (
	"fmt"
	"io"
	"os"
	"poc/sftp-client/config"
	"poc/sftp-client/sftp"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		panic(err)
	}

	config := config.New()

	client := sftp.New(config.SFTP)

	defer client.Close()

	remoteFile := config.SFTP.RemotePath + "/mock.file"
	localfile := "mock.file"

	// List file in remote directory
	files, err := client.ListFiles(config.SFTP.RemotePath)
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		var name, modTime, size string

		name = f.Name()
		modTime = f.ModTime().Format("2006-01-02 15:04:05")
		size = fmt.Sprintf("%12d", f.Size())

		if f.IsDir() {
			name = name + "/"
			modTime = ""
			size = "PRE"
		}
		// Output each file name and size in bytes
		fmt.Fprintf(os.Stdout, "%19s %12s %s\n", modTime, size, name)
	}

	// Download remote file.
	file, err := client.Download(remoteFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	f, err := os.Create(localfile)
	if err != nil {
		panic(err)
	}
	io.Copy(f, file)

	// Download as file
	err = client.DownloadToFile(remoteFile, localfile)
	if err != nil {
		panic(err)
	}

	// Get file info
	info, err := client.Info(remoteFile)
	if err != nil {
		panic(err)
	}
	fmt.Println(info.Name())

	// Upload file
	fileToUpload, err := os.Open(localfile)
	if err != nil {
		panic(err)
	}
	defer fileToUpload.Close()

	var reader io.Reader = fileToUpload
	client.Upload(reader, remoteFile)

	// Upload from file
	err = client.UploadFromFile(localfile, remoteFile)
	if err != nil {
		panic(err)
	}
}
