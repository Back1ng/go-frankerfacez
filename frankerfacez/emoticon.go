package frankerfacez

import (
	"time"
)

type Emoticon struct {
	// ID The integer ID of an emote. IDs will always be numbers in API v1,
	// though in a future release we may switch to using strings.
	ID int `json:"id"`

	// Name The name of an emote. This is the text string in chat that
	// should be replaced with an instance of the emote.
	Name string `json:"name"`

	// Height The height of the emote at 1x scale in pixels
	Height int `json:"height"`

	// Width The width of the emote at 1x scale in pixels
	Width int `json:"width"`

	// Public A public emote can be freely added to channels and collections on FrankerFaceZ.
	// Non-public emotes require the owner's permission in the form of an unlock code.
	Public bool `json:"public"`

	// Hidden Whether the emote should be displayed in UI like menus that list available emotes.
	// This is rarely used, and was intended to create Easter eggs.
	Hidden bool `json:"hidden"`

	// Modifier Whether the emote is a modifier emote.
	Modifier bool `json:"modifier"`

	// ModifierFlags A bitset of flags for customizing how a modifier functions. (Emote Effects)
	ModifierFlags int `json:"modifier_flags"`

	// UsageCount The number of channels and/or collections that the emote has been added to.
	UsageCount int `json:"usage_count"`

	// Urls A map of DPI scales to emote URLs. There will always be a 1 entry and there may also
	// be 2 and 4 entries for 2x and 4x DPI scales as well. Previously, URLs were sent as network-path
	// references (see RFC 3986 for details) but now all URLs are absolute. Sorry for the inconvenience this has caused.
	Urls map[string]string `json:"urls"`

	// Animated same like in Urls
	Animated map[string]string `json:"animated"`

	// CreatedAt The datetime when the emote was created
	CreatedAt time.Time `json:"created_at"`

	// LastUpdated The datetime when the emote was last modified
	LastUpdated time.Time `json:"last_updated"`
}
