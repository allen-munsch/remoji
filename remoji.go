package main

import (
    "bufio"
    "flag"
    "fmt"
    "io"
    "os"
    "regexp"
)

var inPlace = flag.Bool("i", false, "edit file in place")

// Emoji + variation selector remover
var emojiRegex = regexp.MustCompile(
    `[\x{1F300}-\x{1F5FF}` + // Misc symbols & pictographs
        `\x{1F600}-\x{1F64F}` + // Emoticons
        `\x{1F680}-\x{1F6FF}` + // Transport & map symbols
        `\x{2600}-\x{26FF}` +   // Misc symbols
        `\x{2700}-\x{27BF}` +   // Dingbats
        `\x{1F1E6}-\x{1F1FF}` + // Flags
        `\x{1F900}-\x{1F9FF}` + // Supplemental symbols
        `\x{1FA70}-\x{1FAFF}` + // Extended-A
        `\x{FE0F}` +            // Variation selector-16
        `\x{FE0E}` +            // Variation selector-15
        `]`)

func stripEmoji(s string) string {
    return emojiRegex.ReplaceAllString(s, "")
}

func processReader(r io.Reader, w io.Writer) error {
    scanner := bufio.NewScanner(r)
    for scanner.Scan() {
        fmt.Fprintln(w, stripEmoji(scanner.Text()))
    }
    return scanner.Err()
}

func main() {
    flag.Parse()

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

