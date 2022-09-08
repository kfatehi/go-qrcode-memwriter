This module wraps [go-qrcode](github.com/yeqown/go-qrcode)'s [standard writer](https://github.com/yeqown/go-qrcode/tree/main/writer/standard) replacing the call to `os` with [afero](https://github.com/spf13/afero#memory-backed-storage) (A FileSystem Abstraction System for Go) for those of us that don't need to write to an actual file.

Example

qr.go

```go
package main

import (
	"fmt"

	qrcode_memwriter "github.com/kfatehi/go-qrcode-memwriter"
	"github.com/spf13/afero"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

func generate_qr(AppFs afero.Fs, path string, content string) (string, error) {
	qrc, err := qrcode.New(content)
	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
		return "", err
	}

	w, err := qrcode_memwriter.New(AppFs, path, standard.WithQRWidth(4))
	if err != nil {
		fmt.Printf("standard.New failed: %v", err)
		return "", err
	}

	if err = qrc.Save(w); err != nil {
		fmt.Printf("could not save image: %v", err)
		return "", err
	}

	return path, nil
}
```