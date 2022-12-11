/*
 * Copyright (c) 2022. Veteran Software
 *
 * Discord API Wrapper - A custom wrapper for the Discord REST API developed for a proprietary project.
 *
 * This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public
 * License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later
 * version.
 *
 * This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied
 * warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License along with this program.
 * If not, see <http://www.gnu.org/licenses/>.
 */

package api

// NewComponent - Build a new Component
func NewComponent() *Component {
	return &Component{}
}

// SetType - sets the Component type
func (c *Component) SetType(t ComponentType) *Component {
	c.Type = t

	return c
}

// SetCustomID - sets the CustomID of the component
func (c *Component) SetCustomID(t string) *Component {
	c.CustomID = t

	return c
}

// IsDisabled - is the Component disabled
func (c *Component) IsDisabled() bool {
	return c.Disabled
}

// SetDisabled - sets the Disabled state of the Component
func (c *Component) SetDisabled(d bool) *Component {
	c.Disabled = d

	return c
}

// SetButtonStyle - sets the style of the Button
func (c *Component) SetButtonStyle(s ButtonStyle) *Component {
	c.Style = s

	return c
}

// SetTextInputStyle - sets the style of the TextInput
func (c *Component) SetTextInputStyle(s TextInputStyle) *Component {
	c.Style = s

	return c
}

// SetEmoji - sets the Emoji for the Button
func (c *Component) SetEmoji(e *Emoji) *Component {
	c.Emoji = e

	return c
}

// SetURL - sets the URL for the Button
func (c *Component) SetURL(u string) *Component {
	c.URL = u

	return c
}

// NewModalResponse - Build a new response containing a modal
//
//goland:noinspection GoUnusedExportedFunction
func NewModalResponse() *InteractionResponseModal {
	return &InteractionResponseModal{
		CallbackType: Modal,
		Data:         &InteractionCallbackDataModal{},
	}
}

// SetCustomID - sets the CustomID of the InteractionResponseModal
func (i *InteractionResponseModal) SetCustomID(c string) *InteractionResponseModal {
	i.Data.CustomID = c

	return i
}

// SetTitle - sets the Title of the InteractionResponseModal
func (i *InteractionResponseModal) SetTitle(t string) *InteractionResponseModal {
	i.Data.Title = t

	return i
}

// AddComponent - adds a single Component to the InteractionResponseModal
func (i *InteractionResponseModal) AddComponent(c *Component) *InteractionResponseModal {
	i.Data.Components = append(i.Data.Components, *c)

	return i
}

// NewMessageResponse - Build a new response containing a message
//
//goland:noinspection GoUnusedExportedFunction
func NewMessageResponse() *InteractionResponseMessages {
	return &InteractionResponseMessages{
		Data: &InteractionCallbackDataMessages{},
	}
}

// SetType - sets the Type of the InteractionResponseMessages object
func (i *InteractionResponseMessages) SetType(t *InteractionCallbackType) *InteractionResponseMessages {
	i.Type = *t

	return i
}

// SetTts - sets the TTS flag
func (i *InteractionResponseMessages) SetTts(tts bool) *InteractionResponseMessages {
	i.Data.TTS = tts

	return i
}

// SetContent - sets the content of the response message
func (i *InteractionResponseMessages) SetContent(content string) *InteractionResponseMessages {
	i.Data.Content = content

	return i
}

// AddEmbed - adds a single Embed to the response
func (i *InteractionResponseMessages) AddEmbed(e *Embed) *InteractionResponseMessages {
	i.Data.Embeds = append(i.Data.Embeds, *e)

	return i
}

// AddEmbeds - add multiple Embed objects (max 10) to the response
func (i *InteractionResponseMessages) AddEmbeds(e *[]Embed) *InteractionResponseMessages {
	i.Data.Embeds = append(i.Data.Embeds, *e...)

	return i
}

// SetEphemeral - sets the Ephemeral flag to the response message
func (i *InteractionResponseMessages) SetEphemeral() *InteractionResponseMessages {
	i.Data.Flags = i.Data.Flags | Ephemeral

	return i
}

// AddFlag - bit shifts a new flag into the flags of the response message
func (i *InteractionResponseMessages) AddFlag(f MessageFlags) *InteractionResponseMessages {
	if f == SuppressEmbeds || f == Ephemeral {
		i.Data.Flags = i.Data.Flags | f

		return i
	}

	return i
}

// AddComponent - adds a single Component to the response message
func (i *InteractionResponseMessages) AddComponent(c *Component) *InteractionResponseMessages {
	i.Data.Components = append(i.Data.Components, *c)

	return i
}

// AddAttachment - adds a single attachment to the response message
func (i *InteractionResponseMessages) AddAttachment(a *Attachment) *InteractionResponseMessages {
	i.Data.Attachments = append(i.Data.Attachments, *a)

	return i
}

// NewAutocompleteResponse - Build a new response containing a modal
//
//goland:noinspection GoUnusedExportedFunction
func NewAutocompleteResponse() *InteractionResponseAutocomplete {
	return &InteractionResponseAutocomplete{
		Type: AutocompleteResult,
		Data: &InteractionCallbackDataAutocomplete{},
	}
}

// AddChoice - adds a single choice to the autocomplete response
func (i *InteractionResponseAutocomplete) AddChoice(c *ApplicationCommandOptionChoice) *InteractionResponseAutocomplete {
	i.Data.Choices = append(i.Data.Choices, c)

	return i
}

// AddChoices - adds multiple choices to the autocomplete response
func (i *InteractionResponseAutocomplete) AddChoices(c []*ApplicationCommandOptionChoice) *InteractionResponseAutocomplete {
	i.Data.Choices = append(i.Data.Choices, c...)

	return i
}
