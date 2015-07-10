# tagtype

Type validation for Golang structs.
It leverages the struct tags and reflect to do so.

# Example

```go
package main

import tagtype "github.com/creack/tagtype"

type Example struct {
	ID interface{} `tt:"int,[]byte"`
}

func main() {
	ret := tagtype.Validate(Example{
		ID: []byte("1234"),
	})
	println(ret)
}
```

# Note

- If the `tt` tag is not present or if the value is nil, the field is skipped.
- Made for fun after the GopherCon 2015 talk from @3rf and @shelman. You should not use that package.
