package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/STLnick/import-analyzer/utils"
)

type ImportNode struct {
	path        string
	isNode      bool
	occurrences int
	children    *[]ImportNode
}

func (ig *ImportNode) Print(depth int) {
	pad := strings.Repeat(" ", depth)

	fmt.Printf("%s{\n", pad)
	fmt.Printf("%s  path: %s\n", pad, ig.path)
	fmt.Printf("%s  isNode: %v\n", pad, ig.isNode)
	fmt.Printf("%s  occurrences: %d\n", pad, ig.occurrences)
	fmt.Printf("%s  children:\n", pad)
	for _, child := range *(ig.children) {
		child.Print(depth + 2)
	}
	fmt.Printf("%s}\n", pad)
}

func NewImportNode(path string, occurrences int, isNode bool) ImportNode {
	slice := make([]ImportNode, 0, 3)
	return ImportNode{
		path:        path,
		occurrences: occurrences,
		isNode:      isNode,
		children:    &slice,
	}
}

type ImportResult struct {
    count int
    paths *[]string
}

func processStatementAsImportNode(tree []ImportNode, s string) []ImportNode {
	if !strings.Contains(s, "import") {
		return tree
	}

	var currentGroup *ImportNode
	var workingPath string
	var isLastChunk bool
	splitPath := utils.ExtractPathParts(s)

	for splitIdx, pathChunk := range splitPath {
		if splitIdx == 0 {
			workingPath = pathChunk
			rootIdx := slices.IndexFunc(tree, func(ig ImportNode) bool {
				return ig.path == workingPath
			})

			if rootIdx == -1 {
				ig := NewImportNode(workingPath, 1, false)
				currentGroup = &ig
				tree = append(tree, ig)
			} else {
				currentGroup = &tree[rootIdx]
				(*currentGroup).occurrences += 1
			}
		} else {
			isLastChunk = splitIdx == len(splitPath)-1
			workingPath += "/" + pathChunk
			childIdx := slices.IndexFunc(*((*currentGroup).children), func(c ImportNode) bool {
				return c.path == workingPath
			})

			if childIdx == -1 {
				childIg := NewImportNode(workingPath, 1, isLastChunk)
				*(*currentGroup).children = append(*(*currentGroup).children, childIg)
				currentGroup = &childIg
			} else {
				currentGroup = &(*((*currentGroup).children))[childIdx]
				(*currentGroup).occurrences += 1
			}
		}
	}

	return tree
}

func sortChildren(node ImportNode, listMap *[]ImportResult) {
    var idx int

    for _, child := range *(node.children) {
        idx = slices.IndexFunc(*listMap, func(ir ImportResult) bool {
            return ir.count == child.occurrences
        })
        if idx == -1 {
            ir := ImportResult{count: child.occurrences, paths: &[]string{child.path}}
            *listMap = append(*listMap, ir)
        } else {
            ir := (*listMap)[idx]
            *(ir.paths) = append(*(ir.paths), child.path)
        }
       
        if len(*(child.children)) > 0 {
            sortChildren(child, listMap)
        }
    }
}

func sortByHighestOccurrences(tree []ImportNode) []ImportResult {
    listMap := make([]ImportResult, 0, 5)

    for _, root := range tree {
        rootRes := ImportResult{
            count: root.occurrences,
            paths: &[]string{root.path},
        }
        listMap = append(listMap, rootRes)
        sortChildren(root, &listMap)
    }

    return listMap
}

func main() {
	var tree []ImportNode
	var lines []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	slices.Sort(lines)
	for _, line := range lines {
		tree = processStatementAsImportNode(tree, line)
	}

    resultMap := sortByHighestOccurrences(tree)
    fmt.Println("Result Map:")
    for _, v := range resultMap {
        fmt.Printf(" :: count=%d paths=(%v)\n", v.count, *(v.paths))
    }
}
