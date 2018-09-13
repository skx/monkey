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
	tmp.Contents = "1f8b08000000000004ff8c545f6be346107fd7a7f8bd94b3b9d8b293b897d4755fae2e04daeb9173a110025e6b47d692d5aed81d253525dfbd8c56d2396d291524b276677fff66a43ccff21cbbca4494c6120aef58191751b6ae60e35d847a56c6aa83251c4ed054aad6f2054cdd58aac931691887e562f18d20356d20d4de3dd1699e257033a7796230118d0a0c5fc2b7019195d32ae8993587a0c278e0aec4c9b77831b1027bd4ea895054ca1d29ca3357f48fa367723595c6911631150512a877d6c21169391de8d01aabb18facd814f3a3df43954c01a40d1b77040f615c404581e14af1d760849eea03694d1a856f4e6287e54ce135e1a53245254e8d2b6c2b35c62510c2911c05c5a4fb8466c6318526105318d2daf97f95d846ea98f7a66eac72bc077b6f87437d62dabb770c6503297d42a59ee9acdeb8c8ca5ad228894522c39a2742f4dff58d92ebe87124c6ac058e86abf6302f7c9dc7a73ff29eb89b17b98e7e80fc8f4a71fec9bf20d0e0bd7331e6f74c211aefc608bb19544e8f197045827116145e0c57d2d5300cc599fe5e2466c6352d432b563966be65791a3b2e88803848b3d0c598f529fca2b8a25a664359e9ba4c29cb1c64857791f1f90e1b5ccd97d7cbd5ede5b7ababd5cded87dbab75bfbb0536c0e5fcc3f2e6f246feae57b78bebd57a40ff2818ca7144a09242e806ce77be679572da5294b1937e6b943ee0cbeec7bb4f1772fbf5b75d77dfdedf8bd4a4a6db16cec5a0e0cb4e2ab1c1f26c657b7f8f0d2ed7a3cb9f4c880c4aafb0c4af1c5408ea24d0c3bb8f52aa26690353fc990181b80dae5f7b583c62bace5e47d85d450814df2062963e01cada813062fbfbc7ede79db43791bca115849ea167b5c4b2da5ac6060f8feb2c034c89092cb949a77b8a1fb0ecabd1d59b948174baf4610283ef53f9742c133b3d68d3c66a12baa78b64f8c13c4ed7dd71c0bc7f9f7ebeca42f72f05d1ab3acfe067f53f92b52afecde26827b1430c2d46a5892d597df86a7ab64c1a5f25903cc73ed5c1b5d6eebbaf50dd585318464581e6d96bf657000000ffff427a7ca6f1050000"
	tmp.Length = 1521
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
