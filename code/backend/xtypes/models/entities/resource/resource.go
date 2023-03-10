package resource

// syncme => entities/consts/resources.ts

const (
	SockRoom   = "sroom"
	Dtable     = "dtable"
	Dgroup     = "dgroup"
	Folder     = "cfolder"
	UserGroup  = "ugroup"
	Module     = "module"
	Repo       = "repo"
	BprintType = "bgroup"
	Lock       = "lock"
	EnvMapper  = "envmapper"
)

type Type struct {
	Name        string
	NeedsTarget bool
}

var Types []Type

func init() {

	Types = []Type{
		{Name: SockRoom, NeedsTarget: true},
		{Name: Dtable, NeedsTarget: true},
		{Name: Dgroup, NeedsTarget: true},
		{Name: Folder, NeedsTarget: true},
		{Name: UserGroup, NeedsTarget: true},
		{Name: Repo, NeedsTarget: true},
		{Name: BprintType, NeedsTarget: true},
		{Name: Lock, NeedsTarget: true},
		{Name: EnvMapper, NeedsTarget: true},
	}

}
