package internal

type CorsCfg struct {
	Methods     []string `toml:"methods"`
	Origins     []string `toml:"urls"`
	Headers     []string `toml:"headers"`
	Credentials bool     `toml:"credentials"`
	Debug       bool     `toml:"api"`
}

type ServerCfg struct {
	ServiceName        string `toml:"service_name"`
	BindAddr           string `toml:"bind_addr_http"`
	ReadTimeout        int    `toml:"read_timeout"`
	WriteTimeout       int    `toml:"write_timeout"`
	Protocol           string `toml:"protocol"`
	FileTLSCertificate string `toml:"tls_certificate_file"`
	FileTLSKey         string `toml:"tls_key_file"`
}
