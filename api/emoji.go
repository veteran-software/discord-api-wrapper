package api

import (
	"fmt"
	"net/http"
)

/*
Emoji

ID: emoji id

Name: emoji name

Roles: roles allowed to use this emoji

User: user that created this emoji

RequireColons: whether this emoji must be wrapped in colons

Managed: whether this emoji is managed

Animated: whether this emoji is animated

Available: whether this emoji can be used, may be false due to loss of Server Boosts
*/
type Emoji struct {
	ID            *Snowflake `json:"id"`
	Name          string     `json:"name"`
	Roles         []Role     `json:"roles,omitempty"`
	User          *User      `json:"user,omitempty"`
	RequireColons bool       `json:"require_colons,omitempty"`
	Managed       bool       `json:"managed,omitempty"`
	Animated      bool       `json:"animated,omitempty"`
	Available     bool       `json:"available,omitempty"`
}

// ListGuildEmojis
// Returns a list of emoji objects for the given guild.
func (g *Guild) ListGuildEmojis() (method string, route string) {
	return http.MethodGet, fmt.Sprintf("%s/guilds/%s/emojis", api, g.ID.String())
}

// GetGuildEmoji
// Returns an emoji object for the given guild and emoji IDs.
func (g *Guild) GetGuildEmoji(emoji Emoji) (method string, route string) {
	return http.MethodGet, fmt.Sprintf("%s/guilds/%s/emojis/%s", api, g.ID.String(), emoji.ID.String())
}

func (g *Guild) CreateGuildEmoji() (method string, route string) {
	return http.MethodPost, fmt.Sprintf("%s/guilds/%s/emojis", api, g.ID.String())
}
