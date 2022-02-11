package api

type Application struct {
	ID                  Snowflake        `json:"id"`
	Name                string           `json:"name"`
	Icon                *string          `json:"icon"`
	Description         string           `json:"description"`
	RPCOrigins          []string         `json:"rpc_origins"`
	BotPublic           bool             `json:"bot_public"`
	BotRequireCodeGrant bool             `json:"bot_require_code_grant"`
	TermsOfServiceURL   string           `json:"terms_of_service_url,omitempty"`
	PrivacyPolicyURL    string           `json:"privacy_policy_url,omitempty"`
	Owner               User             `json:"owner,omitempty"`
	Summary             string           `json:"summary"`
	VerifyKey           string           `json:"verify_key"`
	Team                *Team            `json:"team"`
	GuildID             Snowflake        `json:"guild_id,omitempty"`
	PrimarySkuID        Snowflake        `json:"primary_sku_id"`
	Slug                string           `json:"slug,omitempty"`
	CoverImage          string           `json:"cover_image,omitempty"`
	Flags               ApplicationFlags `json:"flags,omitempty"`
}

type ApplicationFlags int64

const (
	GatewayPresence               ApplicationFlags = 1 << 12
	GatewayPresenceLimited        ApplicationFlags = 1 << 13
	GatewayGuildMembers           ApplicationFlags = 1 << 14
	GatewayGuildMembersLimited    ApplicationFlags = 1 << 15
	VerificationPendingGuildLimit ApplicationFlags = 1 << 16
	Embedded                      ApplicationFlags = 1 << 17
	GatewayMessageContent         ApplicationFlags = 1 << 18
	GatewayMessageContentLimited  ApplicationFlags = 1 << 19
)
