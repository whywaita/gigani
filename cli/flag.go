package cli

import (
	"errors"
	"net/url"
)

func validateFlagURL(rawURL string) error {
	_, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return err
	}

	return nil
}

func validateFlag(format, sort string) error {
	if err := validateFlagOutputFormat(format); err != nil {
		return err
	}
	if err := validateFlagSort(sort); err != nil {
		return err
	}

	return nil
}

func validateFlagOutputFormat(outputFormat string) error {
	supportFormat := []string{"markdown", "json"}

	for _, format := range supportFormat {
		if outputFormat == format {
			// outputFormat is support!
			return nil
		}
	}

	return errors.New(outputFormat + " is unsupported output format")
}

func validateFlagSort(sortType string) error {
	supportType := []string{"post", "time"}

	for _, Type := range supportType {
		if sortType == Type {
			// sortType is support!
			return nil
		}
	}

	return errors.New(sortType + " is unsupported sort type")
}
