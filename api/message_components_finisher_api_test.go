package api

import (
	"reflect"
	"testing"
)

func TestComponent_GetType(t *testing.T) {
	type fields struct {
		Type ComponentType
	}
	tests := []struct {
		name   string
		fields fields
		want   ComponentType
	}{
		{
			name:   "Action Row",
			fields: fields{Type: ComponentTypeActionRow},
			want:   ComponentTypeActionRow,
		},
		{
			name:   "Button",
			fields: fields{Type: ComponentTypeButton},
			want:   ComponentTypeButton,
		},
		{
			name:   "Select Menu",
			fields: fields{Type: ComponentTypeSelectMenu},
			want:   ComponentTypeSelectMenu,
		},
		{
			name:   "Text Input",
			fields: fields{Type: ComponentTypeTextInput},
			want:   ComponentTypeTextInput,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Component{
				Type: tt.fields.Type,
			}
			if got := c.GetType(); got != tt.want {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponent_SetType(t *testing.T) {
	type fields struct {
		Type ComponentType
	}
	type args struct {
		t ComponentType
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Component
	}{
		{
			name:   "Action Row",
			fields: fields{Type: ComponentTypeActionRow},
			args:   args{t: ComponentTypeActionRow},
			want: &Component{
				Type: ComponentTypeActionRow,
			},
		},
		{
			name:   "Button",
			fields: fields{Type: ComponentTypeButton},
			args:   args{t: ComponentTypeButton},
			want: &Component{
				Type: ComponentTypeButton,
			},
		},
		{
			name:   "Select Menu",
			fields: fields{Type: ComponentTypeSelectMenu},
			args:   args{t: ComponentTypeSelectMenu},
			want: &Component{
				Type: ComponentTypeSelectMenu,
			},
		},
		{
			name:   "Text Input",
			fields: fields{Type: ComponentTypeTextInput},
			args:   args{t: ComponentTypeTextInput},
			want: &Component{
				Type: ComponentTypeTextInput,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Component{
				Type: tt.fields.Type,
			}
			if got := c.SetType(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponent_GetCustomID(t *testing.T) {
	type fields struct {
		CustomID string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "Empty String Test",
			fields: fields{CustomID: ""},
			want:   "",
		},
		{
			name:   "Single Character Test",
			fields: fields{CustomID: "G"},
			want:   "G",
		},
		{
			name:   "Long String Test",
			fields: fields{CustomID: "The quick brown fox jumps over the lazy dog"},
			want:   "The quick brown fox jumps over the lazy dog",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Component{
				CustomID: tt.fields.CustomID,
			}
			if got := c.GetCustomID(); got != tt.want {
				t.Errorf("GetCustomID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponent_SetCustomID(t *testing.T) {
	type fields struct {
		CustomID string
	}
	type args struct {
		t string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Component
	}{
		{
			name:   "Empty String Test",
			fields: fields{CustomID: ""},
			args:   args{t: ""},
			want: &Component{
				CustomID: "",
			},
		},
		{
			name:   "Single Character Test",
			fields: fields{CustomID: "G"},
			args:   args{t: "G"},
			want: &Component{
				CustomID: "G",
			},
		},
		{
			name:   "Long String Test",
			fields: fields{CustomID: "The quick brown fox jumps over the lazy dog"},
			args:   args{t: "The quick brown fox jumps over the lazy dog"},
			want: &Component{
				CustomID: "The quick brown fox jumps over the lazy dog",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Component{
				CustomID: tt.fields.CustomID,
			}
			if got := c.SetCustomID(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetCustomID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponent_IsDisabled(t *testing.T) {
	type fields struct {
		Disabled bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "Is Disabled",
			fields: fields{Disabled: true},
			want:   true,
		},
		{
			name:   "Is Not Disabled",
			fields: fields{Disabled: false},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Component{
				Disabled: tt.fields.Disabled,
			}
			if got := c.IsDisabled(); got != tt.want {
				t.Errorf("IsDisabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponent_SetDisabled(t *testing.T) {
	type fields struct {
		Disabled bool
	}
	type args struct {
		d bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Component
	}{
		{
			name:   "Set Disabled",
			fields: fields{Disabled: true},
			args:   args{d: true},
			want: &Component{
				Disabled: true,
			},
		},
		{
			name:   "Set Enabled",
			fields: fields{Disabled: false},
			args:   args{d: false},
			want: &Component{
				Disabled: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Component{
				Disabled: tt.fields.Disabled,
			}
			if got := c.SetDisabled(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDisabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponent_GetButtonStyle(t *testing.T) {
	type fields struct {
		Style interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   ButtonStyle
	}{
		{
			name:   "Primary",
			fields: fields{Style: ButtonPrimary},
			want:   ButtonPrimary,
		},
		{
			name:   "Secondary",
			fields: fields{Style: ButtonSecondary},
			want:   ButtonSecondary,
		},
		{
			name:   "Success",
			fields: fields{Style: ButtonSuccess},
			want:   ButtonSuccess,
		},
		{
			name:   "Danger",
			fields: fields{Style: ButtonDanger},
			want:   ButtonDanger,
		},
		{
			name:   "Link",
			fields: fields{Style: ButtonLink},
			want:   ButtonLink,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Component{
				Style: tt.fields.Style,
			}
			if got := c.GetButtonStyle(); got != tt.want {
				t.Errorf("GetButtonStyle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponent_SetButtonStyle(t *testing.T) {
	type fields struct {
		Style interface{}
	}
	type args struct {
		s ButtonStyle
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Component
	}{
		{
			name:   "Primary",
			fields: fields{Style: ButtonPrimary},
			args:   args{s: ButtonPrimary},
			want: &Component{
				Style: ButtonPrimary,
			},
		},
		{
			name:   "Secondary",
			fields: fields{Style: ButtonSecondary},
			args:   args{s: ButtonSecondary},
			want: &Component{
				Style: ButtonSecondary,
			},
		},
		{
			name:   "Success",
			fields: fields{Style: ButtonSuccess},
			args:   args{s: ButtonSuccess},
			want: &Component{
				Style: ButtonSuccess,
			},
		},
		{
			name:   "Danger",
			fields: fields{Style: ButtonDanger},
			args:   args{s: ButtonDanger},
			want: &Component{
				Style: ButtonDanger,
			},
		},
		{
			name:   "Link",
			fields: fields{Style: ButtonLink},
			args:   args{s: ButtonLink},
			want: &Component{
				Style: ButtonLink,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Component{
				Style: tt.fields.Style,
			}
			if got := c.SetButtonStyle(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetButtonStyle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponent_GetTextInputStyle(t *testing.T) {
	type fields struct {
		Style interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   TextInputStyle
	}{
		{
			name:   "Short",
			fields: fields{Style: TextInputShort},
			want:   TextInputShort,
		},
		{
			name:   "Paragraph",
			fields: fields{Style: TextInputParagraph},
			want:   TextInputParagraph,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Component{
				Style: tt.fields.Style,
			}
			if got := c.GetTextInputStyle(); got != tt.want {
				t.Errorf("GetTextInputStyle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponent_SetTextInputStyle(t *testing.T) {
	type fields struct {
		Style interface{}
	}
	type args struct {
		s TextInputStyle
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Component
	}{
		{
			name:   "Short",
			fields: fields{Style: TextInputShort},
			args:   args{s: TextInputShort},
			want: &Component{
				Style: TextInputShort,
			},
		},
		{
			name:   "Paragraph",
			fields: fields{Style: TextInputParagraph},
			args:   args{s: TextInputParagraph},
			want: &Component{
				Style: TextInputParagraph,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Component{
				Style: tt.fields.Style,
			}
			if got := c.SetTextInputStyle(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetButtonStyle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewComponent(t *testing.T) {
	tests := []struct {
		name string
		want *Component
	}{
		{
			name: "New Component",
			want: &Component{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewComponent(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewComponent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponent_GetEmoji(t *testing.T) {
	type fields struct {
		Emoji *Emoji
	}
	tests := []struct {
		name   string
		fields fields
		want   *Emoji
	}{
		{
			name: "Custom Emoji",
			fields: fields{Emoji: &Emoji{
				ID:       StringToSnowflake("941127649168871454"),
				Name:     "glitch",
				Animated: false,
			}},
			want: &Emoji{
				ID:       StringToSnowflake("941127649168871454"),
				Name:     "glitch",
				Animated: false,
			},
		},
		{
			name: "Unicode Emoji",
			fields: fields{Emoji: &Emoji{
				ID:       nil,
				Name:     "🔥",
				Animated: false,
			}},
			want: &Emoji{
				ID:       nil,
				Name:     "🔥",
				Animated: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Component{
				Emoji: tt.fields.Emoji,
			}
			if got := c.GetEmoji(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEmoji() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponent_SetEmoji(t *testing.T) {
	type fields struct {
		Emoji *Emoji
	}
	type args struct {
		e *Emoji
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Component
	}{
		{
			name: "Custom Emoji",
			fields: fields{Emoji: &Emoji{
				ID:       StringToSnowflake("941127649168871454"),
				Name:     "glitch",
				Animated: false,
			}},
			args: args{e: &Emoji{
				ID:       StringToSnowflake("941127649168871454"),
				Name:     "glitch",
				Animated: false,
			}},
			want: &Component{
				Emoji: &Emoji{
					ID:       StringToSnowflake("941127649168871454"),
					Name:     "glitch",
					Animated: false,
				},
			},
		},
		{
			name: "Unicode Emoji",
			fields: fields{Emoji: &Emoji{
				ID:       nil,
				Name:     "🔥",
				Animated: false,
			}},
			args: args{e: &Emoji{
				ID:       nil,
				Name:     "🔥",
				Animated: false,
			}},
			want: &Component{
				Emoji: &Emoji{
					ID:       nil,
					Name:     "🔥",
					Animated: false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Component{
				Emoji: tt.fields.Emoji,
			}
			if got := c.SetEmoji(tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetEmoji() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponent_GetURL(t *testing.T) {
	type fields struct {
		URL string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "URL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Component{
				URL: tt.fields.URL,
			}
			if got := c.GetURL(); got != tt.want {
				t.Errorf("GetURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComponent_SetURL(t *testing.T) {
	type fields struct {
		Type        ComponentType
		CustomID    string
		Disabled    bool
		Style       interface{}
		Label       string
		Emoji       *Emoji
		URL         string
		Options     []SelectOption
		MinValues   int
		MaxValues   int
		Placeholder string
		Components  []Component
		MinLength   int
		MaxLength   int
		Required    bool
		Value       string
	}
	type args struct {
		u string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Component
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Component{
				Type:        tt.fields.Type,
				CustomID:    tt.fields.CustomID,
				Disabled:    tt.fields.Disabled,
				Style:       tt.fields.Style,
				Label:       tt.fields.Label,
				Emoji:       tt.fields.Emoji,
				URL:         tt.fields.URL,
				Options:     tt.fields.Options,
				MinValues:   tt.fields.MinValues,
				MaxValues:   tt.fields.MaxValues,
				Placeholder: tt.fields.Placeholder,
				Components:  tt.fields.Components,
				MinLength:   tt.fields.MinLength,
				MaxLength:   tt.fields.MaxLength,
				Required:    tt.fields.Required,
				Value:       tt.fields.Value,
			}
			if got := c.SetURL(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
