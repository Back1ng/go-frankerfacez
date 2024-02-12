package utils_test

import (
	"github.com/Back1ng/go-frankerfacez/utils"
	"net/url"
	"testing"
)

func TestReqToQueryValues(t *testing.T) {
	cases := []struct {
		name  string
		want  *url.Values
		given struct {
			Query string `json:"query"`
		}
	}{
		{
			name: "Convert basic struct to values",
			want: &url.Values{"query": []string{"test"}},
			given: struct {
				Query string `json:"query"`
			}{
				Query: "test",
			},
		},
		{
			name: "Convert basic empty struct to values",
			want: &url.Values{"query": []string{""}},
			given: struct {
				Query string `json:"query"`
			}{
				Query: "",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			values, err := utils.ReqToQueryValues(tc.given)
			if err != nil {
				t.Error(err)
			}

			if values.Encode() != tc.want.Encode() {
				t.Error("Value is not equal to wanted value")
			}
		})
	}
}
