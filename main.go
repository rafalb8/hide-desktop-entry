package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var HomeDir = func() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return home
}()

// Paths
var (
	SystemAppPath = "/usr/share/applications"
	LocalAppPath  = HomeDir + "/.local/share/applications"
)

const hiddenFileContents = "[Desktop Entry]\nHidden=true\n"

func appList(path string) []string {
	list, _ := filepath.Glob(filepath.Join(path, "*.desktop"))
	for i := range list {
		list[i] = filepath.Base(list[i])
	}
	return list
}

func hide(list []string) {
	for _, file := range list {
		path := filepath.Join(LocalAppPath, file)
		err := os.WriteFile(path, []byte(hiddenFileContents), 0644)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	list := listDifference(appList(SystemAppPath), appList(LocalAppPath))
	fmt.Println("========================")
	width := fmt.Sprint(len(fmt.Sprint(len(list))))
	for i, app := range list {
		fmt.Printf("%"+width+"d: %s\n", i, app)
	}
	fmt.Println("========================")
	fmt.Println("==> Select entries to hide (eg: 1 2 3 or 1-3)")
	fmt.Print("==> ")
	toHide := RangesSlice(ReadRanges(), list)

	fmt.Println("========================")
	fmt.Println("Entries to be hidden:")
	for _, v := range toHide {
		fmt.Println("*", v)
	}
	fmt.Println("========================")
	fmt.Print("==> Are you sure? (y/N): ")
	var confirm string
	fmt.Scanf("%s", &confirm)
	if strings.ToLower(confirm) == "y" {
		hide(toHide)
	}
}
