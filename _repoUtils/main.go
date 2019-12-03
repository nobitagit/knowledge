package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func isMdFile(fileName []byte) bool {
	return bytes.HasSuffix(fileName, []byte(".md")) || bytes.HasSuffix(fileName, []byte(".MD"))
}

func getBaseDir(fullPath string) string {
	return strings.Split(fullPath, string(os.PathSeparator))[0]
}

func isBlackListed(fileName []byte) bool {
	ignoredDirs := [2]string{".git", "_repoUtils"}

	dir := filepath.Dir(string(fileName))

	// top level files are always ignored, so we return early
	if dir == "." {
		return true
	}
	// derive the top level directory from the path
	baseDir := getBaseDir(dir)
	var contained bool
	for _, ignoredDir := range ignoredDirs {
		contained = ignoredDir == baseDir
	}

	return contained
}

func main() {
	cmd := exec.Command("git", "log", "--pretty=format:", "--name-only", "--since='30 days ago'", "--stat")
	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatal("Git command failed")
	}
	r := bytes.NewReader(out)
	scanner := bufio.NewScanner(r)

	var filePaths [][]byte
	var dirStats map[string]int

	for scanner.Scan() {
		filePath := scanner.Bytes()
		isMd := isMdFile(filePath)
		if isMd == true && !isBlackListed(filePath) {
			filePaths = append(filePaths, filePath)
			k := string(filePath)
			dirStats[k] = 0
			cmd2 := exec.Command("git", "log", "--pretty=format:", "--name-only", "--since='30 days ago'", "--stat")
			fmt.Printf("%v contains md\n", scanner.Text())
		}
	}

	fmt.Printf("%s\n", filePaths)
}
