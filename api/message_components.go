/*
 * Copyright (c) 2022. Veteran Software
 *
 * Discord API Wrapper - A custom wrapper for the Discord REST API developed for a proprietary project.
 *
 * This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public
 * License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied
 * warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License along with this program.
 * If not, see <http://www.gnu.org/licenses/>.
 */

package api

// Component - Components are a new field on the message object, so you can use them whether you're sending messages or responding to a slash command or other interaction.
//
// The top-level component's field is an array of Action Row components.
type Component struct {
	Type        ComponentType  `json:"type"`                  // ComponentType; valid for all types
	CustomID    string         `json:"custom_id,omitempty"`   // a developer-defined identifier for the button, max 100 characters
	Disabled    bool           `json:"disabled,omitempty"`    // whether the button is disabled, default false
	Style       interface{}    `json:"style,omitempty"`       // one of ButtonStyle
	Label       string         `json:"label,omitempty"`       // text that appears on the button, max 80 characters
	Emoji       *Emoji         `json:"emoji,omitempty"`       // name, id, and animated
	URL         string         `json:"url,omitempty"`         // a URL for link-style buttons
	Options     []SelectOption `json:"options,omitempty"`     // the choices in the select, max 25
	MinValues   int            `json:"min_values,omitempty"`  // the minimum number of items that must be chosen; default 1, min 0, max 25
	MaxValues   int            `json:"max_values,omitempty"`  // the maximum number of items that can be chosen; default 1, max 25
	Placeholder string         `json:"placeholder,omitempty"` // custom placeholder text if nothing is selected, max 100 characters
	Components  []Component    `json:"components,omitempty"`  // a list of child components
	MinLength   int            `json:"min_length,omitempty"`  // the minimum input length for a text input
	MaxLength   int            `json:"max_length,omitempty"`  // the maximum input length for a text input
	Required    bool           `json:"required,omitempty"`    // whether this component is required to be filled
	Value       string         `json:"value,omitempty"`       // a pre-filled value for this component
}

// ComponentType - The type of component
type ComponentType int

const (
	ComponentTypeActionRow  ComponentType = iota + 1 // A container for other components
	ComponentTypeButton                              // A clickable button
	ComponentTypeSelectMenu                          // A select menu for picking from choices
	ComponentTypeTextInput                           // A text input object
)

// Button - Buttons are interactive components that render on messages.
//
// They can be clicked by users, and send an interaction to your app when clicked.
//
//   - Buttons must be sent inside an ComponentTypeActionRow
//   - An ComponentTypeActionRow can contain up to 5 buttons
type Button struct {
	Type     ComponentType `json:"type"`                // ComponentType for a button
	Style    ButtonStyle   `json:"style"`               // one of ButtonStyle
	Label    string        `json:"label,omitempty"`     // text that appears on the button, max 80 characters
	Emoji    Emoji         `json:"emoji,omitempty"`     // name, id, and animated
	CustomID string        `json:"custom_id,omitempty"` // a developer-defined identifier for the button, max 100 characters
	URL      string        `json:"url,omitempty"`       // a URL for link-style buttons
	Disabled bool          `json:"disabled,omitempty"`  // whether the button is disabled, default false
}

// ButtonStyle - Buttons come in a variety of styles to convey different types of actions.
//
// These styles also define what fields are valid for a button.
//
//	Non-link buttons must have a custom_id, and cannot have a URL
//	Link buttons must have a URL, and cannot have a custom_id
//	Link buttons do not send an interaction to your app when clicked
type ButtonStyle int

//goland:noinspection SpellCheckingInspection
const (
	ButtonPrimary   ButtonStyle = iota + 1 // color: blurple; requires field: custom_id
	ButtonSecondary                        // color: grey; requires field: custom_id
	ButtonSuccess                          // color: green; requires field: custom_id
	ButtonDanger                           // color: red; requires field: custom_id
	ButtonLink                             // color: grey; requires field: url
)

// SelectMenu - Select menus support single-select and multi-select behavior, meaning you can prompt a user to choose just one item from a list, or multiple.
//
// When a user finishes making their choice by clicking out of the dropdown or closing the half-sheet, your app will receive an interaction.
//
//	Select menus must be sent inside an Action Row
//	An Action Row can contain only one select menu
//	An Action Row containing a select menu cannot also contain buttons
type SelectMenu struct {
	Type        ComponentType  `json:"type"`                  // ComponentTypeSelectMenu for a select menu
	CustomID    string         `json:"custom_id"`             // a developer-defined identifier for the button, max 100 characters
	Options     []SelectOption `json:"options"`               // the choices in the select, max 25
	Placeholder string         `json:"placeholder,omitempty"` // custom placeholder text if nothing is selected, max 150 characters
	MinValues   int64          `json:"min_values,omitempty"`  // the minimum number of items that must be chosen; default 1, min 0, max 25
	MaxValues   int64          `json:"max_values,omitempty"`  // the maximum number of items that can be chosen; default 1, max 25
	Disabled    bool           `json:"disabled,omitempty"`    // disable the select, default false
}

// SelectOption - Represents a single select menu option
type SelectOption struct {
	Label       string `json:"label"`                 // the user-facing name of the option, max 25 characters
	Value       string `json:"value"`                 // the dev-define value of the option, max 100 characters
	Description string `json:"description,omitempty"` // an additional description of the option, max 50 characters
	Emoji       *Emoji `json:"emoji,omitempty"`       // id, name, and animated
	Default     bool   `json:"default,omitempty"`     // will render this option as selected by default
}

// TextInput - Text inputs are an interactive component that render on modals. They can be used to collect short-form or long-form text.
type TextInput struct {
	Type        ComponentType  `json:"type"`                  // ComponentTypeTextInput for a text input
	CustomID    string         `json:"custom_id"`             // a developer-defined identifier for the input, max 100 characters
	Style       TextInputStyle `json:"style"`                 // the TextInputStyle
	Label       string         `json:"label"`                 // the label for this component
	MinLength   uint           `json:"min_length,omitempty"`  // the minimum input length for a text input, min 0, max 4000
	MaxLength   uint           `json:"max_length,omitempty"`  // the maximum input length for a text input, min 1, max 4000
	Required    bool           `json:"required,omitempty"`    // whether this component is required to be filled, default true
	Value       string         `json:"value,omitempty"`       // a pre-filled value for this component, max 4000 characters
	Placeholder string         `json:"placeholder,omitempty"` // custom placeholder text if the input is empty, max 100 characters
}

// TextInputStyle - Denotes if a text input is short form or paragraph form
type TextInputStyle int

const (
	TextInputShort     TextInputStyle = iota + 1 // A single-line input
	TextInputParagraph                           // A multi-line input
)
