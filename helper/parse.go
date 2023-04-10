package helper

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

func Base64Parse(t string) string {
	a := t[:100] + t[132:]
	return a
}

func ImageParse(s string, chat_id string) {
	// remove the prefix
	s = strings.TrimPrefix(s, "data:image/png;base64,")
	// decode base64 string
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// create a file
	file, err := os.Create("./uploads/" + chat_id + ".png")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	// write the data to the file
	file.Write(decoded)
	fmt.Println("File created: image.png")
}