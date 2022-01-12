package tutorials

import (
	"fyne.io/fyne/v2"
)

// Tutorial defines the data structure for a tutorial
type Tutorial struct {
	Title, Intro string
	View         func(w fyne.Window) fyne.CanvasObject
}

var (
	// Tutorials defines the metadata for each tutorial
	Tutorials = map[string]Tutorial{

		"Base64": {"Encode/Decode",
			"Use Base64 encode/decode text.",
			encodeTab,
		},
		"Sign": {"Sign",
			"sign text.",
			signTab,
		},
	}

	// TutorialIndex  defines how our tutorials should be laid out in the index tree
	TutorialIndex = map[string][]string{
		"": {"Base64", "Sign"},
	}
)
