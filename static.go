//
// This file was generated via github.com/skx/implant/
//
// Local edits will be lost.
//
package main

import (
	"bytes"
	"compress/gzip"
	"encoding/hex"
	"errors"
	"io/ioutil"
)

//
// EmbeddedResource is the structure which is used to record details of
// each embedded resource in your binary.
//
// The resource contains the (original) filename, relative to the input
// directory `implant` was generated with, along with the original size
// and the compressed/encoded data.
//
type EmbeddedResource struct {
	Filename string
	Contents string
	Length   int
}

//
// RESOURCES is a simple array containing one entry for each embedded
// resource.
//
// It is exposed to callers via the `getResources()` function.
//
var RESOURCES []EmbeddedResource

//
// Populate our resources
//
func init() {

	var tmp EmbeddedResource

	tmp.Filename = "data/stdlib.mon"
	tmp.Contents = "1f8b08000000000004ff8c92bd0adb401084fb7b8a690212b6f5533b4e139c2a450a170161f0d95aa125a793b89f80097af7705a598953a511d2ddec7c338bca5295252e3d7b746c088fd106cdd6a38bf61178b41efaa766a3ef86707fa2a54e4713f6e0613234900dd4822deaaafa909ca6e808c3687fd0b35062ce051542608f49bb80b1c3181d7cd0b6d5ae3d18be3bed6440ad535fd8f90012469ad016da39fd4cf7af70e8922a930be4f8a50047213abb9e35d515f951cd9beda52738f26f8e3848466dcc0be871fefef9fced82d09340deb0c96125ac5443219d46137042733d2a0570870c866cb6e4cef109f5aac6a2679c501f150074a3cbc0f828f27c93a53aabe9147d9fb9e56b2f851bbee6320ef06e27af73f25b1eb28835d5df3bf8aaff63b346fb7f2a6e75848e54a8da920a4daa367f4a1f6ac938a78594256ea2838dc6dcd20fc1c364f8c1013d392ad4ac7e070000fffffeed60d192020000"
	tmp.Length = 658
	RESOURCES = append(RESOURCES, tmp)

}

//
// Return the contents of a resource.
//
func getResource(path string) ([]byte, error) {
	for _, entry := range RESOURCES {
		//
		// We found the file contents.
		//
		if entry.Filename == path {
			var raw bytes.Buffer
			var err error

			// Decode the data.
			in, err := hex.DecodeString(entry.Contents)
			if err != nil {
				return nil, err
			}

			// Gunzip the data to the client
			gr, err := gzip.NewReader(bytes.NewBuffer(in))
			if err != nil {
				return nil, err
			}
			defer gr.Close()
			data, err := ioutil.ReadAll(gr)
			if err != nil {
				return nil, err
			}
			_, err = raw.Write(data)
			if err != nil {
				return nil, err
			}

			// Return it.
			return raw.Bytes(), nil
		}
	}
	return nil, errors.New("Failed to find resource")
}

//
// Return the available resources.
//
func getResources() []EmbeddedResource {
	return RESOURCES
}
