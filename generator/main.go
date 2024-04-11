//go:build generator

package main

import (
	"flag"
	"fmt"
	"github.com/sleepinggenius2/gosmi/smi"
	"os"
	"os/exec"
	"strings"

	"github.com/sleepinggenius2/gosmi"
	"github.com/sleepinggenius2/gosmi/types"
)

type VerbStructure struct {
	Name        string
	Description string
	Columns     []*EnrichNode
	Index       []*EnrichNode
}

func (v *VerbStructure) WriteComments() string {
	return fmt.Sprintf(
		"// %s %s\n",
		strings.Title(v.Name),
		strings.Replace(v.Description, "\n", "\n// ", -1),
	)

}
func (v *VerbStructure) WriteFile() error {
	file, err := os.Create(fmt.Sprintf("../beans/%s.go", v.Name))
	if err != nil {
		return err
	}
	defer file.Close()
	var sb strings.Builder
	sb.WriteString("package beans\n\n")
	sb.WriteString("import (\n")
	sb.WriteString("\t\"fmt\"\n")
	sb.WriteString("\t\"reflect\"\n")
	sb.WriteString(")\n\n")
	sb.WriteString(v.WriteStructMain())
	sb.WriteString("\n\n")
	sb.WriteString(v.WriteStructParams())
	_, err = file.WriteString(sb.String())

	return err
}

func (v *VerbStructure) WriteStructMain() string {
	var sb strings.Builder
	nameTitled := strings.Title(v.Name)
	for _, indexNode := range v.Index {
		sb.WriteString(writeEnum(nameTitled, indexNode) + "\n")
	}

	sb.WriteString(v.WriteComments())
	sb.WriteString(fmt.Sprintf("type %s struct {\n", strings.Title(v.Name)))
	indexNames := make([]string, 0)
	conditionalZeroIndex := make([]string, 0)
	for _, column := range v.Index {
		titledCName := strings.Title(column.Name)
		indexNames = append(indexNames, titledCName)
		conditionalZeroIndex = append(conditionalZeroIndex, fmt.Sprintf("reflect.ValueOf(c.%s).IsZero()", titledCName))
		comments := column.WriteComments()
		comments = strings.Replace(strings.TrimSuffix(comments, "\n"), "\n", "\n\t", -1)
		sb.WriteString("\t" + comments + "\n")
		sb.WriteString(fmt.Sprintf("\t%s %s\n", strings.Title(column.Name), column.TypeString(nameTitled)))
	}
	sb.WriteString(fmt.Sprintf("\tParams *%sParams\n", nameTitled))
	sb.WriteString("}\n\n")

	sb.WriteString(fmt.Sprintf(`func New%sList() *%s {
	return &%s{}
}`+"\n\n", nameTitled, nameTitled, nameTitled))

	sb.WriteString(fmt.Sprintf("func New%s(\n", nameTitled))
	for _, column := range v.Index {
		sb.WriteString(fmt.Sprintf("\t%s %s,\n", column.Name, column.TypeString(nameTitled)))
	}
	sb.WriteString(fmt.Sprintf("\tparams *%sParams,\n", nameTitled))
	sb.WriteString(fmt.Sprintf(") *%s {\n", nameTitled))
	sb.WriteString(fmt.Sprintf("\treturn &%s{\n", nameTitled))
	for _, column := range v.Index {
		sb.WriteString(fmt.Sprintf("\t\t%s: %s,\n", strings.Title(column.Name), column.Name))
	}
	sb.WriteString("\t\tParams: params,\n")
	sb.WriteString("\t}\n")

	sb.WriteString("}\n\n")

	sb.WriteString(
		fmt.Sprintf(`func (c *%s) Name() string {
	return "%s"
}`+"\n\n",
			nameTitled,
			nameTitled,
		),
	)

	sb.WriteString(
		fmt.Sprintf(`func (c *%s) GetParams() BeanType {
	return c.Params
}`+"\n\n",
			nameTitled,
		),
	)
	sb.WriteString(
		fmt.Sprintf(`func (c *%s) GetParamsType() reflect.Type {
	return reflect.TypeOf(c.Params)
}`+"\n\n",
			nameTitled,
		),
	)
	sb.WriteString(fmt.Sprintf("func (c *%s) Path() string {\n\t", nameTitled))
	sb.WriteString(`path := "/config/" + c.Name()` + "\n")
	sb.WriteString(fmt.Sprintf("if %s {\n\t\treturn path\n\t}\n", strings.Join(conditionalZeroIndex, " &&\n\t\t")))
	sb.WriteString(`return path + "/" + fmt.Sprint(c.` + strings.Join(indexNames, `) + "/" + fmt.Sprint(c.`) + ")\n")
	sb.WriteString("}\n\n")
	return sb.String()
}

func (v *VerbStructure) WriteStructParams() string {
	var sb strings.Builder
	tableTitled := strings.Title(v.Name)
	for _, column := range v.Columns {
		enum := writeEnum(tableTitled, column)
		if enum != "" {
			sb.WriteString(enum + "\n")
		}
	}

	sb.WriteString(fmt.Sprintf("type %sParams struct {\n", tableTitled))
	for _, column := range v.Columns {
		comments := column.WriteComments()
		comments = strings.Replace(strings.TrimSuffix(comments, "\n"), "\n", "\n\t", -1)
		sb.WriteString("\t" + comments + "\n")
		sb.WriteString(fmt.Sprintf("\t%s %s `json:\"%s,omitempty\"`\n", strings.Title(column.Name), column.TypeString(tableTitled), column.Name))
	}
	sb.WriteString("}\n\n")

	sb.WriteString(fmt.Sprintf(`func (p %sParams) iMABean(){}`+"\n", tableTitled))
	return sb.String()

}

func cleanName(name string) string {
	return strings.ReplaceAll(strings.Title(name), "-", "")
}

type EnrichNode struct {
	gosmi.SmiNode
}

func (e *EnrichNode) TypeString(tableName string) string {
	if e.SmiType.BaseType == types.BaseTypeEnum {
		return tableName + strings.Title(e.Name)
	}
	return baseTypeToType(e.SmiType.BaseType)
}

func (e *EnrichNode) WriteComments() string {
	return "// " + strings.Replace(e.Description, "\n", "\n// ", -1) + "\n"

}

func writeEnum(tableName string, node *EnrichNode) string {
	if node.SmiType.BaseType != types.BaseTypeEnum || node.SmiType.Enum == nil {
		return ""
	}

	var sb strings.Builder
	sb.WriteString(
		fmt.Sprintf("type %s%s %s\n", tableName, strings.Title(node.Name), baseTypeToType(node.SmiType.Enum.BaseType)),
	)
	sb.WriteString("const (\n")
	for _, v := range node.SmiType.Enum.Values {
		sb.WriteString(fmt.Sprintf(
			"\t%s%s_%s %s = %d\n",
			tableName,
			strings.Title(node.Name),
			strings.Title(cleanName(v.Name)),
			tableName+strings.Title(node.Name),
			v.Value),
		)
	}
	sb.WriteString(")\n")
	return sb.String()
}

func baseTypeToType(baseType types.BaseType) string {
	switch baseType {
	case types.BaseTypeInteger32:
		return "int32"
	case types.BaseTypeUnsigned32:
		return "uint32"
	case types.BaseTypeFloat32:
		return "float32"
	case types.BaseTypeFloat64:
		return "float64"
	case types.BaseTypeEnum:
		return "enum"
	case types.BaseTypeUnsigned64:
		return "uint64"
	case types.BaseTypeInteger64:
		return "int64"
	default:
		return "string"
	}
}

type arrayStrings []string

var modules arrayStrings
var paths arrayStrings

var prefix string
var debug bool

func (a arrayStrings) String() string {
	return strings.Join(a, ",")
}

func (a *arrayStrings) Set(value string) error {
	*a = append(*a, value)
	return nil
}

func main() {
	flag.BoolVar(&debug, "d", false, "Debug")
	flag.Var(&modules, "m", "Module to load")
	flag.Var(&paths, "p", "Path to add")
	flag.Parse()

	Init()

	oid := flag.Arg(0)
	if oid == "" {
		fmt.Println("Usage: smi [options] <oid>")
		flag.PrintDefaults()
		return
	}
	Subtree(oid)

	Exit()
}

func Init() {
	gosmi.Init()

	for _, path := range paths {
		gosmi.AppendPath(path)
	}

	for i, module := range modules {
		moduleName, err := gosmi.LoadModule(module)
		if err != nil {
			fmt.Printf("Init Error: %s\n", err)
			return
		}
		if debug {
			fmt.Printf("Loaded module %s\n", moduleName)
		}
		modules[i] = moduleName
	}

	if debug {
		path := gosmi.GetPath()
		fmt.Printf("Search path: %s\n", path)
		loadedModules := gosmi.GetLoadedModules()
		fmt.Println("Loaded modules:")
		for _, loadedModule := range loadedModules {
			fmt.Printf("  %s (%s)\n", loadedModule.Name, loadedModule.Path)
		}
	}
}

func Exit() {
	gosmi.Exit()
}

func Subtree(oid string) {
	var node gosmi.SmiNode
	var err error
	if (oid[0] >= '0' && oid[0] <= '9') || oid[0] == '.' {
		node, err = gosmi.GetNodeByOID(types.OidMustFromString(oid))
	} else {
		node, err = gosmi.GetNode(oid)
	}
	if err != nil {
		fmt.Printf("Subtree Error: %s\n", err)
		return
	}

	t := node.AsTable()

	columns := make([]gosmi.SmiNode, 0)
	for _, k := range t.ColumnOrder {
		column, ok := t.Columns[k]
		if !ok {
			continue
		}
		columns = append(columns, column)
	}

	vss := findAllVerbs(node)
	for _, vs := range vss {
		err := vs.WriteFile()
		if err != nil {
			fmt.Printf("Error writing file: %s\n", err)
		}
	}

	cmd := exec.Command("go", "fmt", "github.com/orange-cloudfoundry/goalteon/beans")
	cmd.Dir = ".."
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}

func findAllVerbs(node gosmi.SmiNode) []*VerbStructure {
	oidLenBegin := node.OidLen
	if node.Kind != types.NodeTable {
		rawNode := smi.GetNextNode(node.GetRaw(), types.NodeTable)
		if rawNode == nil {
			return nil
		}
		node = gosmi.CreateNode(rawNode)
	}
	verbs := make([]*VerbStructure, 0)
	prev := ""
	for {
		if node.GetRaw() == nil {
			return verbs
		}
		t := node.AsTable()
		rowNameSplit := make([]string, 0)
		rowNode := smi.GetNextNode(node.GetRaw(), types.NodeRow)
		if rowNode != nil {
			rowNameSplit = splitNameUpperCase(rowNode.Name.String())
		}
		columns := make([]*EnrichNode, 0)
		for _, k := range t.ColumnOrder {
			column, ok := t.Columns[k]
			if !ok {
				continue
			}
			column.Name = columnRename(prev, column.Name, rowNameSplit)
			prev = column.Name
			columns = append(columns, &EnrichNode{SmiNode: column})
		}
		indexes := make([]*EnrichNode, 0)
		for _, indexNode := range t.Index {
			indexes = append(indexes, &EnrichNode{SmiNode: indexNode})
		}
		verbs = append(verbs, &VerbStructure{
			Name:        t.Name,
			Description: t.Description,
			Columns:     columns,
			Index:       indexes,
		})

		rawNode := smi.GetNextNode(node.GetRaw(), types.NodeTable)
		if rawNode == nil {
			return verbs
		}
		if rawNode.OidLen <= oidLenBegin {
			return verbs
		}
		node = gosmi.CreateNode(rawNode)
	}

}

func splitNameUpperCase(srcStrName string) []string {
	newNameArr := make([]string, 0)
	breakIdx := make([]int, 0)
	prevUpper := false
	for i, c := range srcStrName {
		if i == 0 {
			continue
		}
		if c >= 'A' && c <= 'Z' {
			if !prevUpper && i > 1 {
				breakIdx = append(breakIdx, i)
			}
			prevUpper = true
		} else {
			prevUpper = false
		}
	}

	prevIndex := 0
	for _, i := range breakIdx {
		newNameArr = append(newNameArr, srcStrName[prevIndex:i])
		prevIndex = i
	}
	newNameArr = append(newNameArr, srcStrName[prevIndex:])
	return newNameArr
}

func columnRename(prevName string, name string, tableNameSplit []string) string {
	curPrefix := ""
	curPropArr := splitNameUpperCase(name)
	curWordLen := 0
	for i, curWord := range tableNameSplit {
		if i == len(tableNameSplit) {
			break
		}
		if strings.ToLower(curWord) == strings.ToLower(curPropArr[i]) {
			curPrefix += curPropArr[i]
			curWordLen = len(curPropArr[i])
		} else if strings.ToLower(curWord) == strings.ToLower(curPropArr[i]+"s") {
			curPrefix += curPropArr[i]
			curWordLen = len(curPropArr[i])
			break
		} else {
			break
		}
	}

	if prevName == name[len(curPrefix):] {
		name = name[len(curPrefix)-curWordLen:]
	} else {
		name = name[len(curPrefix):]
	}
	switch name {
	case "String":
		return "StringVal"
	}
	return name
}
