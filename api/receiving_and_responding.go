package api

/* INTERACTION OBJECT */

/*
Interaction

https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-interaction-structure

An interaction is the base "thing" that is sent when a user invokes a command, and is the same for Slash Commands and other future interaction types (such as Component).

Interaction.Data is always present on ApplicationCommand interaction types. It is optional for future-proofing against new interaction types

Interaction.Member is sent when the command is invoked in a guild, and Interaction.User is sent when invoked in a DM

--------

ID: id of the interaction

ApplicationID: id of the application this interaction is for

Type: the type of interaction

Data: the command data payload

GuildID: the guild it was sent from

ChannelID: the channel it was sent from

Member: guild member data for the invoking user, including permissions

User: user object for the invoking user, if invoked in a DM

Token: a continuation token for responding to the interaction

Version: read-only property, always 1

Message: for components, the message they were attached to
*/
type Interaction struct {
	ID            Snowflake       `json:"id"`
	ApplicationID Snowflake       `json:"application_id"`
	Type          InteractionType `json:"type"`
	Data          InteractionData `json:"data,omitempty"`
	GuildID       Snowflake       `json:"guild_id,omitempty"`
	ChannelID     Snowflake       `json:"channel_id,omitempty"`
	Member        GuildMember     `json:"member,omitempty"`
	User          *User           `json:"user,omitempty"`
	Token         string          `json:"token"`
	Version       int             `json:"version"` // read-only property, always `1`
	Message       *Message        `json:"message,omitempty"`
	Locale        string          `json:"locale,omitempty"`
	GuildLocale   string          `json:"guild_locale,omitempty"`
}

/*
InteractionType

https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-interaction-type
*/
type InteractionType int

const (
	InteractionTypePing InteractionType = iota + 1
	InteractionTypeApplicationCommand
	InteractionTypeMessageComponent
	InteractionTypeApplicationCommandAutocomplete
	InteractionTypeModalSubmit
)

/*
InteractionData

https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-interaction-data-structure

ID: the ID of the invoked command

Name: the name of the invoked command

Resolved: converted users + roles + channels

Options: the params + values from the user

CustomID: for components, the custom_id of the component

ComponentType: for components, the type of the component
*/
type InteractionData struct {
	/* Application Command */
	ID       Snowflake                                  `json:"id,omitempty"`
	Name     string                                     `json:"name,omitempty"`
	Type     ApplicationCommandType                     `json:"type,omitempty"`
	Resolved ResolvedData                               `json:"resolved,omitempty"`
	Options  []*ApplicationCommandInteractionDataOption `json:"options,omitempty"`

	/* Component, Modal Submit */
	CustomID string `json:"custom_id,omitempty"`

	/* Component */
	ComponentType ComponentType `json:"component_type,omitempty"`

	/* Component (Select) */
	Values []string `json:"values,omitempty"`

	/* User Command, Message Command */
	TargetID Snowflake `json:"target_id,omitempty"`

	/* Modal Submit */
	Components []Component `json:"components,omitempty"`
}

/*
ResolvedData

https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-resolved-data-structure

Users: the IDs and DiscordUser objects

Members: the IDs and partial GuildMember objects

Roles: the IDs and GuildRole objects

Channels: the IDs and partial GuildChannel objects

Attachments: the ids and attachment objects
*/
type ResolvedData struct {
	Users       map[Snowflake]User        `json:"users,omitempty"`
	Members     map[Snowflake]GuildMember `json:"members,omitempty"`
	Roles       map[Snowflake]Role        `json:"roles,omitempty"`
	Channels    map[Snowflake]Channel     `json:"channels,omitempty"`
	Messages    map[Snowflake]Message     `json:"messages,omitempty"`
	Attachments map[Snowflake]Attachment  `json:"attachments,omitempty"`
}

/* MESSAGE INTERACTION OBJECT */

/*
MessageInteraction

https://discord.com/developers/docs/interactions/receiving-and-responding#message-interaction-object-message-interaction-structure

This is sent on the message object when the message is a response to an Interaction.

--------

ID: id of the interaction

Type: the type of interaction

Name: the name of the ApplicationCommand

User: the user who invoked the interaction
*/
type MessageInteraction struct {
	ID     Snowflake       `json:"id"`
	Type   InteractionType `json:"type"`
	Name   string          `json:"name"`
	User   User            `json:"user"`
	Member GuildMember     `json:"member,omitempty"`
}

/* INTERACTION RESPONSE OBJECT */

/*
InteractionResponseMessages

https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-interaction-response-structure

After receiving an interaction, you must respond to acknowledge it. You can choose to respond with a message immediately using type 4, or you can choose to send a deferred response with type 5. If choosing a deferred response, the user will see a loading state for the interaction, and you'll have up to 15 minutes to edit the original deferred response using Edit Original Interaction Response.

Interaction responses can also be public—everyone can see it—or "ephemeral"—only the invoking user can see it. That is determined by setting flags to 64 on the InteractionCallbackDataMessages.

--------

Type: the type of response

Data: an optional response message
*/
type InteractionResponseMessages struct {
	Type InteractionCallbackType          `json:"type"`
	Data *InteractionCallbackDataMessages `json:"data,omitempty"`
}

/*
InteractionResponseAutocomplete

https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-interaction-response-structure

After receiving an interaction, you must respond to acknowledge it. You can choose to respond with a message immediately using type 4, or you can choose to send a deferred response with type 5. If choosing a deferred response, the user will see a loading state for the interaction, and you'll have up to 15 minutes to edit the original deferred response using Edit Original Interaction Response.

Interaction responses can also be public—everyone can see it—or "ephemeral"—only the invoking user can see it. That is determined by setting flags to 64 on the InteractionCallbackDataMessages.

--------

Type: the type of response

Data: options for the autocomplete result
*/
type InteractionResponseAutocomplete struct {
	Type InteractionCallbackType              `json:"type"`
	Data *InteractionCallbackDataAutocomplete `json:"data,omitempty"`
}

type InteractionResponseModal struct {
	CallbackType InteractionCallbackType       `json:"type"`
	Data         *InteractionCallbackDataModal `json:"data,omitempty"`
}

/*
InteractionCallbackType

https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-interaction-callback-type

Pong: ACK a Ping

ChannelMessageWithSource: respond to an interaction with a message

DeferredChannelMessageWithSource: ACK an interaction and edit a response later, the user sees a loading state

DeferredUpdateMessage: for components, ACK an interaction and edit the original message later; the user does not see a loading state; edit the message using EditOriginalInteractionResponse

UpdateMessage: for components, edit the message the component was attached to
*/
type InteractionCallbackType int

const (
	Pong InteractionCallbackType = iota + 1
	_
	_
	ChannelMessageWithSource
	DeferredChannelMessageWithSource
	DeferredUpdateMessage
	UpdateMessage
	AutocompleteResult
	Modal // ** Not available for MODAL_SUBMIT and PING interactions.
)

/*
InteractionCallbackDataMessages

https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-interaction-callback-data-structure

BUG(KeithRusso): Not all message fields are currently supported.

--------

TTS: is the response TTS

Content: message content

Embeds: supports up to 10 embeds

AllowedMentions: AllowedMentionType object

Flags: set to 64 to make your response Ephemeral

Components: message components
*/
type InteractionCallbackDataMessages struct {
	TTS             bool             `json:"tts"`
	Content         string           `json:"content"`
	Embeds          []Embed          `json:"embeds,omitempty"`
	AllowedMentions *AllowedMentions `json:"allowed_mentions"`
	Flags           MessageFlags     `json:"flags,omitempty"`
	Components      []Component      `json:"components,omitempty"`
	Attachments     []Attachment     `json:"attachments,omitempty"`
}

type InteractionCallbackDataAutocomplete struct {
	Choices []*ApplicationCommandOptionChoice `json:"choices"`
}

type InteractionCallbackDataModal struct {
	CustomID   string      `json:"custom_id"`  // a developer-defined identifier for the component, max 100 characters
	Title      string      `json:"title"`      // the title of the popup modal
	Components []Component `json:"components"` // between 1 and 5 (inclusive) components that make up the modal
}

/* HELPER FUNCTIONS */

func (i *Interaction) BuildResponse(embeds []*Embed) *InteractionResponseMessages {
	ir := &InteractionResponseMessages{
		Data: &InteractionCallbackDataMessages{},
	}

	for _, embed := range embeds {
		ir.Data.Embeds = append(ir.Data.Embeds, *embed)
	}

	if i.Type == InteractionTypeApplicationCommand {
		ir.Type = ChannelMessageWithSource

	} else {
		ir.Type = UpdateMessage
	}

	return ir
}
