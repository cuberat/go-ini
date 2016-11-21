Package ini provides INI file read functionality in Go.

## Installation

To get the latest changes

    go get github.com/cuberat/go-ini

## Getting started

File:

```ini
foo=bar
[db]
user=myuser
password=mypassword
```

Code:

```go
import (
    "fmt"
    "github.com/cuberat/go-ini/ini"
)

conf, err := ini.LoadFromFile("my/file/path.conf")
fmt.Printf("%v\n", conf)
```

Output:

```bash
map[default:map[foo:bar] db:map[user:myuser password:mypassword]]
```
