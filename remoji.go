package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	inPlace    = flag.Bool("i", false, "edit file in place")
	grepMode   = flag.Bool("grep", false, "recursively find files containing emojis")
	showDetail = flag.Bool("o", false, "with --grep, show line numbers and emojis found")
)

// Emoji + variation selector remover
var emojiRegex = regexp.MustCompile(
	`[\x{1F300}-\x{1F5FF}` + // Misc symbols & pictographs
		`\x{1F600}-\x{1F64F}` + // Emoticons
		`\x{1F680}-\x{1F6FF}` + // Transport & map symbols
		`\x{2600}-\x{26FF}` + // Misc symbols
		`\x{2700}-\x{27BF}` + // Dingbats
		`\x{1F1E6}-\x{1F1FF}` + // Flags
		`\x{1F900}-\x{1F9FF}` + // Supplemental symbols
		`\x{1FA70}-\x{1FAFF}` + // Extended-A
		`\x{FE0F}` + // Variation selector-16
		`\x{FE0E}` + // Variation selector-15
		`]`)

func stripEmoji(s string) string {
	return emojiRegex.ReplaceAllString(s, "")
}

func containsEmoji(s string) bool {
	return emojiRegex.MatchString(s)
}

func findEmojis(s string) []string {
	return emojiRegex.FindAllString(s, -1)
}

func uniqueEmojis(emojis []string) []string {
	seen := make(map[string]bool)
	var result []string
	for _, e := range emojis {
		// Skip variation selectors on their own
		if e == "\uFE0F" || e == "\uFE0E" {
			continue
		}
		if !seen[e] {
			seen[e] = true
			result = append(result, e)
		}
	}
	return result
}

func processReader(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Fprintln(w, stripEmoji(scanner.Text()))
	}
	return scanner.Err()
}

func grepEmojis(root string, showDetail bool) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories and hidden files/folders
		if info.IsDir() {
			if strings.HasPrefix(info.Name(), ".") && path != root {
				return filepath.SkipDir
			}
			return nil
		}

		// Skip hidden files
		if strings.HasPrefix(info.Name(), ".") {
			return nil
		}

		// Skip binary files (simple heuristic: check extension)
		ext := strings.ToLower(filepath.Ext(path))
		binaryExts := map[string]bool{
			".png": true, ".jpg": true, ".jpeg": true, ".gif": true,
			".ico": true, ".webp": true, ".svg": true,
			".pdf": true, ".zip": true, ".tar": true, ".gz": true,
			".exe": true, ".bin": true, ".so": true, ".dylib": true,
			".woff": true, ".woff2": true, ".ttf": true, ".eot": true,
			".mp3": true, ".mp4": true, ".wav": true, ".avi": true,
			".db": true, ".sqlite": true,
		}
		if binaryExts[ext] {
			return nil
		}

		// Read file
		data, err := os.ReadFile(path)
		if err != nil {
			return nil
		}

		content := string(data)
		if !containsEmoji(content) {
			return nil
		}

		if !showDetail {
			fmt.Println(path)
			return nil
		}

		// Show detailed output with line numbers
		scanner := bufio.NewScanner(strings.NewReader(content))
		lineNum := 0
		var matches []string

		for scanner.Scan() {
			lineNum++
			line := scanner.Text()
			emojis := findEmojis(line)
			if len(emojis) > 0 {
				unique := uniqueEmojis(emojis)
				matches = append(matches, fmt.Sprintf("%d:%s", lineNum, strings.Join(unique, "")))
			}
		}

		if len(matches) > 0 {
			fmt.Printf("%s  %s\n", path, strings.Join(matches, " "))
		}

		return nil
	})
}

func main() {
	flag.Parse()

	if *grepMode {
		root := "."
		if len(flag.Args()) > 0 {
			root = flag.Args()[0]
		}
		if err := grepEmojis(root, *showDetail); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	}

	if len(flag.Args()) == 0 {
		// stdin → stdout
		processReader(os.Stdin, os.Stdout)
		return
	}

	path := flag.Args()[0]

	if *inPlace {
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		cleaned := stripEmoji(string(data))
		if err := os.WriteFile(path, []byte(cleaned), 0644); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	}

	// file → stdout
	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()
	processReader(f, os.Stdout)
}
