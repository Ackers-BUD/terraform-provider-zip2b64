package client

import (
	"encoding/base64"
	"testing"
)

func TestDodgyZip(t *testing.T) {
	response, err := ZipExtract("DodgyInput", "filename")
	if response != "" {
		t.Errorf("Expected '' got '%s'", response)
	}

	if err == nil {
		t.Error("Expect Dodgy Zip to fail")
	}
}

func TestFileInZip(t *testing.T) {
	zip := []byte("PK\x03\x04\n\x00\x00\x00\x00\x00\x14OOQ\xddDYc\v\x00\x00\x00\v\x00\x00\x00\t\x00\x00\x00file1.txttextfileonePK\x03\x04\n\x00\x00\x00\x00\x00\x1aOOQq?\xb5]\t\x00\x00\x00\t\x00\x00\x00\t\x00\x00\x00file2.txttextfile2PK\x01\x02?\x00\n\x00\x00\x00\x00\x00\x14OOQ\xddDYc\v\x00\x00\x00\v\x00\x00\x00\t\x00$\x00\x00\x00\x00\x00\x00\x00 \x00\x00\x00\x00\x00\x00\x00file1.txt\n\x00 \x00\x00\x00\x00\x00\x01\x00\x18\x00#~\xfb\xa7\x85\xa2\xd6\x01#~\xfb\xa7\x85\xa2\xd6\x01\x00\xa8\xdfE\xe7\x8b\xd6\x01PK\x01\x02?\x00\n\x00\x00\x00\x00\x00\x1aOOQq?\xb5]\t\x00\x00\x00\t\x00\x00\x00\t\x00$\x00\x00\x00\x00\x00\x00\x00 \x00\x00\x002\x00\x00\x00file2.txt\n\x00 \x00\x00\x00\x00\x00\x01\x00\x18\x00u\xa7v\xaf\x85\xa2\xd6\x01u\xa7v\xaf\x85\xa2\xd6\x01\xa8\x9eˏ\x85\xa2\xd6\x01PK\x05\x06\x00\x00\x00\x00\x02\x00\x02\x00\xb6\x00\x00\x00b\x00\x00\x00\x00\x00")
	Base64EncodeZipFile := base64.StdEncoding.EncodeToString(zip)
	Base64file, _ := ZipExtract(Base64EncodeZipFile, "file1.txt")

	if base64.StdEncoding.EncodeToString([]byte("textfileone")) != Base64file {
		t.Error("Failed to extract file and decode from zip")
	}
}

func TestNotFoundInZip(t *testing.T) {
	zip := []byte("PK\x03\x04\n\x00\x00\x00\x00\x00\x14OOQ\xddDYc\v\x00\x00\x00\v\x00\x00\x00\t\x00\x00\x00file1.txttextfileonePK\x03\x04\n\x00\x00\x00\x00\x00\x1aOOQq?\xb5]\t\x00\x00\x00\t\x00\x00\x00\t\x00\x00\x00file2.txttextfile2PK\x01\x02?\x00\n\x00\x00\x00\x00\x00\x14OOQ\xddDYc\v\x00\x00\x00\v\x00\x00\x00\t\x00$\x00\x00\x00\x00\x00\x00\x00 \x00\x00\x00\x00\x00\x00\x00file1.txt\n\x00 \x00\x00\x00\x00\x00\x01\x00\x18\x00#~\xfb\xa7\x85\xa2\xd6\x01#~\xfb\xa7\x85\xa2\xd6\x01\x00\xa8\xdfE\xe7\x8b\xd6\x01PK\x01\x02?\x00\n\x00\x00\x00\x00\x00\x1aOOQq?\xb5]\t\x00\x00\x00\t\x00\x00\x00\t\x00$\x00\x00\x00\x00\x00\x00\x00 \x00\x00\x002\x00\x00\x00file2.txt\n\x00 \x00\x00\x00\x00\x00\x01\x00\x18\x00u\xa7v\xaf\x85\xa2\xd6\x01u\xa7v\xaf\x85\xa2\xd6\x01\xa8\x9eˏ\x85\xa2\xd6\x01PK\x05\x06\x00\x00\x00\x00\x02\x00\x02\x00\xb6\x00\x00\x00b\x00\x00\x00\x00\x00")
	Base64EncodeZipFile := base64.StdEncoding.EncodeToString(zip)
	_, err := ZipExtract(Base64EncodeZipFile, "file3.txt")
	if err == nil {
		t.Error("Failed to NotFound in Zip.  It was found?")
	}
}
