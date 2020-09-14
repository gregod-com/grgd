package interfaces

// IHelper ...
type IHelper interface {
	CheckUserProfile(logger ILogger) string
	CheckFlag(flag string) bool
	CheckFlagArg(flag string) string
}
