package account

type Authenticator interface {
	AccountPermissions(account *Account) ([]*string, error)
	AccountGroups(account *Account) ([]*Group, error)

	Login(account *Account) (token string, err error)

	Groups(filter GroupFilter) ([]*Group, error)
	Group(id string) (*Group, error)
	CreateGroup(role *Group) error
	UpdateGroup(role *Group) error

	Accounts(filter AccountFilter) ([]*Account, error)
	Account(id string) (*Account, error)
	CreateAccount(a *Account) error
	UpdateAccount(a *Account) error

	JoinGroup(g *Group, a *Account) error
	LeaveGroup(g *Group, a *Account) error

	ModificationAllowed() bool
}
