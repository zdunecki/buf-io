package file

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const fileNameBuffer = 64
const fileSizeBuffer = 10

type File struct {
	Name []byte
	Size []byte
	Content []byte
}

func CreateFormat(path string) []File {
	fi, err := os.Stat(path)

	if err != nil {
		panic(err)
		return nil
	}

	var result []File

	switch mode := fi.Mode(); {
	case mode.IsDir():

		files, err := ioutil.ReadDir(path)
		if err != nil {
			panic(err)
		}

		for _, f := range files {
			if f.IsDir() {
				continue
			}

			fileName, fileSize, content, err := createFormat(path + "/" + f.Name())
			if err != nil {
				panic(err)
			}

			result = append(result, File{
				fileName,
				fileSize,
				content,
			})
		}

	case mode.IsRegular():
		fileName, fileSize, content, err := createFormat(path)
		if err != nil {
			panic(err)
		}

		result = append(result, File{
			fileName,
			fileSize,
			content,
		})
	}

	return result
}

func createFormat(path string) ([]byte, []byte, []byte, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil, nil, nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, nil, nil, err
	}

	fileName := fillString(fileInfo.Name(), fileNameBuffer)
	fileSize := fillString(strconv.Itoa(int(fileInfo.Size())), fileSizeBuffer)

	reader := bufio.NewReader(file)
	content, _ := ioutil.ReadAll(reader)

	return []byte(fileName), []byte(fileSize), content, nil
}

func fillString(returnString string, toLength int) string {
	for {
		lengthString := len(returnString)
		if lengthString < toLength {
			returnString += ":"
			continue
		}
		break
	}
	return returnString
}

func GetName(message []byte) string {
	return strings.Split(string(message), ":")[0]
}

func GetSize(message []byte) (int, error) {
	return strconv.Atoi(strings.Split(string(message), ":")[0])
}
