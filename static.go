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
	tmp.Contents = "1f8b08000000000004ff8c54c18ae34610bdeb2bde25accd8e2d7b934388e35cb213184836cbae038161c06d75c92aa6d52dba4b333161fe3d945ad27a9210a2837197aadeabf7aa5a655994250e0d27d4ec0855f062d827d4bdaf84834f304f869d3939c2e9024bb5e99ddc80dbce514b5ec8823db69bcd578ad4f591d006ff48977591c1794debccc0099d89825023f411498cb726da95e35334712eb8ab71093d9e39359080d63c12aac6f833253d4b43ff28bd6ad752cd9eac36d35024857ae31c3c91d5ea48a79e9dc5318911aed6e77084a98522c8b2b03f4326336e6092c24863e48b314a4fed89ac258b2a741795235a53054b786eb86a5429fbcaf59ac33e8310cee4291a213b3ab4622f14bb48427172eb10feb5c53ed1c07ce4b673c6cb1112829b8a46c76cf06f04c64532f682c63cd1553efb24c639b2a849b44581e347420adf8d83d2e71c7026c1aa07ce2c4d7f5a57a12dd3e31fe5483cec8b3ee73041fe47a62aff109e1169d23ea898fd7ba29838f8d9c261078db7b307d290625c1985679646a71aa7a5b8ea7f6c122bf65d2fb0464c8955e8454ff3c4151150057917061b0b8dfe62a4a15617c3381db9aea8a4a20a3e093ede618fafd7db6f7663e0768f77eb6f77c578fc7c787ff7017b6ca6f79f0fef7ffded803db65791db4f9fb0c7bbdd40a89c3f714c02cab7499d301e264673d1b14cd710b5662df20b2cf1670144923efa3176bf79c07257bcccb087861029bd42c42adf46e3dc449870fbfb8fb71f0fea74267945ab0823c3c8ea4834da3bc11ef70fbba200b8c6028efc62e87b891fb01db331e473f6404daf435c80f17d4e5fce692a6704edfad42ce270bac982eff961b91bca017efb36ff7dd1c0f0938d18bbbaf6e067f33f9c7526fd4de22c27b343056de64e335b967aff45f46a9b7b7c5143ca12c79c07df3b771c3e086de7b862414391d6c54bf157000000ffff8473b5ae7c050000"
	tmp.Length = 1404
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
