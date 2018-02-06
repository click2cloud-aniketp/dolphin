// +build !windows

package cli

const (
	defaultBindAddress     = ":9000"
	defaultDataDirectory   = "/data"
	defaultAssetsDirectory = "."
	defaultNoAuth          = "false"
	defaultNoAnalytics     = "false"
	defaultTLSVerify       = "false"
	defaultTLSCACertPath   = "/certs/ca.pem"
	defaultTLSCertPath     = "/certs/cert.pem"
	defaultTLSKeyPath      = "/certs/key.pem"
	defaultSSL             = "false"
	defaultSSLCertPath     = "/certs/dockm.crt"
	defaultSSLKeyPath      = "/certs/dockm.key"
	defaultSyncInterval    = "60s"
)
