package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var targetDirs = []string{"cmd", "internal", "pkg", "sdk"}

type violation struct {
	Path    string
	Line    int
	Message string
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	violations, err := collectViolations()
	if err != nil {
		return err
	}
	if len(violations) == 0 {
		fmt.Println("env guard: ok")
		return nil
	}
	printViolations(violations)
	return fmt.Errorf("env guard: %d violation(s)", len(violations))
}

func collectViolations() ([]violation, error) {
	var violations []violation
	for _, root := range targetDirs {
		items, err := scanDir(root)
		if err != nil {
			return nil, err
		}
		violations = append(violations, items...)
	}
	return violations, nil
}

func scanDir(root string) ([]violation, error) {
	if err := ensureDir(root); err != nil {
		return nil, err
	}

	var violations []violation
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || filepath.Ext(path) != ".go" || allowedPath(path) {
			return err
		}
		items, itemErr := scanFile(path)
		if itemErr != nil {
			return itemErr
		}
		violations = append(violations, items...)
		return nil
	})
	return violations, err
}

func ensureDir(root string) error {
	_, err := os.Stat(root)
	if os.IsNotExist(err) {
		return nil
	}
	return err
}

func allowedPath(path string) bool {
	return strings.HasPrefix(filepath.ToSlash(path), "internal/config/")
}

func scanFile(path string) ([]violation, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}
	return fileViolations(path, fset, file), nil
}

func fileViolations(path string, fset *token.FileSet, file *ast.File) []violation {
	names := osImportNames(file)
	var violations []violation
	ast.Inspect(file, func(node ast.Node) bool {
		call, ok := node.(*ast.CallExpr)
		if ok {
			appendViolation(&violations, envCallViolation(path, fset, names, call))
		}
		return true
	})
	return violations
}

func osImportNames(file *ast.File) map[string]struct{} {
	names := make(map[string]struct{})
	for _, imp := range file.Imports {
		if imp.Path.Value == "\"os\"" {
			names[importName(imp)] = struct{}{}
		}
	}
	return names
}

func importName(imp *ast.ImportSpec) string {
	if imp.Name == nil {
		return "os"
	}
	return imp.Name.Name
}

func appendViolation(violations *[]violation, item *violation) {
	if item != nil {
		*violations = append(*violations, *item)
	}
}

func envCallViolation(
	path string,
	fset *token.FileSet,
	names map[string]struct{},
	call *ast.CallExpr,
) *violation {
	selector, ok := call.Fun.(*ast.SelectorExpr)
	if !ok || !isOSCall(names, selector) || !isEnvMethod(selector.Sel.Name) {
		return nil
	}
	return &violation{
		Path:    path,
		Line:    fset.Position(call.Pos()).Line,
		Message: fmt.Sprintf("direct %s call is forbidden outside config.Manager", selector.Sel.Name),
	}
}

func isOSCall(names map[string]struct{}, selector *ast.SelectorExpr) bool {
	ident, ok := selector.X.(*ast.Ident)
	if !ok {
		return false
	}
	_, ok = names[ident.Name]
	return ok
}

func isEnvMethod(name string) bool {
	return name == "Getenv" || name == "LookupEnv"
}

func printViolations(violations []violation) {
	for _, item := range violations {
		fmt.Printf("%s:%d: %s\n", item.Path, item.Line, item.Message)
	}
}
