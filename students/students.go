package students

import (
    "fmt"
    "time"
    "io/ioutil"
    "math/rand"
    "encoding/json"
)

type Students struct {
    List []string `json:"list"`
}

func New(path string) (*Students, error) {
    file, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }

    s := &Students{}

    err = json.Unmarshal(file, s)
    if err != nil {
        return nil, err
    }
    
    return s, nil
}

func (s *Students) Shuffle() []string {
    list := s.List

    r := rand.New(rand.NewSource(time.Now().Unix()))
    for n := len(list); n > 0; n-- {
        randIndex := r.Intn(n)
        list[n-1], list[randIndex] = list[randIndex], list[n-1]
    }

    return list
}

func (s *Students) ShuffleString() string {
    list := s.Shuffle()

    var str string
    for i, v := range(list) {
        str = fmt.Sprintf("%s%d. %s\n", str, i+1, v);
    }

    return str
}