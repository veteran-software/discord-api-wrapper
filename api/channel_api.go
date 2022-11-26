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

import (
	"net/url"
	"time"
)

// IsValidLength - Checks that the total size of an Embed is valid for sending
func (e *Embed) IsValidLength() bool {
	if len(e.Title) <= titleLimit && len(e.Description) <= descriptionLimit && len(e.Fields) <= fieldCount && len(e.Footer.Text) <= footerTextLimit && len(e.Author.Name) <= authorNameLimit {
		for _, field := range e.Fields {
			if len(field.Name) > fieldNameLimit || len(field.Value) > fieldValueLimit {
				return false
			}
		}

		return true
	}

	return false
}

// String - Converts a Channel into a string for easy output
func (c *Channel) String() string {
	var chanType string

	switch c.Type {
	case GuildText:
		chanType = "GTC:"
	case DM:
		chanType = "DM:"
	case GroupDM:
		chanType = "GDM:"
	case GuildNews:
		chanType = "GNC:"
	case GuildNewsThread:
		chanType = "GNT:"
	case GuildPublicThread:
		chanType = "GPuT:"
	case GuildPrivateThread:
		chanType = "GPrT:"
	}

	return chanType + c.Name + "(" + c.ID.String() + ")"
}

// NewEmbed - Instantiates a new Embed object with the color defaulted to red and the timestamp defaulted to time.Now()
//
//goland:noinspection GoUnusedExportedFunction
func NewEmbed() *Embed {
	return &Embed{
		Title:       "",
		Type:        RichEmbed,
		Description: "",
		URL:         "",
		Timestamp:   time.Now().Format(time.RFC3339),
		Color:       16711680,
		Footer:      nil,
		Image:       nil,
		Thumbnail:   nil,
		Author:      nil,
		Fields:      nil,
	}
}

// SetTitle - Set the Embed title
func (e *Embed) SetTitle(title string) *Embed {
	if len(title) > titleLimit {
		title = title[:titleLimit-4] + " ..."
	}
	e.Title = title

	return e
}

// SetDescription - Set the Embed description
func (e *Embed) SetDescription(description string) *Embed {
	if len(description) > descriptionLimit {
		description = description[:descriptionLimit-4] + " ..."
	}
	e.Description = description

	return e
}

// SetURL - Set the Embed URL
func (e *Embed) SetURL(u string) *Embed {
	// We only check for an error to validate if it's a properly formed URL
	if _, err := url.Parse(u); err != nil {
		return e
	}
	e.URL = u

	return e
}

// SetTimestamp - Set the Embed timestamp
func (e *Embed) SetTimestamp(ts time.Time) *Embed {
	e.Timestamp = ts.UTC().Format(time.RFC3339)

	return e
}

// SetColor - Set the Embed color
func (e *Embed) SetColor(c int64) *Embed {
	e.Color = c

	return e
}

// SetFooter - Set the Footer
func (e *Embed) SetFooter(text string, iconURL string) *Embed {
	e.Footer = newFooter().SetText(text).SetIconURL(iconURL)

	return e
}

// SetImage - Set the Image
func (e *Embed) SetImage(imageURL string) *Embed {
	e.Image = newImage().SetURL(imageURL)

	return e
}

// SetThumbnail - Set the Thumbnail
func (e *Embed) SetThumbnail(thumbnailURL string) *Embed {
	e.Thumbnail = newThumbnail().SetURL(thumbnailURL)

	return e
}

// SetAuthor - Set the Author
func (e *Embed) SetAuthor(name, url string, iconURL *string) *Embed {
	e.Author = newAuthor().SetName(name).SetURL(url).SetIconURL(iconURL)

	return e
}

// AddField - Add a singular Field
func (e *Embed) AddField(name, value string, inline bool) *Embed {
	e.Fields = append(e.Fields, NewField().SetName(name).SetValue(value).SetInline(inline))

	return e
}

// AddFields - Add multiple fields; must create the Field objects first
func (e *Embed) AddFields(fields ...*Field) *Embed {
	if len(fields) == 0 {
		return e
	}

	e.Fields = append(e.Fields, fields...)

	return e
}

/* EMBED FOOTER */

func newFooter() *Footer {
	return &Footer{}
}

// SetText - set the Footer text
func (f *Footer) SetText(text string) *Footer {
	if len(text) > footerTextLimit {
		text = text[:footerTextLimit-4] + " ..."
	}
	f.Text = text

	return f
}

// SetIconURL - set the Footer IconURL
func (f *Footer) SetIconURL(iconURL string) *Footer {
	if _, err := url.Parse(iconURL); err != nil {
		return f
	}
	f.IconURL = iconURL

	return f
}

/* EMBED IMAGE */

func newImage() *Image {
	return &Image{}
}

// SetURL - set the Image URL
func (i *Image) SetURL(u string) *Image {
	if _, err := url.Parse(u); err != nil {
		return i
	}
	i.URL = u

	return i
}

/* EMBED THUMBNAIL */

func newThumbnail() *Thumbnail {
	return &Thumbnail{}
}

// SetURL - set the Thumbnail URL
func (t *Thumbnail) SetURL(u string) *Thumbnail {
	if _, err := url.Parse(u); err != nil {
		return t
	}
	t.URL = u

	return t
}

/* EMBED AUTHOR */

func newAuthor() *Author {
	return &Author{}
}

// SetName - set the Author Name
func (a *Author) SetName(name string) *Author {
	if len(name) > authorNameLimit {
		name = name[:authorNameLimit-4] + " ..."
	}
	a.Name = name

	return a
}

// SetURL - set the Author URL
func (a *Author) SetURL(u string) *Author {
	if _, err := url.Parse(u); err != nil {
		return a
	}
	a.URL = u

	return a
}

// SetIconURL - set the Author IconURL
func (a *Author) SetIconURL(u *string) *Author {
	if u == nil {
		return a
	}

	if _, err := url.Parse(*u); err != nil {
		return a
	}
	a.IconURL = u

	return a
}

// NewField - Create a new base Field to chain against
func NewField() *Field {
	return &Field{}
}

// SetName - set the Field Name
func (f *Field) SetName(name string) *Field {
	if len(name) > fieldNameLimit {
		name = name[:fieldNameLimit-4] + " ..."
	}
	f.Name = name

	return f
}

// SetValue - set the Field Value
func (f *Field) SetValue(value string) *Field {
	if len(value) > fieldValueLimit {
		value = value[:fieldValueLimit-4] + " ..."
	}
	f.Value = value

	return f
}

// SetInline - set the Field Inline
func (f *Field) SetInline(inline bool) *Field {
	f.Inline = inline

	return f
}

// IsInline - Helper function for testing is inline
func (f *Field) IsInline() bool {
	return f.Inline
}
