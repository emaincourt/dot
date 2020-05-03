package archiver

import (
	"compress/flate"

	archiverv3 "github.com/mholt/archiver/v3"
)

type Archiver struct{}

func NewArchiver() *Archiver {
	return &Archiver{}
}

func (c *Archiver) Zip(source string, target string) error {
	zip := archiverv3.Zip{
		CompressionLevel:       flate.DefaultCompression,
		MkdirAll:               true,
		SelectiveCompression:   true,
		ContinueOnError:        false,
		OverwriteExisting:      false,
		ImplicitTopLevelFolder: true,
	}

	if err := zip.Archive([]string{source}, target); err != nil {
		return err
	}

	return nil
}

func (c *Archiver) Unzip(source, target string) error {
	zip := archiverv3.Zip{
		CompressionLevel:       flate.DefaultCompression,
		MkdirAll:               true,
		SelectiveCompression:   true,
		ContinueOnError:        false,
		OverwriteExisting:      false,
		ImplicitTopLevelFolder: true,
	}

	if err := zip.Unarchive(source, target); err != nil {
		return err
	}

	return nil
}
