package ini

import (
    "bufio"
    "io"
    // "log"
    "os"
    "strings"
)

func LoadFromFile(filename string) (map[string]map[string]string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return nil, nil
    }
    defer f.Close()

    // return LoadFromReader(bufio.NewReader(f))
    return LoadFromReader(f)
}

func LoadFromString(buf string) (map[string]map[string]string, error) {
    return LoadFromReader(strings.NewReader(buf))
}

func LoadFromReader(r io.Reader) (map[string]map[string]string, error) {
    in := bufio.NewReader(r)
    conf := make(map[string]map[string]string)
    cur_sect_name := "default"

    cur_sect := make(map[string]string)
    conf[cur_sect_name] = cur_sect

    for line, err := in.ReadString('\n'); err == nil; line, err = in.ReadString('\n') {
        line = strings.TrimSpace(line)
        if len(line) == 0 {
            continue // empty line
        }

        sr := strings.NewReader(line)
        b := []byte{'f'}
        n, err := sr.Read(b)
        if err != nil || n == 0 {
            return conf, err
        }

        c := b[0]

        if c == ';' || c == '#' {
            continue // comment
        }

        if c == '[' {
            i := 0
            for n, err = sr.Read(b); err == nil; n, err = sr.Read(b) {
                i += 1
                if b[0] == ']' {
                    bytes := []byte(line)
                    cur_sect_name = string(bytes[1:i])
                    cur_sect = make(map[string]string)
                    conf[cur_sect_name] = cur_sect
                }
            }
        } else {
            parts := strings.SplitN(line, "=", 2)
            k, v := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
            cur_sect[k] = v
        }
        
    }
    

    // conf["default"] = map[string]string{"foo":"bar","key1":"val1"}

    return conf, nil
}
