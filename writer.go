package qrcode_memwriter

import (
	"os"

	"github.com/spf13/afero"

	"github.com/pkg/errors"
	"github.com/yeqown/go-qrcode/writer/standard"
)

// New creates a standard writer.
func New(AppFs afero.Fs, filename string, opts ...standard.ImageOption) (*standard.Writer, error) {
	fd, err := AppFs.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return nil, errors.Wrap(err, "create file failed")
	}
	return standard.NewWithWriter(fd, opts...), nil
}
