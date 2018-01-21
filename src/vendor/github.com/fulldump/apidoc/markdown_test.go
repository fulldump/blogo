package apidoc

import (
	"fmt"
	"testing"
)

func Test_md_crop_tabs_happy_path(t *testing.T) {
	text := `
			This is my text comment with some blank lines

			And some indented code:

			´´´json

				{
					"name": "fulanitez",
					"address": {
						"street": "Elm",
						"number": 7
					}
				}

			´´´
	`

	expected := `
This is my text comment with some blank lines

And some indented code:

´´´json

	{
		"name": "fulanitez",
		"address": {
			"street": "Elm",
			"number": 7
		}
	}

´´´
	`

	processed := md_crop_tabs(text)

	if expected != processed {
		fmt.Println(expected)
		fmt.Println(processed)
		t.Error("Result does not match.")
	}

}

func Test_md_crop_tabs_first_line(t *testing.T) {
	text := `This is my text comment with some blank lines

			And some indented code:
	`

	expected := `This is my text comment with some blank lines

And some indented code:
	`

	processed := md_crop_tabs(text)

	if expected != processed {
		fmt.Println(expected)
		fmt.Println(processed)
		t.Error("Result does not match.")
	}

}
