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

package ini_test

import (
    "github.com/cuberat/go-ini/ini"
    "reflect"
    "testing"
)

func TestLoad(t *testing.T) {
    buf := "foo=1\n\n; this is a comment\n[test]\nfoo=1\nbar=2\n[section1]\nfoo=bar"
    conf, err := ini.LoadFromString(buf)
    if err != nil {
        t.Error("LoadFromString() failed")
    }

    expecting := map[string]map[string]string{
        "default": map[string]string{"foo":"1"},
        "test": map[string]string{"foo":"1", "bar":"2"},
        "section1": map[string]string{"foo":"bar"},
    }

    if !reflect.DeepEqual(conf, expecting) {
        t.Errorf("didn't get expected values from structured conf\n\tgot %v\n\texpected %v",
            conf, expecting)
    }


    // t.Errorf("conf=%v", conf)
}

func TestFlatLoad(t *testing.T) {
    buf := "foo=1\n\n; this is a comment\n[test]\nfoo=1\nbar=2\n[section1]\nfoo=bar\n"

    flat_conf, err := ini.LoadFromStringFlat(buf)
    if err != nil {
        t.Error("LoadFromStringFlat() failed")
    }

    flat_expecting := map[string]string{
        "default.foo":"1",
        "test.foo":"1",
        "test.bar":"2",
        "section1.foo":"bar",
    }

    if !reflect.DeepEqual(flat_conf, flat_expecting) {
        t.Errorf("didn't get expected values from flat conf\n\tgot %v\n\texpected %v",
            flat_conf, flat_expecting)
    }

}
