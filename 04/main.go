package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Tree map[string]Tree

func (tree Tree) Add(path string, absDir string, dir string, level int) {
	if path == absDir {
		return
	}

	subDir := strings.Split(path, absDir)
	segments := strings.Split(subDir[1], "/")
	segments[0] = dir
	if level == 0 || len(segments) < level+2 {
		tree.add(segments)
	}
}

func (tree Tree) add(segments []string) {
	if len(segments) == 0 {
		return
	}

	nextTree, ok := tree[segments[0]]
	if !ok {
		nextTree = Tree{}
		tree[segments[0]] = nextTree
	}

	nextTree.add(segments[1:])
}

func (tree Tree) Fprint(w io.Writer, root bool, padding string) {
	if tree == nil {
		return
	}

	index := 0

	for k, v := range tree {
		fmt.Fprintf(w, "%s%s\n", padding+getPadding(root, getBoxType(index, len(tree))), k)
		v.Fprint(w, false, padding+getPadding(root, getBoxTypeExternal(index, len(tree))))
		index++
	}
}

type BoxType int

const (
	Regular BoxType = iota
	Last
	AfterLast
	Between
)

func (boxType BoxType) String() string {
	switch boxType {
	case Regular:
		return "\u251c" // ├
	case Last:
		return "\u2514" // └
	case AfterLast:
		return " "
	case Between:
		return "\u254E" // ╎
	default:
		panic("invalid box type")
	}
}

func getBoxType(index int, len int) BoxType {
	if index+1 == len {
		return Last
	} else if index+1 > len {
		return AfterLast
	}
	return Regular
}

func getBoxTypeExternal(index int, len int) BoxType {
	if index+1 == len {
		return AfterLast
	}
	return Between
}

func getPadding(root bool, boxType BoxType) string {
	if root {
		return ""
	}

	return boxType.String() + " "
}

func main() {
	var dir string
	flag.StringVar(&dir, "d", ".", "Directory")
	level := flag.Int("L", 0, "Descend only level directories deep.")
	flag.Parse()

	absDir, derr := filepath.Abs(dir)

	if derr != nil {
		log.Println(derr)
	}

	tree := Tree{}
	err := filepath.WalkDir(absDir,
		func(path string, info os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			tree.Add(path, absDir, dir, *level)
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	tree.Fprint(os.Stdout, true, "")
}
