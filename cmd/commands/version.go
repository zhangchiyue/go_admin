package commands

var (
	VersionMsg *VersionInfo
)

type VersionInfo struct {
	Tag       string `json:"tag"`
	BuildTime string `json:"build_time"`
	Commit    string `json:"commit"`
	Author    string `json:"author"`
}
