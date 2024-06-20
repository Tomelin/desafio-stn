package configs

type Certificate struct {
	Key string "mapstructure:'key'"
	Crt string "mapstructure:'crt'"
}

type ConfigWebServer struct {
	Port            string      "mapstructure:'port'"
	Listen          string      "mapstructure:'listen'"
	Http2Enabled    bool        "mapstructure:'http2_enabled'"
	SslEnabled      bool        "mapstructure:'ssl_enabled'"
	CertificatePath Certificate "mapstructure:'certificate_path'"
	Mode            string      "mapstructure:'mode'"
}

func (c *ConfigWebServer) Validate() {
	if c.Port == "" {
		c.Port = "8443"
	}

	if c.Listen == "" {
		c.Listen = "0.0.0.0"
	}
}
