package interfaces

// IHelper ...
type IHelper interface {
	CheckUserProfile() string
	CheckFlag(flag string) bool
	CheckFlagArg(flag string) string
}
