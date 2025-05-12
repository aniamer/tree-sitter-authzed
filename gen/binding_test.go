package gen

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_authzed "github.com/tree-sitter/tree-sitter-authzed/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_authzed.Language())
	if language == nil {
		t.Errorf("Error loading authzed-parser grammar")
	}
}
