package r6dissect

import (
    "encoding/json"
    "errors"
    "os"
    "sync"

    "github.com/xuri/excelize/v2"
)

var ErrInvalidFile = errors.New("invalid file index")

type MatchReader struct {
    paths     []string
    rounds    []*Round
    queries   []string
    listeners [][]Listener
    mutex     sync.Mutex
}

func NewMatchReader(paths []string) *MatchReader {
    return &MatchReader{
        paths:     paths,
        rounds:    make([]*Round, len(paths)),
        listeners: make([][]Listener, len(paths)),
    }
}

func (m *MatchReader) AddListener(query string, listener Listener) {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    m.queries = append(m.queries, query)
    for i := range m.paths {
        m.listeners[i] = append(m.listeners[i], listener)
    }
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

    r, err := NewReader(f)
    if err != nil {
        return err
    }
    m.rounds[i] = r

    for i := 0; i < len(m.queries); i++ {
        for _, listener := range m.listeners[i] {
            r.Listen(m.queries[i], listener)
        }
    }
    return r.Read()
}

func (m *MatchReader) ToJSON(path string) error {
    data, err := json.MarshalIndent(m.rounds, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(path, data, 0644)
}

func (m *MatchReader) ToXLSX(path string) error {
    f := excelize.NewFile()
    _, err := f.NewSheet("Match")
    if err != nil {
        return err
    }

    err = f.DeleteSheet("Sheet1")
    if err != nil {
        return err
    }

    for i, round := range m.rounds {
        if round != nil {
            round.ToXLSX(f, i)
        }
    }

    return f.SaveAs(path)
}