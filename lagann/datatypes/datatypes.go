package datatypes

type User struct {
	Name     string    `json:"name"`
	Password string    `json:"password"`
	SSHKeys  []*SSHKey `json:"sshkeys"`
}

type SSHKey struct {
	Comment     string `json:"comment"`
	Key         string `json:"key"`
	Fingerprint string `json:"fingerprint"`
}

type App struct {
	Users     []string `json:"user"`
	Name      string   `json:"name"`
	LastBuild *Build   `json:"last_build"`
}

type Build struct {
	User  string `json:"user"`
	ID    string `json:"id"`
	Image string `json:"image"`
	App   string `json:"app"`
}

type Reply struct {
	Message string                   `json:"message"`
	Code    int                      `json:"code"`
	Data    []map[string]interface{} `json:"data"`
}
