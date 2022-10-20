package sftp

type Config struct {
	Server            string
	Host              string
	Port              string
	Username          string
	Password          string
	PrivateKey        string
	RemotePath        string
	ConnectionTimeout int
}
