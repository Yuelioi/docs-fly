package post

import (
	"bytes"
	"docsfly/models"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func textContent(n *html.Node) string {
	var b bytes.Buffer
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			b.WriteString(c.Data)
		}
	}
	return b.String()
}

func loopNode(n *html.Node, entries *[]Toc) {
	if n.Type == html.ElementNode && (n.DataAtom == atom.H1 || n.DataAtom == atom.H2) {
		for _, a := range n.Attr {
			if a.Key == "id" {
				*entries = append(*entries, Toc{
					ID:    a.Val,
					Depth: uint(getDepth(n.DataAtom)),
					Title: textContent(n),
				})
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		loopNode(c, entries)
	}
}

func generateTOC(htmlContent string) ([]byte, error) {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return []byte{'[', ']'}, err
	}

	var entries []Toc

	loopNode(doc, &entries)

	jsonData, err := json.Marshal(entries)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return jsonData, err
	}

	if string(jsonData) == "null" {
		jsonData = []byte{'[', ']'}
	}

	return jsonData, nil
}

func buildFolderTree(folder *Chapter, categories []models.Entry, documents []models.Entry, currentDepth int) {

	var wg sync.WaitGroup

	// 添加文件到当前文件夹
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, doc := range documents {
			if strings.HasPrefix(doc.Filepath, folder.Filepath+"/") && doc.Depth == currentDepth {
				folder.Documents = append(folder.Documents, doc.MetaData)
			}

		}
	}()

	// 添加子文件夹到当前文件夹
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, cat := range categories {
			if strings.HasPrefix(cat.Filepath, folder.Filepath+"/") && cat.Depth == currentDepth {
				childFolder := Chapter{
					MetaData:  cat.MetaData,
					Filepath:  cat.MetaData.Filepath,
					Documents: make([]models.MetaData, 0),
					Children:  make([]Chapter, 0),
				}
				buildFolderTree(&childFolder, categories, documents, currentDepth+1)

				folder.Children = append(folder.Children, childFolder)
			}

		}
	}()

	wg.Wait()

}

func getDepth(node atom.Atom) int {
	switch node {
	case atom.H1:
		return 1
	case atom.H2:
		return 2
	case atom.H3:
		return 3
	case atom.H4:
		return 3
	default:
		return 0
	}

}
