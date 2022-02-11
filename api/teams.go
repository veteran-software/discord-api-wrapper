package api

type Team struct {
	Icon        *string      `json:"icon"`
	ID          Snowflake    `json:"id"`
	Members     []TeamMember `json:"members"`
	Name        string       `json:"name"`
	OwnerUserID Snowflake    `json:"owner_user_id"`
}

type TeamMember struct {
	MembershipState MembershipState `json:"membership_state"`
	Permissions     []string        `json:"permissions"`
	TeamID          Snowflake       `json:"team_id"`
	User            User            `json:"user"`
}

type MembershipState int

const (
	Invited MembershipState = iota + 1
	Accepted
)
