package utils_test

import (
	"github.com/Back1ng/go-frankerfacez/frankerfacez"
	"github.com/Back1ng/go-frankerfacez/utils"
	"net/url"
	"testing"
)

func TestReqToQueryValues(t *testing.T) {
	cases := []struct {
		name  string
		want  *url.Values
		given frankerfacez.ApiV1EmotesRequest
	}{
		{
			name: "Convert basic struct to values",
			want: &url.Values{"q": []string{"test"}},
			given: frankerfacez.ApiV1EmotesRequest{
				Query: "test",
			},
		},
		{
			name: "Convert basic empty struct to values",
			want: &url.Values{"q": []string{""}},
			given: frankerfacez.ApiV1EmotesRequest{
				Query: "",
			},
		},
		{
			name: "Convert boolean true is a string true",
			want: &url.Values{"q": []string{"KEKW"}, "animated": []string{"true"}},
			given: frankerfacez.ApiV1EmotesRequest{
				Query:    "KEKW",
				Animated: true,
			},
		},
		{
			name: "Omit boolean when false",
			want: &url.Values{"q": []string{"KEKW"}},
			given: frankerfacez.ApiV1EmotesRequest{
				Query:    "KEKW",
				Animated: false,
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
				t.Errorf("Value %v is not equal to wanted value %v", values.Encode(), tc.want.Encode())
			}
		})
	}
}
