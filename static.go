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
	tmp.Contents = "1f8b08000000000004ffc456616b1b4710fdae5ff130349189255976dcc4552e505217026d1a1c154a8dc1ebbb39dde2d5eeb13b67572dfeef6576f74e72923a2914ba60cbda9b7933f3e6cd9c67b3d16c8665a3036a6d08a5b3acb40da83b5bb2763640dd2a6dd4b5215c6f5051ad3ac307d0ebd6d09a2c53056d313f3cfc4690dace13d6ceded0663a4ae07a4ad3144107b4ca335c0dd7790456b652be9a187ded951f1cded6d8b80e773a346087b5ba21948db22b0af29d1bfac47527dd8a6a6da992641af224504f8d8125aac4dbd375a74d85abc08a75395db92ba89ac9832acddaaec03d1907504160b851bc2546c2d3fa9aaa8a2a94aedd48392c3ea5ab08778d2e1ba954dbd27462a36d0221acc892574c556668a22d936f3d31f99eada5fb6c8a5da018f94aaf5ba32c5f819d33bd5366ac72f62943194faadaa051b7b463af6d60650c55a889254586d13784e0becb8d92b372581163d2012bcd4d773d2ddd7a166efe98e5c0512f7256ae877cc4522a7fe7eee0a9af3d5631f0774b3e6867070aa30695ad060eb821c1d8210a779a1be9aaef45b1937f4e12136ddb8e512956334c5cc7f26de8b820025241d242a4719459f85971436bd18632d27551298b0e46a5b381f1fe2d0a1c4fe7cfe727a747df9e1c9fbc3c7d717abcc84fcf8002389abe98bf3c7a293fcf4f4e0f9f9f2c7af43782a12c0778aac9fb283817eb9e34ca568682c84efa5da1761e1f963fbc7d77201fbffcba8c9f67e7e7926aca263e9698877d061f96628902f39d9bb3f37314385a0c55fea87d60501a61a15f5928efd546a0fbd9472d56e3f400fbf86b0478e2cedb7c77717889fdc5e87e805d36044fe101222669052863fa800167bfbd397bbf94f6a6200fc20a428e90a31a62b9ed0ca3c0c5e5623402748d310cd971cc7b1faf31cfd688f63a71209dae9d1f43e35532df1fcca49c0cda76a119fbf8ed20157ca12ff717d11dd0cf9ea53fefe522fe4a44e4ac7639f8497d05b346858f4a1cca49d121051d0e99a668a9d48b6dd19379caf15e0899cd7095ec603b63aee2165ab746979ae32e9cc64e09d5b319de483f6499adf42dd96181c259902a1b9065bf39c8c3283d1f242272e9e792691d845ea4bd2778ada74a978a4509e2277bd77714876c5056ecd9b4d686c98f775c92c886fe1d46d6a5fb0645ec5e2053a7b6c86d6ad85613924aee740f95d510cb4101f1bfd09711571a5b633c243c8e46fb3987741e448922c9410f32644a0648aa10c8ac95d893dcb8ec136592de8ae7919b4898abeb40f24aada1ece620d30be1163a9924bac4f513066d3516cb9cb4a42bc2ff47e244bfdb71d8922444f4e4a02852f40c2a270b4b67e2e2086ce722d53a9bc13a46ed3a5b4d2550769acc17a3fbedeec9a5abbc7164c9caff1c3abe7b453571245d0dd5b666a3ed4ae42acafa48a9ecb252f316cb5c85ae6d8da60a4607fe1c676bd58e51dbb22fef8b9c65097dbdd878dda29010e32cb74124126bc07928275eb783d99734340cb1b3b7e479879dc0fd3bc5b83bf29352858f462f594cd9458371eeb124f6a8707a1250606f2f2ee0cfea280295d81934f114a9c8ba2ef1bac0def77b78f204255e15d8fb7daf6f437f24130128a7ce57e38192784a3c2b707c9445984eb22d1bbfb5cdea4c34e6f85e5ccbe8793f1a6565fa9d897c8ccbae6d1fe7321afc1f5caa1d2efffc975c4efe732eff0e0000ffff1a41adf84f0c0000"
	tmp.Length = 3151
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
