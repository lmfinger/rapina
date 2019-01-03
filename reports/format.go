package reports

import (
	"encoding/json"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// Used by style
const (
	DEFAULT = iota + 1

	// Number format
	NUMBER
	INDEX
	PERCENT

	EMPTY

	// Text position
	LEFT
	RIGHT
	CENTER
)

// formatFont directly maps the styles settings of the fonts.
type formatFont struct {
	Bold      bool   `json:"bold"`
	Italic    bool   `json:"italic"`
	Underline string `json:"underline"`
	Family    string `json:"family"`
	Size      int    `json:"size"`
	Color     string `json:"color"`
}

type formatAlignment struct {
	Horizontal      string `json:"horizontal"`
	Indent          int    `json:"indent"`
	JustifyLastLine bool   `json:"justify_last_line"`
	ReadingOrder    uint64 `json:"reading_order"`
	RelativeIndent  int    `json:"relative_indent"`
	ShrinkToFit     bool   `json:"shrink_to_fit"`
	TextRotation    int    `json:"text_rotation"`
	Vertical        string `json:"vertical"`
	WrapText        bool   `json:"wrap_text"`
}

type formatBorder struct {
	Type  string `json:"type"`
	Color string `json:"color"`
	Style int    `json:"style"`
}

// formatStyle directly maps the styles settings of the cells.
type formatStyle struct {
	Border []formatBorder `json:"border"`
	Fill   struct {
		Type    string   `json:"type"`
		Pattern int      `json:"pattern"`
		Color   []string `json:"color"`
		Shading int      `json:"shading"`
	} `json:"fill"`
	Font       *formatFont      `json:"font"`
	Alignment  *formatAlignment `json:"alignment"`
	Protection *struct {
		Hidden bool `json:"hidden"`
		Locked bool `json:"locked"`
	} `json:"protection"`
	NumFmt        int     `json:"number_format"`
	DecimalPlaces int     `json:"decimal_places"`
	CustomNumFmt  *string `json:"custom_number_format"`
	Lang          string  `json:"lang"`
	NegRed        bool    `json:"negred"`
}

//
// newFormat provides a struct to create style for cells
//
func newFormat(format int, position int, bold bool) (f *formatStyle) {
	f = &formatStyle{}

	custom := ""
	switch format {
	case PERCENT:
		custom = "0%;-0%;- "
	case INDEX:
		custom = "0.0;-0.0;-"
	case NUMBER:
		custom = "_-* #,##0,_-;_-* (#,##0,);_-* \"-\"_-;_-@_-"
	}

	if custom != "" {
		f.CustomNumFmt = &custom
	}

	switch position {
	case RIGHT:
		f.Alignment = &formatAlignment{Horizontal: "right"}
	case CENTER:
		f.Alignment = &formatAlignment{Horizontal: "center"}
	}

	if bold {
		f.Font = &formatFont{Bold: true}
	}

	return
}

func (f *formatStyle) bold(enabled bool) {
	f.Font = &formatFont{Bold: enabled}
}

func (f *formatStyle) size(s int) {
	f.Font = &formatFont{Size: s}
}

func (f formatStyle) newStyle(e *excelize.File) (style int) {
	json, err := json.Marshal(f)
	if err == nil {
		style, err = e.NewStyle(string(json))
		if err != nil {
			style = 0
		}
	}

	return
}
