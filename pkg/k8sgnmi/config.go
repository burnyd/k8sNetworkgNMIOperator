package k8sgnmi

type GNMI_CFG struct {
	Token      string
	Url        string
	Server     string
	Origin     string
	Path       string
	Mode       string
	StreamMode string
	Addr       string
	Username   string
	Password   string
	Host       string
	Port       string
	OcModel    string
	JsonData   []byte
}
