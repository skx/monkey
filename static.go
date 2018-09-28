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
	tmp.Contents = "1f8b08000000000004ffac555f6fdbc6137cd7a798971f2221faebc4bfd855d517d7050cb469e0a84001c3804ee4523cf87847dc2ded0a85bf7bb177475a498db40f25604b2477677767e6568bc568b1c0b6d601953684c25956da06549d2d583b1ba01e95366a6f08fb234aaa5467780addb4861ab24c25b4c56ab9fc9f20b59d2734ce3ed0713e4ae07a4ef3544107b4ca335c05d7790456b654be9c19bdf7ca0f0937158eaec3930e35d8a1510f84a256f64041eeb9a6bfa59eb45b52a52d95d24c4d9e04ea8d31b044a5647bda77da94d80556ac8bf9c1eda02a260f2a356b7b00f7644ca182c070adf88518294fcd9eca924a14ae3dca382c39852b094fb52e6a9954dbc27412a36d02211cc892574c656668a62d936f3d31f99eadad7bb5c52e50acbcd34d6b94e51dd839d32765c64a67df3094f1a4ca236af54827f1da0656c650898a585a6418fd4008eebb2c945c07870331661d70d05c77fb79e19a4578f863910b47bfc875703de4372265f28fee099efad9e314037f8fe4837676a0307a50d972e0806b128c13a2f0a4b916557d6f8a93fe73939869db768c52b15a60e63a96bb417141046482e48548e328b3f08be29a1af18632a2bab894c507a3c2d9c0f874830ddecd57ef57e79767ff3f7f777e71f9e1f2dd3abfbd0636c0d9fcc3eae2ec42fede9f5f2edf9faf7bf42bc15096033c55e47d349c8b73cf6a654b43416c277a97a89cc7e7ed8f371fa7f2f1eb6fdbf8797d7b2bada66ee26ba9b9ec3bf8bc95486cb03a79727d7b8b0dced6c3943f691f18948eb0d0af2c94f7ea28d0fdd9472551e3f40213fc39023c71e76d7e76b7bcc7643d7a1e60b735c153f80211b3b40294317dc180ebdfafae3f6d45de54e48bb282902be4aa86589e7686b1c1ddfd7a34027485310cd971ec7b821fb0cad188f13a71204a57ce8fa1f17d0a9f0c61324e066dbb508f7dbc9ba681eff4fd641dd301fdf66dfafa2c0fe2bf4444eeea94839fd5bf60d6a8f0d588c338a93a64a0e5d069aa9646bd7b197ab64a3d3e0b218b0576290eb6336617b750d31a5d688ebb701e9512aa170b5c891eb2cc0efa91ecb040e12c481535c8b23f4ef36114cd078b885dfa73c9d404a11769ef095eeba9d485627182e4c9def51dc54336382b6a36afb461f2e3b83ea72789c96a838acbc8fd6b1e90d22fcaca163e155732e218d8c41277fa3e2289a015c643a3e31834c955d325b9c90ed82473e4db69861cbc11dd2090d923518b2c58ce39b5c76d24e585ccfc8babe32f8f70160de92aa8b635476d0f2296f0fa954eecb24ef90c6b1ba342d7b6465309a303bfc679a3da9ef0ca16ff25d5dcb4d8086814e2f4f87c834c6eda81c97f60f0af000000ffffa94fea96ac080000"
	tmp.Length = 2220
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
