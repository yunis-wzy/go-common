package env

import "os"

const (
	PSMUnknown     = "-"
	ENVBOE         = "boe"
	ENVBOEI18N     = "boei18n"
	ENVPPE         = "ppe"
	ENVPPEI18N     = "ppei18n"
	ENVPROD        = "prod"
	ENVPRODI18N    = "prodi18n"
	ENVSTAGING     = "staging"
	ENVSTAGINGI18N = "stagingi18n"

	IDCCN   = "cn"
	IDCI18N = "i18n"
)

var psm string
var env string
var idc string
var cluster string

func init() {
	psm = os.Getenv("TCE_PSM")
	if psm == "" {
		psm = PSMUnknown
	}
	env = os.Getenv("ENV")
	if env != ENVBOE && env != ENVPROD && env != ENVPPE && env != ENVSTAGING && env != ENVPRODI18N && env != ENVBOEI18N && env != ENVPPEI18N && env != ENVSTAGINGI18N {
		env = ENVBOE
	}
	idc = os.Getenv("IDC")
	if idc != IDCCN && idc != IDCI18N {
		idc = IDCCN
	}
}

// PSM .
func PSM() string {
	return psm
}

// SetPSM is used for unit test.
func SetPSM(psm_ string) {
	psm = psm_
}

func Env() string {
	return env
}

func IDC() string {
	return idc
}
