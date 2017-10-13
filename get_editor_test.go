package cryptex

import (
	"os"
	"reflect"
	"testing"
)

func TestGetEditorEnvPnzrEditor(t *testing.T) {
	os.Unsetenv("DEFAULT_EDITOR")
	os.Unsetenv("EDITOR")
	tests := []struct {
		in     string
		editor string
		args   []string
	}{
		{"vi", "vi", []string{}},
		{"", "nano", []string{}},
		{"atom -f true", "atom", []string{"-f", "true"}},
	}

	for _, test := range tests {
		func() {
			defer func() {
				os.Unsetenv("DEFAULT_EDITOR")
				os.Unsetenv("EDITOR")
			}()

			os.Setenv("DEFAULT_EDITOR", test.in)

			editor, args := getEditor()
			if !reflect.DeepEqual(args, test.args) {
				t.Fatalf("want %q, but %q:", test.args, args)
			}

			if editor != test.editor {
				t.Fatalf("want %q, but %q:", test.editor, editor)
			}
		}()
	}
}

func TestGetEditorEnvEditor(t *testing.T) {
	os.Unsetenv("DEFAULT_EDITOR")
	os.Unsetenv("EDITOR")
	tests := []struct {
		in     string
		editor string
		args   []string
	}{
		{"vi", "vi", []string{}},
		{"", "nano", []string{}},
		{"atom -f true", "atom", []string{"-f", "true"}},
	}

	for _, test := range tests {
		func() {
			defer func() {
				os.Unsetenv("DEFAULT_EDITOR")
				os.Unsetenv("EDITOR")
			}()

			os.Setenv("EDITOR", test.in)

			editor, args := getEditor()

			if !reflect.DeepEqual(args, test.args) {
				t.Fatalf("want %q, but %q:", test.args, args)
			}

			if editor != test.editor {
				t.Fatalf("want %q, but %q:", test.editor, editor)
			}
		}()
	}
}

func TestGetEditorNoEnv(t *testing.T) {
	os.Unsetenv("DEFAULT_EDITOR")
	os.Unsetenv("EDITOR")

	editor, args := getEditor()

	if !reflect.DeepEqual(args, []string{}) {
		t.Fatalf("want %q, but %q:", []string{}, args)
	}

	if editor != "nano" {
		t.Fatalf("want %q, but %q:", "nano", editor)
	}

}
