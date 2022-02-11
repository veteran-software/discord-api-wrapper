package api

/* MESSAGE COMPONENT API */

func NewComponent() *Component {
	return &Component{}
}

func (c *Component) GetType() ComponentType {
	return c.Type
}

func (c *Component) SetType(t ComponentType) *Component {
	c.Type = t

	return c
}

func (c *Component) GetCustomID() string {
	return c.CustomID
}

func (c *Component) SetCustomID(t string) *Component {
	c.CustomID = t

	return c
}

func (c *Component) IsDisabled() bool {
	return c.Disabled
}

func (c *Component) SetDisabled(d bool) *Component {
	c.Disabled = d

	return c
}

func (c *Component) GetButtonStyle() ButtonStyle {
	return c.Style.(ButtonStyle)
}

func (c *Component) SetButtonStyle(s ButtonStyle) *Component {
	c.Style = s

	return c
}

func (c *Component) GetTextInputStyle() TextInputStyle {
	return c.Style.(TextInputStyle)
}

func (c *Component) SetTextInputStyle(s TextInputStyle) *Component {
	c.Style = s

	return c
}

func (c *Component) GetEmoji() *Emoji {
	return c.Emoji
}

func (c *Component) SetEmoji(e *Emoji) *Component {
	c.Emoji = e

	return c
}

func (c *Component) GetURL() string {
	return c.URL
}

func (c *Component) SetURL(u string) *Component {
	c.URL = u

	return c
}

/*
NewModalResponse

Build a new response containing a modal
*/
func NewModalResponse() *InteractionResponseModal {
	return &InteractionResponseModal{
		CallbackType: Modal,
		Data:         &InteractionCallbackDataModal{},
	}
}

func (i *InteractionResponseModal) SetCustomID(c string) *InteractionResponseModal {
	i.Data.CustomID = c

	return i
}

func (i *InteractionResponseModal) SetTitle(t string) *InteractionResponseModal {
	i.Data.Title = t

	return i
}

func (i *InteractionResponseModal) AddComponent(c *Component) *InteractionResponseModal {
	i.Data.Components = append(i.Data.Components, *c)

	return i
}

/*
NewMessageResponse

Build a new response containing a message
*/
func NewMessageResponse() *InteractionResponseMessages {
	return &InteractionResponseMessages{
		Data: &InteractionCallbackDataMessages{},
	}
}

func (i *InteractionResponseMessages) SetType(t InteractionCallbackType) *InteractionResponseMessages {
	i.Type = t

	return i
}

func (i *InteractionResponseMessages) SetTts(tts bool) *InteractionResponseMessages {
	i.Data.TTS = tts

	return i
}

func (i *InteractionResponseMessages) SetContent(content string) *InteractionResponseMessages {
	i.Data.Content = content

	return i
}

func (i *InteractionResponseMessages) AddEmbed(e *Embed) *InteractionResponseMessages {
	i.Data.Embeds = append(i.Data.Embeds, *e)

	return i
}

func (i *InteractionResponseMessages) SetEphemeral() *InteractionResponseMessages {
	i.Data.Flags = i.Data.Flags | Ephemeral

	return i
}

func (i *InteractionResponseMessages) AddFlag(f MessageFlags) *InteractionResponseMessages {
	if f == SuppressEmbeds || f == Ephemeral {
		i.Data.Flags = i.Data.Flags | f

		return i
	}

	return i
}

func (i *InteractionResponseMessages) AddComponent(c *Component) *InteractionResponseMessages {
	i.Data.Components = append(i.Data.Components, *c)

	return i
}

func (i *InteractionResponseMessages) AddAttachment(a *Attachment) *InteractionResponseMessages {
	i.Data.Attachments = append(i.Data.Attachments, *a)

	return i
}

func (i *InteractionResponseMessages) GetEmbeds() []Embed {
	return i.Data.Embeds
}

func (i *InteractionResponseMessages) GetType() InteractionCallbackType {
	return i.Type
}

/*
NewAutocompleteResponse

Build a new response containing a modal
*/
func NewAutocompleteResponse() *InteractionResponseAutocomplete {
	return &InteractionResponseAutocomplete{
		Type: AutocompleteResult,
		Data: &InteractionCallbackDataAutocomplete{},
	}
}

func (i *InteractionResponseAutocomplete) AddChoice(c *ApplicationCommandOptionChoice) *InteractionResponseAutocomplete {
	i.Data.Choices = append(i.Data.Choices, c)

	return i
}
