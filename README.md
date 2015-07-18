#libcaesar
Caesar cipher library

## What is this
This is a library to perform [Caesar cipher](https://en.wikipedia.org/wiki/Caesar_cipher) substitutions on the basic latin characters a-z and A-Z, all other characters will be ignored.

## Usage

```shell
go get github.com/billmcmath/libcaesar
```

```go
package main

import (
    "fmt"
    libcaesar "github.com/billmcmath/libcaesar"
)

func main() {
    source := "THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG, the quick brown fox jumps over the lazy dog."
    key := int8(10)

    result, _ := libcaesar.Encrypt(key, source)
    plain, _ := libcaesar.Decrypt(key, result)
    fmt.Printf("The plain text is \"%s\"\n", source)
    fmt.Printf("The cipher text is \"%s\"\n", result)
    fmt.Printf("The decrypted text is \"%s\"\n", plain)
}
```

## License
This library is licensed under the [MIT license](https://github.com/billmcmath/libcaesar/blob/master/LICENSE)