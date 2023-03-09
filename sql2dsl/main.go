package main

// TODO: 有问题, sdk的问题， 代码没啥问题

import (
	"fmt"
	"github.com/cch123/elasticsql"
	"github.com/jysui123/esql"
	"io/ioutil"
	"os"
	"path/filepath"
)

const dirPath = "sql2dsl"

func main() {

	curPath, e := os.Getwd()
	//curPath, e := filepath.Abs(filepath.Dir(os.Args[0]))
	if e != nil {
		panic(e)
	}
	fmt.Println(curPath)

	sql := readSQL(curPath)
	{
		dsl, esType, e := elasticsql.ConvertPretty(sql)
		if e != nil {
			fmt.Println(e.Error())
			panic(e)
		}
		fmt.Println(dsl)
		fmt.Println(esType)

		writeJSON(curPath, "./output1.json", dsl)
	}
	{
		ne := esql.NewESql()
		dsl, i, e := ne.ConvertPretty(sql)
		if e != nil {
			fmt.Println(e.Error())
			panic(e)
		}
		fmt.Println(i)
		fmt.Println(dsl)
		writeJSON(curPath, "./output2.json", dsl)

	}
}

func readSQL(curPath string) string {
	file, e := os.Open(filepath.Join(curPath, dirPath, "./input.sql"))
	defer file.Close()
	if e != nil {
		panic(e)
	}
	content, e := ioutil.ReadAll(file)
	if e != nil {
		panic(e)
	}
	result := string(content)
	fmt.Println(result)
	return result
}

func writeJSON(curPath, fileName, content string) {
	e := os.WriteFile(filepath.Join(curPath, dirPath, fileName), []byte(content), 0644)
	if e != nil {
		panic(e)
	}

}
