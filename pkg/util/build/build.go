package build

import (
	"encoding/json"
	"net/http"
	"runtime"

	"github.com/prometheus/common/version"
	prom "github.com/prometheus/prometheus/web/api/v1"
)

// Version information passed to Prometheus version package.
// Package path as used by linker changes based on vendoring being used or not,
// so it's easier just to use stable Pyroscope path, and pass it to
// Prometheus in the code.
var (
	Version   string
	Revision  string
	Branch    string
	BuildUser string
	BuildDate string
	GoVersion string
)

func init() {
	version.Version = Version
	version.Revision = Revision
	version.Branch = Branch
	version.BuildUser = BuildUser
	version.BuildDate = BuildDate
	version.GoVersion = runtime.Version()
}

func GetVersion() prom.PrometheusVersion {
	return prom.PrometheusVersion{
		Version:   version.Version,
		Revision:  version.Revision,
		Branch:    version.Branch,
		BuildUser: version.BuildUser,
		BuildDate: version.BuildDate,
		GoVersion: version.GoVersion,
	}
}

func BuildInfoHandler(rw http.ResponseWriter, req *http.Request) {
	resp := struct {
		Status string      `json:"status"`
		Data   interface{} `json:"data,omitempty"`
	}{
		Status: "success",
		Data:   GetVersion(),
	}
	rw.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(rw).Encode(&resp)
}
