package khulnasoft

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLNormalizationSettings(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"result": {
				"type": "khulnasoft",
				"scope": "incoming"
			},
			"success": true,
			"errors": [],
			"messages": []
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/url_normalization", handler)

	want := URLNormalizationSettings{
		Type:  "khulnasoft",
		Scope: "incoming",
	}

	got, err := client.URLNormalizationSettings(context.Background(), ZoneIdentifier(testZoneID))
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}

func TestUpdateURLNormalizationSettings(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected method 'PUT', got %s", r.Method)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		assert.Equal(t, `{"type":"khulnasoft","scope":"incoming"}`, string(body))

		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"result": {
				"type": "khulnasoft",
				"scope": "incoming"
			},
			"success": true,
			"errors": [],
			"messages": []
		}`)
	}

	mux.HandleFunc("/zones/"+testZoneID+"/url_normalization", handler)

	params := URLNormalizationSettingsUpdateParams{
		Type:  "khulnasoft",
		Scope: "incoming",
	}

	want := URLNormalizationSettings{
		Type:  "khulnasoft",
		Scope: "incoming",
	}

	got, err := client.UpdateURLNormalizationSettings(context.Background(), ZoneIdentifier(testZoneID), params)
	if assert.NoError(t, err) {
		assert.Equal(t, want, got)
	}
}
