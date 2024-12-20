package sessionrecording

import (
	"compress/gzip"
	"io"

	"github.com/gravitational/trace"
)

// gzipReader wraps file, on close close both gzip writer and file
type gzipReader struct {
	io.ReadCloser
	inner io.ReadCloser
}

// Close closes file and gzip writer
func (f *gzipReader) Close() error {
	var errors []error
	if f.ReadCloser != nil {
		errors = append(errors, f.ReadCloser.Close())
		f.ReadCloser = nil
	}
	if f.inner != nil {
		errors = append(errors, f.inner.Close())
		f.inner = nil
	}
	return trace.NewAggregate(errors...)
}

func newGzipReader(reader io.ReadCloser) (*gzipReader, error) {
	gzReader, err := gzip.NewReader(reader)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	// older bugged versions of teleport would sometimes incorrectly inject padding bytes into
	// the gzip section of the archive. this causes gzip readers with multistream enabled (the
	// default behavior) to fail. we  disable multistream here in order to ensure that the gzip
	// reader halts when it reaches the end of the current (only) valid gzip entry.
	gzReader.Multistream(false)
	return &gzipReader{
		ReadCloser: gzReader,
		inner:      reader,
	}, nil
}
