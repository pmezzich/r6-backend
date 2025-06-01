package r6dissect

import (
    "os"
    "r6dissect/reader"
    "r6dissect/match"
)

func ParseMatchReplay(filePath string) (*match.Match, error) {
    f, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    parsedReader, err := reader.ReadReplay(f)
    if err != nil {
        return nil, err
    }

    parsedMatch, err := match.Parse(parsedReader)
    if err != nil {
        return nil, err
    }

    return parsedMatch, nil
}
