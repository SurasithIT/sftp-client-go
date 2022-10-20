package config

import (
	"log"
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

	pk, err := os.ReadFile(os.Getenv("SFTP_PRIVATE_KEY")) // required only if private key authentication is to be used
	if err != nil {
		log.Fatalln(err)
	}

	config.SFTP.PrivateKey = string(pk)
	config.SFTP.KeyExchanges = []string{"diffie-hellman-group-exchange-sha256", "diffie-hellman-group14-sha256"} // optional

	timeout, err := strconv.Atoi(os.Getenv("SFTP_COONECTION_TIMEOUT"))
	if err != nil {
		panic(err)
	}
	config.SFTP.ConnectionTimeout = timeout

	return &config
}
