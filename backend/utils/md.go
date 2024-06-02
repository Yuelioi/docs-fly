package utils

import (
	"bytes"
	"docsfly/models"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"go.abhg.dev/goldmark/frontmatter"
	netHtml "golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	fences "github.com/stefanfritsch/goldmark-fences"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func ReadMd(filepath string) ([]byte, error) {
	// 读取Markdown文件内容

	mdBytes, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	return mdBytes, nil
}

// 读取文件的Meta信息
func ReadMetas(path string, info os.FileInfo) (*[]models.MetaData, error) {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, errors.New("文件不存在")
	}

	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading Meta:", err)
		return nil, err
	}
	metadatas := make([]models.MetaData, 0)
	err = json.Unmarshal(content, &metadatas)

	if err != nil {
		return nil, err
	}

	return &metadatas, nil
}

func SearchMeta(datas *[]models.MetaData, info os.FileInfo, order uint) *models.MetaData {
	for _, data := range *datas {
		if data.Name == info.Name() {
			return &data
		}
	}
	return CreateMeta(info, order)
}

// 读取Markdown文件的Meta信息
func ReadMarkdownMeta(path string, info os.FileInfo, order uint) (metadata *models.MetaData, err error) {

	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return CreateMeta(info, order), err
	}

	markdown := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
		),
	)

	var buf bytes.Buffer
	context := parser.NewContext()
	if err := markdown.Convert(content, &buf, parser.WithContext(context)); err != nil {
		return nil, err
	}
	metaReader := meta.Get(context)

	metadata = &models.MetaData{
		Name:   Transform[string](metaReader["Name"]),
		Title:  Transform[string](metaReader["Title"]),
		Order:  Transform[uint](metaReader["Order"]),
		Status: Transform[uint](metaReader["Status"]),
	}
	return metadata, nil
}

func GenerateMeta(docs models.Document) string {
	return fmt.Sprintf(`---
display_name: %v
order: %v
---
`, docs.Title, docs.Order)

}

func GetMeta(metaReader map[string]interface{}) string {
	metaString := "---\n"
	for key, value := range metaReader {
		switch v := value.(type) {
		case []interface{}:
			metaString += key + ":\n"
			for _, item := range v {
				metaString += fmt.Sprintf("    - %v\n", item)
			}

		default:
			metaString += fmt.Sprintf("%s: %v\n", key, value)
		}
	}
	return metaString + "---"
}

func InitMarkdownMeta(docs models.Document) error {

	content, err := os.ReadFile(docs.Filepath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}
	re := regexp.MustCompile(`(?s)\A---.*?---`)

	match := re.Find(content)
	isMatch := len(match) > 0

	if isMatch {
		contentWithoutMeta := re.ReplaceAll(content, []byte(""))

		markdown := goldmark.New(
			goldmark.WithExtensions(
				meta.Meta,
			),
		)

		var buf bytes.Buffer
		context := parser.NewContext()
		if err := markdown.Convert(content, &buf, parser.WithContext(context)); err != nil {
			return err
		}
		metaReader := meta.Get(context)
		if metaReader["NameZH"] != nil {
			metaReader["display_name"] = metaReader["NameZH"]
		} else {
			metaReader["display_name"] = docs.Name
		}

		if metaReader["Order"] != nil {
			metaReader["order"] = metaReader["Order"]
		} else {
			metaReader["order"] = docs.Order
		}

		metaString := GetMeta(metaReader)
		os.WriteFile(docs.Filepath, []byte(metaString+string(contentWithoutMeta)), 0644)

	} else {
		meta := GenerateMeta(docs)
		fmt.Printf("%v\n", meta+string(content))
		os.WriteFile(docs.Filepath, []byte(meta+string(content)), 0644)
	}

	return nil
}

// 获取文件名(不包含后缀)
func PureFileName(file string) (fileNameWithoutExt string) {
	fileName := filepath.Base(file)
	fileNameWithoutExt = fileName[:len(fileName)-len(filepath.Ext(fileName))]
	return
}

// 初始化Meta信息
func CreateMeta(info os.FileInfo, order uint) *models.MetaData {
	return &models.MetaData{
		Name:  info.Name(),
		Title: PureFileName(info.Name()),
		Order: order,
	}
}

// 修改meta数据,把零值替换成替代值
func UpdateMeta(meta *models.MetaData, Name, Title string, Order uint, Status uint) {

	if IsZeroType(meta.Name) {
		meta.Name = Name
	}
	if IsZeroType(meta.Title) {
		meta.Title = Title
	}
	if meta.Order == 0 {
		meta.Order = Order
	}
	if IsZeroType(meta.Status) {
		meta.Status = Status
	}

}

type preWrapper struct {
	start func(code bool, styleAttr string) string
	end   func(code bool) string
}

func (p preWrapper) Start(code bool, styleAttr string) string {
	return p.start(code, styleAttr)
}

func (p preWrapper) End(code bool) string {
	return p.end(code)
}

func codeAddPreClass(ctx highlighting.CodeBlockContext) (options []chromahtml.Option) {
	lang, ok := ctx.Language()
	if ok {
		codePreWrapper := preWrapper{
			start: func(code bool, styleAttr string) string {
				if code {
					return fmt.Sprintf(`<div class="code-block"><pre class="language-%s" data-language="%[1]s" tabindex="0"%[2]s><code>`, string(lang), styleAttr)
				}

				return fmt.Sprintf(`<pre tabindex="0"%s>`, styleAttr)
			},
			end: func(code bool) string {
				if code {
					return `</code></pre><button class="copy-button" >copy</button></div>`
				}

				return `</pre>`
			},
		}

		op := chromahtml.WithPreWrapper(codePreWrapper)
		options = append(options, op)
	}

	return
}

func MarkdownToHTML(source []byte) []byte {

	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM,
			&frontmatter.Extender{},
			&fences.Extender{},
			highlighting.NewHighlighting(
				highlighting.WithStyle("base16-snazzy"),

				highlighting.WithCodeBlockOptions(codeAddPreClass),

				highlighting.WithFormatOptions(
					chromahtml.BaseLineNumber(1),
					chromahtml.LineNumbersInTable(true),

					chromahtml.WithAllClasses(true),
					chromahtml.WithClasses(true),

					chromahtml.WithLineNumbers(true),
				),
			),
		),

		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
			html.WithUnsafe(),
		),
	)
	var buf bytes.Buffer
	if err := md.Convert(source, &buf); err != nil {
		panic(err)
	}

	return buf.Bytes()
}

// Entry represents a single entry in the table of contents.
type Entry struct {
	ID    string
	Level int
	Title string
}

// ByID is a sortable list of entries.
type ByID []Entry

func (b ByID) Len() int           { return len(b) }
func (b ByID) Less(i, j int) bool { return b[i].ID < b[j].ID }
func (b ByID) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

// GenerateTOC generates a table of contents for the given HTML document.
func GenerateTOC(htmlContent string) (string, error) {
	doc, err := netHtml.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return "", err
	}

	var entries []Entry
	var f func(n *netHtml.Node)
	f = func(n *netHtml.Node) {
		if n.Type == netHtml.ElementNode && (n.DataAtom == atom.H1 || n.DataAtom == atom.H2 || n.DataAtom == atom.H3) {
			for _, a := range n.Attr {
				if a.Key == "id" {
					entries = append(entries, Entry{
						Level: int(n.DataAtom - atom.H1 + 1),
						ID:    a.Val,
						Title: textContent(n),
					})
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	// Sort entries by level and then by title
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Level != entries[j].Level {
			return entries[i].Level < entries[j].Level
		}
		return entries[i].Title < entries[j].Title
	})

	// Generate the TOC HTML
	toc := new(bytes.Buffer)
	tmpl := template.Must(template.New("toc").Parse(`<nav>
{{ range . }}<a href="#{{ .ID }}">{{ repeat .Level "-" }} {{ .Title }}</a>{{ end }}
</nav>`))
	tmpl.Funcs(template.FuncMap{
		"repeat": func(count int, str string) string {
			r := strings.Repeat(str, count)
			return strings.TrimSpace(r)
		},
	})
	if err := tmpl.Execute(toc, entries); err != nil {
		return "", err
	}

	return toc.String(), nil
}

func textContent(n *netHtml.Node) string {
	var b bytes.Buffer
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == netHtml.TextNode {
			b.WriteString(c.Data)
		}
	}
	return b.String()
}
