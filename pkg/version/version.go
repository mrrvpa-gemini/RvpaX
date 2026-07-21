package version

const (
	Version = "1.0.0"
	Build   = "dev"
	Author  = "mrrvpa-gemini"
	Github  = "https://github.com/mrrvpa-gemini/rvpa"
)

func GetVersion() string {
	return Version
}

func GetBuild() string {
	return Build
}
