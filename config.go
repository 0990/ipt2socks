package tproxy2socks

type Config struct {
	Proxy      string
	ListenAddr string
	UDPTimeout int32
}
