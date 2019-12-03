package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("git", "log", "--pretty=format:", "--name-only", "--since='30 days ago'", "--stat")
	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatal("Git command failed")
	}
	r := bytes.NewReader(out)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	//	fmt.Printf("%v\n", string(out))
}
