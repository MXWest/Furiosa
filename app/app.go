package app

/** Furiosa Structs and Interfaces **/

type HealthStatusMsg struct {
	Status  string `json:"state"`
	Message string `json:"message"`
}

/** Capable Structs and Interfaces **/

type SlackStruct struct {
	Enabled  bool     `yaml:"enabled"`
	Channels []string `yaml:"channels"`
}

type ConfigStruct struct {
	Org                        string      `yaml:"org"`
	Repositories               []string    `yaml:"repositories"`
	IgnoredRepositories        []string    `yaml:"ignored_repositories"`
	TrivialCommits             []string    `yaml:"trivial_commits"`
	Slack                      SlackStruct `yaml:"slack"`
}

type ArgsStruct struct {
	DraftRelease      bool
	PreRelease        bool
	CutRelease        bool
	Hotfix            bool
	ReleaseBranchName string
	NoSuffix          bool
	Suffix            string
	Repository        string
	SemverBump        string
	BranchFromName    string
	BranchFromSha     string
}
