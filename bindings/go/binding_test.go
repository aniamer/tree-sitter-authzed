package tree_sitter_authzed_test

import (
	"fmt"
	"os"
	"testing"

	tree_sitter_authzed "github.com/aniamer/tree-sitter-authzed/bindings/go"
	tree_sitter "github.com/tree-sitter/go-tree-sitter"
)

func TestCanLoadGrammar(t *testing.T) {
	schema, err := os.ReadFile("basic_authzed_schema.zed")
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	parser := tree_sitter.NewParser()
	defer parser.Close()
	language := tree_sitter.NewLanguage(tree_sitter_authzed.Language())
	if language == nil {
		t.Errorf("Error loading authzed-parser grammar")
		return
	}
	err = parser.SetLanguage(language)
	if err != nil {
		t.Error(err)
		return
	}
	tree := parser.Parse(schema, nil)

	if tree == nil {
		t.Error("Couldn't parse schema")
		return
	}
	defer tree.Close()

	rootNode := tree.RootNode()
	fmt.Printf("%s\n", rootNode.ToSexp())

	query, query_err := tree_sitter.NewQuery(language, `((definition) @identifier)`)
	if query_err != nil {
		t.Error(err)
	}

	defer query.Close()

	qc := tree_sitter.NewQueryCursor()
	defer qc.Close()

	captures := qc.Captures(query, tree.RootNode(), schema)

	for match, index := captures.Next(); match != nil; match, index = captures.Next() {
		fmt.Printf(
			"Capture %d: %s\n",
			index,
			match.Captures[index].Node.Utf8Text(schema),
		)
	}
}
