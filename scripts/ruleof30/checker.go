package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
)

const (
	maxFileLines    = 300
	maxFuncLines    = 30
	maxStructFields = 20
)

type violation struct {
	Path    string
	Line    int
	Message string
}

func (v violation) String() string {
	if v.Line > 0 {
		return fmt.Sprintf("%s:%d: %s", v.Path, v.Line, v.Message)
	}
	return fmt.Sprintf("%s: %s", v.Path, v.Message)
}

func collectViolations(root string) ([]violation, error) {
	files, err := listGoFiles(root)
	if err != nil {
		return nil, err
	}

	var violations []violation
	for _, path := range files {
		fileViolations, err := checkFile(path)
		if err != nil {
			return nil, err
		}
		violations = append(violations, fileViolations...)
	}
	return violations, nil
}

func listGoFiles(root string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() && shouldSkipDir(d.Name()) {
			return filepath.SkipDir
		}
		if d.IsDir() || filepath.Ext(path) != ".go" {
			return nil
		}
		files = append(files, path)
		return nil
	})
	sort.Strings(files)
	return files, err
}

func shouldSkipDir(name string) bool {
	switch name {
	case ".git", ".idea", ".memory", "bin", "vendor":
		return true
	default:
		return false
	}
}

func checkFile(path string) ([]violation, error) {
	source, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}
	return checkSource(path, source)
}

func checkSource(path string, source []byte) ([]violation, error) {
	if isGenerated(source) {
		return nil, nil
	}

	fset, file, err := parseGoFile(path, source)
	if err != nil {
		return nil, err
	}

	violations := fileViolations(path, source)
	violations = append(violations, nodeViolations(path, fset, file)...)
	return violations, nil
}

func parseGoFile(path string, source []byte) (*token.FileSet, *ast.File, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, source, 0)
	if err != nil {
		return nil, nil, fmt.Errorf("parse %s: %w", path, err)
	}
	return fset, file, nil
}

func fileViolations(path string, source []byte) []violation {
	lines := lineCount(source)
	if lines <= maxFileLines {
		return nil
	}

	return []violation{{
		Path:    path,
		Line:    1,
		Message: fmt.Sprintf("file has %d lines (max %d)", lines, maxFileLines),
	}}
}

func lineCount(source []byte) int {
	if len(source) == 0 {
		return 0
	}
	return bytes.Count(source, []byte{'\n'}) + 1
}

func isGenerated(source []byte) bool {
	for _, line := range bytes.SplitN(source, []byte{'\n'}, 6) {
		if bytes.Contains(line, []byte("Code generated")) {
			return true
		}
	}
	return false
}

func nodeViolations(path string, fset *token.FileSet, file *ast.File) []violation {
	var violations []violation
	ast.Inspect(file, func(node ast.Node) bool {
		addViolation(&violations, funcDeclViolation(path, fset, node))
		addViolation(&violations, funcLitViolation(path, fset, node))
		addViolation(&violations, structViolation(path, fset, node))
		return true
	})
	return violations
}

func addViolation(violations *[]violation, violation *violation) {
	if violation != nil {
		*violations = append(*violations, *violation)
	}
}

func funcDeclViolation(path string, fset *token.FileSet, node ast.Node) *violation {
	fn, ok := node.(*ast.FuncDecl)
	if !ok {
		return nil
	}

	lines := nodeLineSpan(fset, fn)
	if lines <= maxFuncLines {
		return nil
	}

	return &violation{
		Path:    path,
		Line:    fset.Position(fn.Pos()).Line,
		Message: fmt.Sprintf("function %s has %d lines (max %d)", fn.Name.Name, lines, maxFuncLines),
	}
}

func funcLitViolation(path string, fset *token.FileSet, node ast.Node) *violation {
	fn, ok := node.(*ast.FuncLit)
	if !ok {
		return nil
	}

	lines := nodeLineSpan(fset, fn)
	if lines <= maxFuncLines {
		return nil
	}

	return &violation{
		Path:    path,
		Line:    fset.Position(fn.Pos()).Line,
		Message: fmt.Sprintf("function literal has %d lines (max %d)", lines, maxFuncLines),
	}
}

func structViolation(path string, fset *token.FileSet, node ast.Node) *violation {
	structType, ok := node.(*ast.StructType)
	if !ok {
		return nil
	}

	fields := structFieldCount(structType)
	if fields <= maxStructFields {
		return nil
	}

	return &violation{
		Path:    path,
		Line:    fset.Position(structType.Pos()).Line,
		Message: fmt.Sprintf("struct has %d fields (max %d)", fields, maxStructFields),
	}
}

func nodeLineSpan(fset *token.FileSet, node ast.Node) int {
	start := fset.Position(node.Pos()).Line
	end := fset.Position(node.End()).Line
	return end - start + 1
}

func structFieldCount(structType *ast.StructType) int {
	total := 0
	for _, field := range structType.Fields.List {
		if len(field.Names) == 0 {
			total++
			continue
		}
		total += len(field.Names)
	}
	return total
}

func printViolations(violations []violation) {
	for _, violation := range violations {
		fmt.Println(violation.String())
	}
}
