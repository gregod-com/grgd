package helper

import (
	"grgd/interfaces"
	"os"
)

// Helper ...
type Helper struct{}

// CheckUserProfile ...
func (h *Helper) CheckUserProfile(logger interfaces.ILogger) string {
	var profilename string
	u, ok := os.LookupEnv("USER")
	if !ok {
		logger.Fatal("failed to lookup USER ENV VAR! Exiting")
	}

	if u == "root" {
		logger.Fatal("You should not run this app as root! Exiting")
	}

	profilename = u

	if p := h.CheckFlagArg("profile"); p != "" {
		profilename = p
	}
	logger.Trace("Found profile")
	return profilename
}

// CheckFlagArg ...
func (h *Helper) CheckFlagArg(flag string) string {
	for k, v := range os.Args {
		if v == "--"+flag && len(os.Args) > k+1 {
			return os.Args[k+1]
		}
	}
	return ""
}

// CheckFlag ...
func (h *Helper) CheckFlag(flag string) bool {
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
