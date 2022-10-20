package sftp

import (
	sftpConfig "poc/sftp-client/config/sftp"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type Client struct {
	config     sftpConfig.Config
	sshClient  *ssh.Client
	sftpClient *sftp.Client
}
