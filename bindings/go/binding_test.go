package tree_sitter_authzed_test

import (
	"fmt"
	"testing"

	tree_sitter_authzed "github.com/aniamer/tree-sitter-authzed/bindings/go"
	tree_sitter "github.com/tree-sitter/go-tree-sitter"
)

func TestCanLoadGrammar(t *testing.T) {
	//schema, err := os.ReadFile("basic_authzed_schema.zed")
	//if err != nil {
	//	t.Fatalf("Failed to read file: %v", err)
	//}

	parser := tree_sitter.NewParser()
	defer parser.Close()
	language := tree_sitter.NewLanguage(tree_sitter_authzed.Language())
	if language == nil {
		t.Errorf("Error loading authzed-parser grammar")
	}

	fmt.Printf("languge %v", language.Version())
	parser.SetLanguage(language)
	tree := parser.Parse([]byte(`
		definition user {
			relation friend: user
			permission view = friend
    		}`), nil)

	if tree == nil {
		t.Error("Couldn't parse schema")
	}
	defer tree.Close()
	// root := tree.RootNode()
	// fmt.Println(root.ToSexp())
}
