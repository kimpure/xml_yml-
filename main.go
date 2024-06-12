package main

import (
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

type Node struct {
    XMLName xml.Name
    Content []byte `xml:",innerxml"`
    Nodes   []Node `xml:",any"`
}

func changeYml(element Node, indent int) string {
    var returnVar string
    indentStr := strings.Repeat("    ", indent)

    if returnVar == "" {
        returnVar = fmt.Sprintf("%s:\n%s", element.XMLName.Local, "    ")
    } else {
        returnVar = fmt.Sprintf("%s\n%s%s:", returnVar, indentStr, element.XMLName.Local)
    }

    for _, key := range element.Nodes {
        if len(key.Nodes) > 0 {
            returnVar += changeYml(key, indent+1)
        } else {
            returnVar = fmt.Sprintf("%s\n%s    %s: %s", returnVar, indentStr, key.XMLName.Local, strings.TrimSpace(string(key.Content)))
        }
    }

    return returnVar
}

func main() {
    // XML 파일 읽기
    xmlFile, err := os.Open("comp.xml")
    if err != nil {
        fmt.Println("Error opening XML file:", err)
        return
    }
    defer xmlFile.Close()

    byteValue, _ := ioutil.ReadAll(xmlFile)

    var root Node
    xml.Unmarshal(byteValue, &root)

    // YAML 형식으로 변환
    ymlContent := changeYml(root, 0)

    // YAML 파일로 저장
    ymlFilePath := "comp.yml"
    err = ioutil.WriteFile(ymlFilePath, []byte(ymlContent), 0644)
    if err != nil {
        fmt.Println("Error writing YAML file:", err)
        return
    }
}
