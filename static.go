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
	tmp.Contents = "1f8b08000000000004ffc4566d4f1c3717fdbebfe2083d4f321bf61542134a365545a914a94d23b295aa122acccc9d1d0b8f3db23dd06dc57fafaeed995d4842d2a8552dc1b29efb76ce3df70ed3e9603ac5b2920ea55484dc682fa476285b9d7b69b483b81652894b45b85ca3a052b4ca8f20eb46514dda5301a9319fcdfecf919ad6126aa3af683d19c4e072429398413a34c27a9812a6b5705ee842d862ace4a515b6777855626d5adc4857c11bd4e28a905742afc8f1775fd17bae5be516544a4d05175391250ef5582968a282bd2d5db65215b8705e78994f56e602a2f4644185f452afe03b3246108ec3f84af80d319c9eea4b2a0a2a909b66cd703cfbe4a620dc5432af18a9d4b96ad946ea1884b0224d56782a124363a93dd9c69227dbb1b5341f2cb17514325fc8ba5142fb0b786354e794182b8c7eec21942551ac51896bdab297da79a1141528c973891e4a5e119cf93a358acfca60451ee31658495fb59793dcd45377f5fb34250e7ae1b3325dc8072c19f96b73034b1df680a2e7ef9aac9346f714060d0a5df41cf88a38c61651b891bee2aeda4e145bf5a7223196ba693d0ae1c51463d37afed6779c23028c206a21d038482cfc287c45356b4328ee3aabd4b30e06b9d1cee3cd2b2cb03f993f9d1f1cee7d75b07ff0fcf0d9e1fe517a7a022c80bdc9b3f9f3bde7fcf3f4e070f6f4e0a88b7ecc3184f60e964ab23608ce04dce34ae8429163d971bf0b94c6e2edf2bb57af47fcf1d3cfcbf079727acaa5c66ac263ce39eb2a78bb644b2c30dfba39393dc5027b473dcaefa5751e144798e9171ac25ab1e6d0ddeca364ab2c3ec0107f0e004bbeb53add9dcdce313c1adcf6619715c192bb1311e3b80284525d4287935f8e4fde2cb9bd31c99db41c216548591579be6d95c70267e7478301204b6450a4b350f7102f314fd608f63272c09d2e8dcd20f1229a0f7b3386938236adab321bbe8d22e033793e3c0aee80dcdd8d7fdef245f8158948556d73f083f80c669570f720f670627630a0595f69cc16a19e6d408fe7b1c65b26643ac545b4836e95ba085ba86e94cca50fbb70123ac5544fa738e67ef0325bc96bd2fd0285d120915720eded7a9486917bde4b84e5d2cda5a7da31bd887b8fe335960a990bcf4a603fdebbb6a53064bdb242cf26a5549e6cb6e51245d6f76f1658e7ee2b2c42f71ca932b6856f63c3369ae05252a7bb50490d010e1660ff33791ee272634b647dc159301aa61ae2b993258824251da590b11820aa824326ad849ea4c6259f2093f8563c0ddc04c24c593ae2576a09a1d7a3442f985bc86812e962d7f718d445c696a9682e9785ff51e258bf9b71d890c44474e460b188d953503e49583211174660331711eb740a6d3c4ad3ea62c28992d3787e34b8ddec9e045da48dc34b96ffe790e1ddcbaa0923694a88a6516ba9572c5756d63da57a93949ab658e2cab54da3241550d2f90f71568b2643a9f30ede27394b12fa7cb1f9bac1825364496ebd4838571fe7ae9c7cddf4669fd25037c4a7549b6b822251307d3795f4e41a91134a6bea409af3fc92b9cb43ba53deca3a4b4d8e95adb0c0ce6fd9bb776e77984d9e0cffb7131ace0f4dcb83560b9ff3a25c8d825a86611547f5b0410a16362b0f7fc6b767f3f321429c5b9072145e247c923e98a4f8382ca86e4002346f85545f80cd7e14dbe4c9374306f84f819b7d01b8a59575dfb5470f838caa4f2d6372fa1590eeee23e555c99476fd0dd446aa33a89e993b2fee63a3afc9faa098ed843c64cadc901de7c2dddbe15d7a130cb675f4e006eaa68995b613e5f3a185143497636b63b327ef1c7eefe778b9c0ceb73b78f408395e2cb0f3ebce96f8c261d172807c626c91451aba93637781fdbd404d77a26d5ed98d6d5a73711e537ecbae79f08c6b2f520bbbb5dd1fa2b36d9a87e90c06ff059d628bce3ffe269de37f83cebf020000ffff2b56dfea9e0e0000"
	tmp.Length = 3742
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
