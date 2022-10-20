package sftp

type Config struct {
	Server            string
	Host              string
	Port              string
	Username          string
	Password          string
	PrivateKey        string
	KeyExchanges      []string
	RemotePath        string
	ConnectionTimeout int
}
