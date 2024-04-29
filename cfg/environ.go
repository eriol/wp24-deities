package cfg // import "github.com/eriol/wp24-deities/cfg"

import "os"

const (
	clientId     = "WP24_DEITIES_CLIENT_ID"
	clientSecret = "WP24_DEITIES_CLIENT_SECRET"
	domain       = "WP24_DEITIES_DOMAIN"
)

func GetClientId() string {
	return os.Getenv(clientId)
}

func GetClientSecret() string {
	return os.Getenv(clientSecret)
}

func GetDomain() string {
	return os.Getenv(domain)
}
