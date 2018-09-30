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
	tmp.Contents = "1f8b08000000000004ffc4577f6f1bc711fd9f9f622ab40919933c518e1bbb321d04ae03b868d340568ba28a0badeee6785b2f770fbb7b9299c2dfbd78b37bc7134d2b6ddaa2072416f766e7c79b373fae28264541978d0e546bc3543a1b95b681eace96513b1b48dd2a6dd48d61bad951c5b5ea4c9c93deb686b76c2357a42dad4e4f7f054d6de799b6cebee3dd729294eb252f93051da8553e92abc9759e4254b652be5a187de3951f2ebcae69e73abad3a1a1e868abde31958db21b0ef81d1bfee8eac8dd8a6b6db982330d7b86aacf8d21cb5ce1b6e79b4e9b8aae43545197cb8dbb265547f6c4958eda6e28f660cc4905a8898d8a7b60609eb7375c555c51e9da1dc289b853ba8ae9aed1658348b52d4d07196d9312a60d5bf62a7295115a681bd9b79e23fb1ead4b77d4c52eb058bed6dbd6281baf293a67fa4b19b1cad9cf2329e359553b6ad42d8fe4b50d5119c315d51ce16224a3df3105f79b9c283c1b471b8eb4e888363a36ddcdb274db22bc7b5f64c3e00be2f9cedd91e73e22f16d40e5967dd0ce0ec008b394ad86c862c3d0310a9fee746c902bdfa77ae455364d0b6ddb2e52a5a22a68e1ba885f431ea1910811a40c0b38931cdb1f546c788b8c2b835c827b11d99d94ce8648dfbfa6353d5eaebe5c3d7976f6eb278f9f3c7df6d5b3c7e7f9ed2ba235d1d9f2abd5d3b3a7f8efcb27cf4ebf7c72de6b7f091dcac6409e6bf65e68e424ee45a36c6538804cc86245b5f3f4e6f2b7afbf9be39f3ffee952fe7d757101579337f21a364f7b0fde5c4292d6b41a9dbcbab8a0359d9d0f517eab7d88c4a93001bfb2a4bc573ba8ee2b9a6a484dd30b9ad13f26449e63e76d3ebb3a7d4bb3f3c98741ed65c3e439dcd3488b54d8ca98de60a0577f79f9eafb4ba43719b967161ab2856cd570c4696722ade9eaedf96442a46b9a92613b15bf67f48256599a445e270c90e9daf929697a9ec4678318c2c94adb2e34532fbfe629e02bfd76762ed789f4a347e9cf0f3890ff2520b257630c7eaffe05648d0a07210ee124eb84804e074f93b514ead53ee8c52af9f8018014055d2739b29d31d7d25bb6add1a58ed2e1969229405d14f43a08fa1b7dcb39f5c4db36eebec6fb810182ed32bd982606ec3d0d6cea19add7233f81280842d1772c908dc1aa95093cc6ea2548813e99dc18cc3a4bacca86d846bf9be78e207a55efacab711191e8c8db801c536aa9d0d77aae74a922e8887b68e9be63a9f4c14a0aaed626b29f8eaee438413a90e854e2c02f436ba190043e9c26d6ec89095732dd7a559992120ead09f7aff45bd10076d5341d1c9e8ad02cfb901ed81eac0853f3cf7956999c214a68436526ac1023b327df11fcd3c0bd106c043057d78131ad6b527637cff00ab6a49348820b573f42d0565364213b0d771f040e45b4afc93d4800a20707bc82cefddb210e9d81136aed8b33c55a14645da4da75b65ac2502e89c5ea7cf261df007fe77254b90706443c90eb6627a024568688367d2cf0bf3b6da71843bef712a17b5ad3c989f8889f3f894451d0b7ce8f082f6e1fe0834808a2af85f739972871ebec420a94545589db815be555745e148d0ad6df6f2a40f4d15afc176f419fc10e94b99478cbef875991557a7ab4cec04c339b67e9ea71e67d4cba01ebbc486a59a886c05c4daa6dcd4edb0d6a1c359db231702fba0c591e62399fa16b5ba3b922a3433c96b4ad6aa754db729cb20773948b7728c03c7f0e330438d3e489db96d630314093e1153e0c7aee1772dcb683d8710c65380990a8c1a2a00bc6fec40394381ee049d5ea93486edda023d889089233071d6d92a332f462dcd4fb7bd9e379aa51b358c9f897c0cd6281205192b9e0c6397fe33c96d9f426c87040b123abe227781c44e6480ce9453f7ee00bd2b58239fcf83884a2a06fb0cb80443f3af685b3dcd7392c2973a776211b04a1a5ef187a4e673d29b2abe2290c496fc1d4ac499151f80a90e6a403852d96658fe55fca652265caca1bcd3e49dd8b13dd298504d307652e588a3bb9aae879fe0ba35e3630791292fd409523408f624f9d5f1cce52987c63debcb953adf874ab4cc74029deb93e11b6e2f71c8e154eb843e5a839dd649480befab34202c4db1b996838bd199daab4aee1f83ef7904599ad780525d86ba06826fbdd2791d1180d2afb901e68807221a89f8b577d357d203681f7d0257845491fc8f01ce891307a3d441fabca053f98a6be36f4be3270efa732946a05dc2f0a2916494f6a78522047f3e17ccc852d58fd42b2302a97812f080b684bcd649f0f00c77152237bf2fe2e28854d1b894158233ede132291a0b4dc64aee8c56a4e7a0c60b2039d89a6191c9c67b2c2ce98acf7f6d44685e6d89a8af387b6d4e53bde85e9ec3f5b562f78eb6e990cab0ac3eaaed19143ab4aa6dabbad64ecd8aa90cf4cf47a9bd32509f1bcc1a2f0b7e90f3f8447b3e9f28bd92ff74b83ebb0506e552cf155b2491d37d545e22f0486fc66eca638bd5abd9d115ad6215befe12bb8efbfdb864992bd9d838ac328c9111c9925fda203821d10ea13a3242d0e6c6a191fe2277d7278c00b9973027cf44a9b9f81bcff24f2cb2fbe9e01feff16f4a73f03fa4bafb703a73e7b38c8d410723a00ce618a0e23454e0076cf3ec13b11614a6640e6de37fc4b676fd98f3b50328885cbb83bf68b5285832fa924b18c4e04c62c47db91267f8c23483dbeec510727d2f48fb7200895945a8b7c37e1269a081a53492fd674f2cd097df61995f47c4d277f3da17d69c8d32b2897ce57d30443ff9458651f9f0934fd0363e5b26cfc5e162d0a8f34ae6c5fc85cca4d99b719dadcced337d64370766dfb309c22f0ff80538de0fcf1df8473b1fe1fc0f9cf000000ffff631723de7f160000"
	tmp.Length = 5759
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
