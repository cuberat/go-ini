// Copyright (c) 2016 Don Owens <don@regexguy.com>.  All rights reserved.
//
// This software is released under the BSD license:
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
//
//  * Redistributions of source code must retain the above copyright
//    notice, this list of conditions and the following disclaimer.
//
//  * Redistributions in binary form must reproduce the above
//    copyright notice, this list of conditions and the following
//    disclaimer in the documentation and/or other materials provided
//    with the distribution.
//
//  * Neither the name of the author nor the names of its
//    contributors may be used to endorse or promote products derived
//    from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS
// FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE
// COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT,
// INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
// HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT,
// STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED
// OF THE POSSIBILITY OF SUCH DAMAGE.

package ini

import (
    "bufio"
    "io"
    // "log"
    "os"
    "strings"
)

type my_ini struct {
    foo string
}

func LoadFromFile(filename string) (map[string]map[string]string, error) {
    o := new(my_ini)
    return o.load_from_file(filename)
}

func LoadFromReader(r io.Reader) (map[string]map[string]string, error) {
    o := new(my_ini)
    return o.load_from_reader(r)
}

func LoadFromString(buf string) (map[string]map[string]string, error) {
    o := new(my_ini)
    return o.load_from_reader(strings.NewReader(buf))
}

func LoadFromFileFlat(filename string) (map[string]string, error) {
    o := new(my_ini)

    conf, err := o.load_from_file(filename)
    if err != nil {
        return nil, err
    }

    flat := o.structured_to_flat(conf, ".")
    return flat, nil
}

func LoadFromReaderFlat(r io.Reader) (map[string]string, error) {
    o := new(my_ini)
    conf, err := o.load_from_reader(r)
    if err != nil {
        return nil, err
    }

    flat := o.structured_to_flat(conf, ".")
    return flat, nil

}

func LoadFromStringFlat(buf string) (map[string]string, error) {
    o := new(my_ini)
    conf, err := o.load_from_reader(strings.NewReader(buf))
    if err != nil {
        return nil, err
    }

    flat := o.structured_to_flat(conf, ".")
    return flat, nil
}

func (o *my_ini) load_from_file(filename string) (map[string]map[string]string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return nil, nil
    }
    defer f.Close()

    return o.load_from_reader(f)
}

func (o *my_ini) load_from_reader(r io.Reader) (map[string]map[string]string, error) {
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
                    break
                }
            }

            // FIXME: if b[0] != ']' -- return error

        } else {
            parts := strings.SplitN(line, "=", 2)
            k, v := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
            cur_sect[k] = v
        }
        
    }
    
    return conf, nil
}

func (o *my_ini) structured_to_flat(conf_in map[string]map[string]string, del string) map[string]string {
    conf := make(map[string]string)
    for sect_name, sect := range conf_in {
        for k, v := range sect {
            name := sect_name + del + k
            conf[name] = v
        }
    }

    return conf
}
