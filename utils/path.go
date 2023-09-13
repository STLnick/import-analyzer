package utils

import "strings"

func removeChars(src string, remove string) string {
    var trimmed string
    var char string
    for _, c := range src {
        char = string(c)
        if !strings.Contains(remove, char) {
            trimmed += char
        }
    }
    return trimmed
}

func ExtractPathParts(s string) []string {
    var path string
    tokens := strings.Fields(s)

    for i, token := range tokens {
        if token == "from" {
            path = tokens[i+1]
        }
    }

    path = removeChars(path, "'\";")
    return strings.Split(path, "/")
}

