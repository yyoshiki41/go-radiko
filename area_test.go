package radiko

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestAreaID(t *testing.T) {
	areaID, err := AreaID()
	if err != nil {
		t.Errorf("Failed to download player.swf: %s", err)
	}
	if areaID == "" {
		t.Errorf("Invalid area id: %s", areaID)
	}
}

func TestProcessSpanNode(t *testing.T) {
	s := `document.write('<span class="JP13">TOKYO JAPAN</span>');`
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		t.Errorf("Parse HTML: %s", err)
	}

	areaID := processSpanNode(doc)
	if areaID != "JP13" {
		t.Errorf("Failed to process span node.\nAreaID: %s",
			areaID)
	}
}
