package fieldlabs

type Action int

const (
	ActionCreate Action = iota
	ActionDestroy
)

type Params struct {
	NamePrefix       string
	Action           Action
	EnvironmentsJSON string
	// alternative to JSON, handy for exports from sheets / google forms
	EnvironmentsCSV string
	LabsJSON        string

	InstanceJSONOutput string

	APIToken      string
	APIOrigin     string
	GraphQLOrigin string
	KURLSHOrigin  string
	IDOrigin      string

	// invite members based on Environment.Email
	InviteUsers bool
	// vendor web user's email for sending invites, can't use api token
	InviterEmail string
	// vendor web user's password for sending invites, can't use api token
	InviterPassword string
	// sessionToken
	SessionToken string
}
