package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

func isMdFile(fileName []byte) bool {
	return bytes.HasSuffix(fileName, []byte(".md")) || bytes.HasSuffix(fileName, []byte(".MD"))
}

func getBaseDir(fullPath string) string {
	return strings.Split(fullPath, string(os.PathSeparator))[0]
}

func isBlackListed(fileName []byte) bool {
	ignoredDirs := [3]string{".git", "_repoUtils", ".github"}

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

func prettyPrintResults(data *map[string]int) {
	json, err := json.MarshalIndent(data, "", "   ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nAdditions from last %v to today.\n", time.Now().AddDate(0, 0, -30).Format("2th Jan"))
	fmt.Println("")
	fmt.Println(string(json))
}

func generateFile(dirStats *map[string]int) {
	type Entry struct {
		Topic string
		Count int
	}

	entries := make([]Entry, 0, len(*dirStats))

	for k, v := range *dirStats {
		topic := strings.ReplaceAll(strings.Title(k), "-", " ")
		entry := Entry{Topic: topic, Count: v}
		entries = append(entries, entry)
	}
	sort.SliceStable(entries, func(i, j int) bool {
		return entries[i].Count > entries[j].Count
	})

	fmt.Printf("%v\n", entries)

	// Dump all stats to the template file
	f, _ := os.Create("_repoUtils/public/index.html")
	defer f.Close()

	t, err := template.ParseFiles("_repoUtils/tpl.tpl")
	if err != nil {
		log.Fatalf("error parsing template %s", err)
	}
	err = t.Execute(f, entries)
	if err != nil {
		log.Fatalf("error parsing template %s", err)
	}

	log.Output(1, "Index file generated")
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
	dirStats := make(map[string]int)

	for scanner.Scan() {
		filePath := scanner.Bytes()
		isMd := isMdFile(filePath)
		if isMd == true && !isBlackListed(filePath) {
			filePaths = append(filePaths, filePath)
			path := string(filePath)
			log.Output(1, fmt.Sprintf("Analysing %v\n", path))
			// store stats for each top level dir
			dirName := getBaseDir(path)
			_, exists := dirStats[dirName]
			if exists == false {
				dirStats[dirName] = 0
			}
			cmdStat := exec.Command("git", "diff", "@{30.days.ago}", "--numstat", "--", path)
			out, err := cmdStat.CombinedOutput()
			if err != nil {
				msg, _ := fmt.Printf("Git command failed for file %s", path)
				log.Fatal(msg)
			}
			statsLine := bytes.Fields(out)
			if len(statsLine) > 0 {
				additions, _ := strconv.Atoi(string(statsLine[0]))
				dirStats[dirName] = dirStats[dirName] + additions
			}
		}
	}

	prettyPrintResults(&dirStats)
	generateFile(&dirStats)
}