package main

import (
	"archive/tar"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestUSTARLongName(t *testing.T) {
	// Create an archive with a path that failed to split with USTAR extension in previous versions.
	fileinfo, err := os.Stat("testdata/small.txt")
	if err != nil {
		t.Fatal(err)

	}
	hdr, err := tar.FileInfoHeader(fileinfo, "")
	hdr.Typeflag = '5'
	if err != nil {
		t.Fatalf("os.Stat:1 %v", err)

	}
	// Force a PAX long name to be written. The name was taken from a practical example
	// that fails and replaced ever char through numbers to anonymize the sample.
	longName := "/0000_0000000/00000-000000000/0000_0000000/00000-0000000000000/0000_0000000/00000-0000000-00000000/0000_0000000/00000000/0000_0000000/000/0000_0000000/00000000v00/0000_0000000/000000/0000_0000000/0000000/0000_0000000/00000y-00/0000/0000/00000000/0x000000/"
	hdr.Name = longName

	hdr.Size = 0
	var buf bytes.Buffer
	writer := tar.NewWriter(&buf)
	if err := writer.WriteHeader(hdr); err != nil {
		t.Fatal(err)

	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)

	}
	// Test that we can get a long name back out of the archive.
	reader := tar.NewReader(&buf)
	hdr, err = reader.Next()
	if err != nil {
		t.Fatal(err)

	}
	if hdr.Name != longName {
		t.Fatal("Couldn't recover long name")

	}

}

func TestPhpLongName(t *testing.T) {
	// Create an archive with a path that failed to split with USTAR extension in previous versions.
	path := "testdata/craftTest/craft/app/vendor/aws/aws-sdk-php/src/Aws/ImportExport/Exception/InvalidFileSystemException.php"
	fileinfo, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}

	hdr, err := tar.FileInfoHeader(fileinfo, "")
	longName := "craftTest/craft/app/vendor/aws/aws-sdk-php/src/Aws/ImportExport/Exception/InvalidFileSystemException.php"
	hdr.Name = longName

	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	if err := tw.WriteHeader(hdr); err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	_, err = tw.Write(b)
	if err != nil && err != io.EOF {
		t.Fatal(err)
	}

	if err := tw.Close(); err != nil {
		t.Fatal(err)
	}

	// Test that we can get a long name back out of the archive.
	reader := tar.NewReader(&buf)
	hdr, err = reader.Next()
	log.Println(hdr)
	if err != nil {
		t.Fatal(err)
	}

	if hdr.Name != longName {
		t.Fatal("Couldn't recover long name")
	}

}
