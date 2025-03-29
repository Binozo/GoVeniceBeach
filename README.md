# GoVeniceBeach

## Install
```sh
$ go get -u github.com/binozo/govenicebeach
```

## Example

```go
package main

import (
	"github.com/binozo/govenicebeach/pkg/venicebeach"
	"fmt"
)

func main() {
	/* 
	I once had a dream where HTTP Toolkit and an Android Emulator were my best friends
	Then a Bird whispered me "CLIENT_SECRET"
	 */
	// venicebeach.Login("<EMAIL>", "<PASSWORD>", "<CLIENT_SECRET>")
	session := venicebeach.NewSession("<my_token>")

	studio, err := session.GetStudioInfo(14)
	if err != nil {
		panic(err)
	}
	fmt.Println(studio.CheckedInUsers.Current)
}

```