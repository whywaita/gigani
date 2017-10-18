package cli

import (
	"errors"
	"net/url"
)

func validateFlag(url, outputFormat string) error {
	// validation flag
	var err error

	err = validateFlagURL(url)
	if err != nil {
		return err
	}

	err = validateFlagOutputFormat(outputFormat)
	if err != nil {
		return err
	}

	// maybe good
	return nil
}

func validateFlagURL(rawURL string) error {
	_, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return err
	}

	return nil
}

func validateFlagOutputFormat(outputFormat string) error {
	supportFormat := [1]string{"markdown"}

	for _, format := range supportFormat {
		if outputFormat == format {
			// outputFormat is support!
			return nil
		}
	}

	err := errors.New(outputFormat + " is unsupported output format")
	return err
}
