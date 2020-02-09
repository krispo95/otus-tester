package tester

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

type TestTask interface {
	Run([]string) []string
}

func RunTest(path string, task TestTask) {
	nr := 0
	for {
		inStrings, err := ReadFile(fmt.Sprintf("%s/test.%d.in", path, nr))
		if err != nil {
			break
		}
		outStrings, err := ReadFile(fmt.Sprintf("%s/test.%d.out", path, nr))
		if err != nil {
			break
		}
		nr++
		result := task.Run(inStrings)
		if len(result) != len(outStrings) {
			fmt.Printf("Тест %+v - false\nlengths not equal, expected len: %d, received len: %d\n",
				nr, len(outStrings), len(result))
			return
		}
		for i := range result {
			if result[i] != outStrings[i] {
				fmt.Printf("Тест %+v - false, results not equal, expected: %s, received: %s\n",
					nr, outStrings[i], result[i])
				return
			}
		}
		fmt.Printf("Тест %+v - true\n", nr)
	}
}

func ReadFile(path string) ([]string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	linesByte := bytes.Split(data, []byte("\n"))
	linesString := make([]string, len(linesByte))
	for i, item := range linesByte {
		str := string(bytes.Trim(item, "\r"))
		linesString[i] = str
	}
	return linesString, nil

}
