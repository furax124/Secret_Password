package main

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	password := readPassword()
	key := generateRandomKey()
	content := readTemplateFile(filepath.Join("..", "main.go"))
	modifiedContent := replacePlaceholders(content, password, key)
	writeToFile(filepath.Join("..", "main.go"), modifiedContent)
	installDependencies()
	compileProject()
	cleanUp()
}

func readPassword() string {
	fmt.Print("Enter your password: ")
	reader := bufio.NewReader(os.Stdin)
	password, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading password:", err)
	}
	return strings.TrimSpace(password)
}

func generateRandomKey() string {
	key := make([]byte, 16)
	if _, err := rand.Read(key); err != nil {
		log.Fatal("Error generating random key:", err)
	}
	return hex.EncodeToString(key)
}

func readTemplateFile(path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Error reading template file:", err)
	}
	return string(content)
}

func replacePlaceholders(content, password, key string) string {
	content = strings.Replace(content, "%PHRASE%", password, -1)
	return strings.Replace(content, "%KEY%", key, -1)
}

func writeToFile(path, content string) {
	if err := ioutil.WriteFile(path, []byte(content), 0644); err != nil {
		log.Fatal("Error writing output file:", err)
	}
	fmt.Println("File successfully written to", path)
}

func installDependencies() {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Error installing dependencies:", err)
	}
}

func compileProject() {
	cmd := exec.Command("garble", "-literals", "-seed=random", "-tiny", "build", "-ldflags=-w -s -H=windowsgui -buildid=", "-trimpath")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Error compiling project:", err)
	}
}

func cleanUp() {
	templatePath := filepath.Join("..", "Template.txt")
	mainPath := filepath.Join("..", "main.go")
	content, err := ioutil.ReadFile(templatePath)
	if err != nil {
		log.Fatal("Error reading template file:", err)
	}
	if err := ioutil.WriteFile(mainPath, content, 0644); err != nil {
		log.Fatal("Error writing main.go file:", err)
	}
	fmt.Println("main.go successfully reset to template.")
}
