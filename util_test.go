package blockfrost

import (
	"net/http"
	"testing"
)

func TestFormatParams(t *testing.T) {
	tests := []struct {
		query APIPagingParams
		want  string
	}{
		{APIPagingParams{Count: 5}, "count=5"},
		{APIPagingParams{Page: 10}, "page=10"},
		{APIPagingParams{}, ""},
		{APIPagingParams{Order: "asc"}, "order=asc"},
		{APIPagingParams{Count: 5, Page: 10}, "count=5&page=10"},
		{APIPagingParams{Count: 5, Page: 10, Order: "desc"}, "count=5&order=desc&page=10"},
	}
	req, err := http.NewRequest(http.MethodGet, "/go", nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, tt := range tests {
		tt, want := tt, tt.want
		v := req.URL.Query()
		t.Run("", func(t *testing.T) {
			got := formatParams(v, tt.query).Encode()
			if got != want {
				t.Fatalf("expected %s got %s", want, got)
			}
		})
	}
}
