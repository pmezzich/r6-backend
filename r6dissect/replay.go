package r6dissect

import (
    "os"
)

func ParseMatchReplay(filePath string) (*Match, error) {
    f, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    parsedReader, err := ReadReplay(f)
    if err != nil {
        return nil, err
    }

    parsedMatch, err := Parse(parsedReader)
    if err != nil {
        return nil, err
    }

    return parsedMatch, nil
}
