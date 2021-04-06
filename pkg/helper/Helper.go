package helper

import (
	"os"

	"github.com/gregod-com/grgd/interfaces"
)

// ProvideHelper ...
func ProvideHelper(logger interfaces.ILogger) interfaces.IHelper {
	h := new(Helper)
	logger.Tracef("provide %T", h)
	h.logger = logger
	return h
}

// Helper ...
type Helper struct {
	logger interfaces.ILogger
}

// CheckUserProfile ...
func (h *Helper) CheckUserProfile() string {
	h.logger.Tracef("[pkg/helper/CheckUserProfile]")
	var profilename string
	u, ok := os.LookupEnv("USER")
	if !ok {
		h.logger.Fatal("failed to lookup USER ENV VAR! Exiting")
	}

	if u == "root" {
		h.logger.Fatal("You should not run this app as root! Exiting")
	}

	profilename = u

	if p := h.CheckFlagArg("profile"); p != "" {
		profilename = p
	}
	h.logger.Trace("Found profile")
	return profilename
}

// CheckFlagArg ...
func (h *Helper) CheckFlagArg(flag string) string {
	h.logger.Tracef("[pkg/helper/CheckFlagArg]")
	for k, v := range os.Args {
		if v == "--"+flag && len(os.Args) > k+1 {
			return os.Args[k+1]
		}
	}
	return ""
}

// CheckFlag ...
func (h *Helper) CheckFlag(flag string) bool {
	h.logger.Tracef("[pkg/helper/CheckFlag]")
	for _, v := range os.Args {
		if v == "-"+flag {
			return true
		}
		if v == "--"+flag {
			return true
		}
	}
	return false
}
