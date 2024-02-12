package frankerfacez

import (
	"time"
)

type Emoticon struct {
	Id            int               `json:"id"`
	Name          string            `json:"name"`
	Height        int               `json:"height"`
	Width         int               `json:"width"`
	Public        bool              `json:"public"`
	Hidden        bool              `json:"hidden"`
	Modifier      bool              `json:"modifier"`
	ModifierFlags int               `json:"modifier_flags"`
	UsageCount    int               `json:"usage_count"`
	Urls          map[string]string `json:"urls"`
	CreatedAt     time.Time         `json:"created_at"`
	LastUpdated   time.Time         `json:"last_updated"`
}
