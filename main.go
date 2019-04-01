package main

import (
	"fmt"
	"github.com/jszwec/csvutil"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// main
// csv -> output csv
func main() {
	if len(os.Args) < 2 {
		panic("ファイルを指定してください")
	}

	// select file
	input, err := read(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%+v\n", input)

	// output struct
	out := inputToOutput(input)

	// ファイル出力
	outputPath := outputPath(os.Args[1])
	err = output(out, outputPath)
	if err != nil {
		log.Fatalln(err)
	}
}

// read
func read(path string) ([]Input, error) {
	records := []Input{}

	b, _ := ioutil.ReadFile(path)
	if err := csvutil.Unmarshal(b, &records); err != nil {
		return nil, err
	}

	return records, nil
}

// inputToOutput
func inputToOutput(param []Input) []Output {
	ret := []Output{}
	for _, v := range param {
		t := Output{
			ID:   v.ID,
			Name: v.Name,
			Cd:   v.Cd,
			Test: v.Test,
		}
		ret = append(ret, t)
	}
	return ret
}

// output
func output(output []Output, path string) error {
	b, _ := csvutil.Marshal(output)
	_ = ioutil.WriteFile(path, b, os.ModePerm)

	return nil
}

// outputPath
func outputPath(path string) string {
	_, fileName := filepath.Split(path)
	fileBase := strings.Split(fileName, filepath.Ext(path))

	// write file path
	return "./" + fileBase[0] + "_out" + filepath.Ext(path)
}
