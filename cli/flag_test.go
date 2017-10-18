package cli

import (
	"errors"
	"testing"
)

func TestValidateFlagURL(t *testing.T) {
	goodURL := "https://github.com"
	badURL := "htttps://github.com/"

	goodErr := validateFlagURL(goodURL)
	badErr := validateFlagURL(badURL)

	// catch Good URL, maybe return nil
	if goodErr != nil {
		t.Fatal(goodErr)
	}

	// catch Bad URL, maybe return err (unsupported scheme)
	errorMsg := errors.New(`Get htttps://github.com/: unsupported protocol scheme "htttps"`)
	if badErr == errorMsg {
		t.Fatal(badErr)
	}
}

func TestValidateFlagOutputFormat(t *testing.T) {
	supportedFormat := "markdown"
	unsupportedFormat := "html"

	goodErr := validateFlagOutputFormat(supportedFormat)
	badErr := validateFlagOutputFormat(unsupportedFormat)

	// catch supported format, maybe return nil
	if goodErr != nil {
		t.Fatal(goodErr)
	}

	// catch unsupported Format, maybe return err (not nil)
	if badErr == nil {
		t.Fatal(badErr)
	}
}
