package config

import (
	"os"
	"poc/sftp-client/config/sftp"
	"strconv"
)

type Config struct {
	SFTP sftp.Config
}

func New() *Config {
	var config Config

	config.SFTP.Host = os.Getenv("SFTP_HOST")
	config.SFTP.Port = os.Getenv("SFTP_PORT")
	config.SFTP.Server = config.SFTP.Host + ":" + config.SFTP.Port
	config.SFTP.Username = os.Getenv("SFTP_USERNAME")
	config.SFTP.Password = os.Getenv("SFTP_PASSWORD")
	config.SFTP.RemotePath = os.Getenv("SFTP_REMOTE_PATH")

	config.SFTP.PrivateKey = os.Getenv("SFTP_PRIVATE_KEY")

	timeout, err := strconv.Atoi(os.Getenv("SFTP_COONECTION_TIMEOUT"))
	if err != nil {
		panic(err)
	}
	config.SFTP.ConnectionTimeout = timeout

	return &config
}
