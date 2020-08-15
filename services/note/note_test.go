package note

import (
	"testing"

	"github.com/donnol/jdnote/app"
)

func TestNoteGetHugoContent(t *testing.T) {
	n := &noteImpl{}
	content := n.getHugoContent("hah", "# hahahah", "2019-12-15", true, []string{}, []string{}, []string{})
	t.Logf("%s\n", content)
	content = n.getHugoContent("hah", "# hahahah", "2019-12-15", false, []string{"Go"}, []string{"Go"}, []string{"Go"})
	t.Logf("%s\n", content)
}

func TestPublish(t *testing.T) {
	n := &noteImpl{}
	ctx := app.DefaultCtx()
	if err := n.Publish(ctx, 45); err != nil {
		t.Fatal(err)
	}
}
