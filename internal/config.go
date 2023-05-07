package internal

type CorsCfg struct {
	Methods     []string
	Origins     []string
	Headers     []string
	Credentials bool
	Debug       bool
}

type ServerCfg struct {
	ServiceName        string
	BindAddrHTTP       string
	ReadTimeout        int
	WriteTimeout       int
	Protocol           string
	FileTLSCertificate string
	FileTLSKey         string
}

type MetricsCfg struct {
	BindAddr   string
	SourceName string
}
