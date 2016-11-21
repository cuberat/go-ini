package ini_test

import (
    "github.com/cuberat/go-ini/ini"
    "reflect"
    "testing"
)

func TestLoad(t *testing.T) {
    buf := "foo=1\n\n; this is a comment\n[test]\nfoo=1\nbar=2\n[section1]\nfoo=bar\n"
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
        t.Errorf("didn't get expected values from conf\n\tgot %v\n\texpected %v",
            conf, expecting)
    }

    // t.Errorf("conf=%v", conf)
}
