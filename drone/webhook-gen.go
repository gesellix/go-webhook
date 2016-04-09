package drone

type Drone struct {
	Docker struct {
		Images []struct {
			Tag string `json:"tag"`
		} `json:"images"`
	} `json:"docker"`
	ExtraVars struct {
		AnotherToken string `json:"another_token"`
		Env          string `json:"env"`
		GitBranch    string `json:"git_branch"`
	} `json:"extra_vars"`
	Name string `json:"name"`
}
