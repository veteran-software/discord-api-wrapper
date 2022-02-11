package api

/* COMPONENT OBJECT */

/*
Component

https://discord.com/developers/docs/interactions/message-components#component-object-component-structure

Type: ComponentType; valid for all types

ButtonStyle: one of ButtonStyle

Label: text that appears on the button, max 80 characters

Emoji: name, id, and animated

CustomID: a developer-defined identifier for the button, max 100 characters

URL: a URL for link-style buttons

Disabled: whether the button is disabled, default false
*/
type Component struct {
	Type ComponentType `json:"type"`

	// Buttons and Select Menus
	CustomID string `json:"custom_id,omitempty"`
	Disabled bool   `json:"disabled,omitempty"`

	// Buttons only
	Style interface{} `json:"style,omitempty"`
	Label string      `json:"label,omitempty"` // And Text inputs
	Emoji *Emoji      `json:"emoji,omitempty"`
	URL   string      `json:"url,omitempty"`

	// Select Menus only
	Options   []SelectOption `json:"options,omitempty"`
	MinValues int            `json:"min_values,omitempty"`
	MaxValues int            `json:"max_values,omitempty"`

	/* Select Menus, Text Inputs */
	Placeholder string `json:"placeholder,omitempty"`

	// Action Rows
	Components []Component `json:"components,omitempty"`

	/* Text Inputs */
	MinLength int    `json:"min_length,omitempty"`
	MaxLength int    `json:"max_length"`
	Required  bool   `json:"required,omitempty"`
	Value     string `json:"value,omitempty"`
}

/*
ComponentType

https://discord.com/developers/docs/interactions/message-components#component-object-component-types

ComponentTypeActionRow: A container for other components

ComponentTypeButton: A clickable button
*/
type ComponentType int

const (
	ComponentTypeActionRow  ComponentType = iota + 1 // A container for other components
	ComponentTypeButton                              // A button object
	ComponentTypeSelectMenu                          // A select menu for picking from choices
	ComponentTypeTextInput                           // A text input object
)

/* BUTTON OBJECT */

/*
Button

Buttons are interactive components that render on messages. They can be clicked by users, and send an interaction to your app when clicked.

    * Buttons must be sent inside an ComponentTypeActionRow
    * An ComponentTypeActionRow can contain up to 5 buttons

--------

Type: ComponentTypeButton for a button

Style: one of ButtonStyle

Label: text that appears on the button, max 80 characters

Emoji: name, id, and animated

CustomID: a developer-defined identifier for the button, max 100 characters

URL: a URL for link-style buttons

Disabled: whether the button is disabled, default false
*/
type Button struct {
	Type     ComponentType `json:"type"`
	Style    ButtonStyle   `json:"style"`
	Label    string        `json:"label,omitempty"`
	Emoji    Emoji         `json:"emoji,omitempty"`
	CustomID string        `json:"custom_id,omitempty"`
	URL      string        `json:"url,omitempty"`
	Disabled bool          `json:"disabled,omitempty"`
}

/*
ButtonStyle

Buttons come in a variety of styles to convey different types of actions. These styles also define what fields are valid for a button.

    Non-link buttons must have a custom_id, and cannot have a URL
    Link buttons must have a URL, and cannot have a custom_id
    Link buttons do not send an interaction to your app when clicked

--------

ButtonPrimary: color: blurple; requires field: custom_id

ButtonSecondary: color: grey; requires field: custom_id

ButtonSuccess: color: green; requires field: custom_id

ButtonDanger: color: red; requires field: custom_id

ButtonLink: color: grey; requires field: url
*/
type ButtonStyle int

const (
	ButtonPrimary   ButtonStyle = iota + 1 // Color: blurple
	ButtonSecondary                        // Color: grey
	ButtonSuccess                          // Color: green
	ButtonDanger                           // Color: red
	ButtonLink                             // Color: grey, navigates to a URL
)

/* SELECT MENU OBJECT */

/*
SelectMenu

https://discord.com/developers/docs/interactions/message-components#select-menu-object-select-menu-structure

CustomID: a developer-defined identifier for the button, max 100 characters

Options: the choices in the select, max 25

Placeholder: custom placeholder text if nothing is selected, max 100 characters

MinValues: the minimum number of items that must be chosen; default 1, min 0, max 25

MaxValues: the maximum number of items that can be chosen; default 1, max 25
*/
type SelectMenu struct {
	Type        ComponentType  `json:"type"`
	CustomID    string         `json:"custom_id"`
	Options     []SelectOption `json:"options"`
	Placeholder string         `json:"placeholder,omitempty"`
	MinValues   int64          `json:"min_values,omitempty"`
	MaxValues   int64          `json:"max_values,omitempty"`
	Disabled    bool           `json:"disabled,omitempty"`
}

/*
SelectOption

https://discord.com/developers/docs/interactions/message-components#select-menu-object-select-option-structure

Label: the user-facing name of the option, max 25 characters

Value: 	the dev-define value of the option, max 100 characters

Description: an additional description of the option, max 50 characters

Emoji: id, name, and animated

Default: will render this option as selected by default
*/
type SelectOption struct {
	Label       string `json:"label"`
	Value       string `json:"value"`
	Description string `json:"description,omitempty"`
	Emoji       *Emoji `json:"emoji,omitempty"`
	Default     bool   `json:"default,omitempty"`
}

type TextInputObject struct {
	Type        ComponentType  `json:"type"`                  // 4 for a text input
	CustomID    string         `json:"custom_id"`             // a developer-defined identifier for the input, max 100 characters
	Style       TextInputStyle `json:"style"`                 // the Text Input Style
	Label       string         `json:"label"`                 // the label for this component
	MinLength   uint           `json:"min_length,omitempty"`  // the minimum input length for a text input, min 0, max 4000
	MaxLength   uint           `json:"max_length,omitempty"`  // the maximum input length for a text input, min 1, max 4000
	Required    bool           `json:"required,omitempty"`    // whether this component is required to be filled, default false
	Value       string         `json:"value,omitempty"`       // a pre-filled value for this component, max 4000 characters
	Placeholder string         `json:"placeholder,omitempty"` // custom placeholder text if the input is empty, max 100 characters
}

type TextInputStyle int

const (
	TextInputShort TextInputStyle = iota + 1
	TextInputParagraph
)
