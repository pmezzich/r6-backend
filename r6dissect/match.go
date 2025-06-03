package r6dissect

import (
    "encoding/json"
    "fmt"
    "os"
)

type Round struct {
    Map      string `json:"map"`
    Operator string `json:"operator"`
    Result   string `json:"result"`
}

type MatchReader struct {
    paths     []string
    rounds    []*Round
    queries   []string
    listeners [][]func(Round)
}

func NewMatchReader(paths []string) *MatchReader {
    return &MatchReader{
        paths:     paths,
        rounds:    make([]*Round, len(paths)),
        listeners: make([][]func(Round), len(paths)),
    }
}

func (m *MatchReader) ReadAll() error {
    for i := range m.paths {
        if err := m.read(i); err != nil {
            return err
        }
    }
    return nil
}

func (m *MatchReader) read(i int) error {
    if i < 0 || i >= len(m.paths) {
        return ErrInvalidFile
    }
    if m.rounds[i] != nil {
        return nil
    }
    f, err := os.Open(m.paths[i])
    if err != nil {
        return err
    }
    defer f.Close()

    var r Round
    if err := json.NewDecoder(f).Decode(&r); err != nil {
        return err
    }
    m.rounds[i] = &r
    for _, listener := range m.listeners[i] {
        listener(r)
    }
    return nil
}

func (m *MatchReader) Listen(i int, query string, listener func(Round)) {
    m.queries = append(m.queries, query)
    m.listeners[i] = append(m.listeners[i], listener)
}

func (m *MatchReader) PrintRounds() {
    for i, r := range m.rounds {
        if r != nil {
            fmt.Printf("Round %d: %+v\n", i, *r)
        }
    }
}
