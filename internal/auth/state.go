package auth

type State struct {
	Id            string `json:"id"`
	RedirectURL   string `json:"redirectUrl"`
	AuthProvider  string `json:"authProvider"`
	GroupsEnabled bool   `json:"groupsEnabled"`
	PersistLogin  bool   `json:"persistLogin"`
}
