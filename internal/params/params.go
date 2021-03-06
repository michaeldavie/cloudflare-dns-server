package params

import (
	"github.com/qdm12/cloudflare-dns-server/internal/models"
	"github.com/qdm12/golibs/logging"
	libparams "github.com/qdm12/golibs/params"
	"github.com/qdm12/golibs/verification"
)

// ParamsReader contains methods to obtain parameters
type ParamsReader interface {
	// DNS getters
	GetProviders() (providers []models.Provider, err error)
	GetPrivateAddresses() (privateAddresses []string)

	// Unbound getters
	GetListeningPort() (listeningPort uint16, err error)
	GetCaching() (caching bool, err error)
	GetVerbosity() (verbosityLevel uint8, err error)
	GetVerbosityDetails() (verbosityDetailsLevel uint8, err error)
	GetValidationLogLevel() (validationLogLevel uint8, err error)

	// Blocking getters
	GetMaliciousBlocking() (blocking bool, err error)
	GetSurveillanceBlocking() (blocking bool, err error)
	GetAdsBlocking() (blocking bool, err error)
	GetUnblockedHostnames() (hostnames []string, err error)
	GetBlockedHostnames() (hostnames []string, err error)
	GetBlockedIPs() (IPs []string, err error)

	// Version getters
	GetVersion() string
	GetBuildDate() string
	GetVcsRef() string
}

type paramsReader struct {
	envParams libparams.EnvParams
	logger    logging.Logger
	verifier  verification.Verifier
}

// NewParamsReader returns a paramsReadeer object to read parameters from
// environment variables
func NewParamsReader(logger logging.Logger) ParamsReader {
	return &paramsReader{
		envParams: libparams.NewEnvParams(),
		logger:    logger,
		verifier:  verification.NewVerifier(),
	}
}
