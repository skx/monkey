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
	tmp.Contents = "1f8b08000000000004ffc4567f6f1c3d11feff3ec5a308faeebdb99f49431bd26d85429082a054e92121d2a038bbb3b7265e7b657b130e94ef8ec6f6ee5dd2104a01bd2bb597b5c733f33cf3cc78e7f3d17c8e552d1d2aa90885d15e48ed5075baf0d268077127a412378a70b3414995e8949f4036ada286b4a7125263b958fc9c3db59d253446dfd266368acee58c663182746885f530154c67e1bcd0a5b0e554c91b2bec70e0bcc2c674b897ae863768c42da1a8855e93e3775fd3574777d22da9929a4a4ea6264becea07a5a0894a3e6de9a693aac4b5f3c2cb62b636d71095270b2aa5977a0ddf93318170ecc6d7c26f89e1f0d4dc50595289c2b41b86e3f94c614ac27d2d8b9a914a5da88e6da48e4e086bd26485a732313495da936d2d79b23d5b2bf36c8a9da310f95a36ad12da5fc31ba3fa4389b1d2e81f3c84b224ca0d6a71473bf6523b2f94a21215794ed143c95b8233bf4c85e2676db0268f6907aca5afbb9b59619ab9bbfddb3c0566bd309e8fe61e967a4421b781953bb24e1a3d10139425743920f335b18f1df8b897bee65ad9bed43b59a5d0984add761ea5f0628ea9e93cbf0d75648f002388150ee48c12b6df0b5f53c315178a6bc9daf35cdd5161b4f3f8748e1c87b3e5ebe5d1f1c12f8e0e8fde1ebf393e3c49bb67400e1cccde2cdf1ebce57faf8f8e17af8f4e7aefa7ec4368ef60a9226b838c4cc03dad852e1539161357b144652c3eaf7e7dfe71c23f7ff8e32afc9e5d5c70aa319bb0cd31177d069f576c891ccb9d95b38b0be438381950fe465ae741b131997ea121ac151b76dd77342ab6cae206c6f8c708b0e43babd3dae5e20ae393d1c3e07655132cb9471e318d8d2d94ea033a9cfde9f4ecd38acb1b833c0acb1e5284145591e7d54e79e4b8bc3a198d00592183229d85bcc7788f65b246b0979103ae74656c068977d17c3c98319ce4b4ed5c9dd9f03689802fe5d5f8241c07e4fe7efcf38117c27f918894d52e07bf13dfc0ac12ee09c4014e8c0e06b418328dd122d4cb2de8e932e6f8c084cce7b88e76d09d52d761b634ad9285f461c2cd42a598eaf91ce72eb0bf9677944a0f6a5abff9c0fb830202b7b3b89145056c3375a4aa31f27c274f669405026f3b0a94ed925509e56897ab531605cfc998c610d66890286a90f67633491321f8157db2a6e2838c447a6a1cd71871a4b2bfd652290be1598e7c8e47baed2874fa102582aba4f264b39d2309278b8e45b40838f84d210f120ac087d5a89aad30399524b7de55926480831c7cfe525e050facae0ad99070168cc62987f870ec214a506a7a9d2497311920b2cd2e93608330927ad299c07fbc702f0237813053558ef8b6ae20f46692e80ddc42469348171ffd8a415d665c859434a7fb2271dc44db9edc92c444f4e4b0aed8e77677c0211371415adbe68c58e77368e351994e97330e945a62ba3c193d6c07e06f4d429566a063c483b86e368194a84ae7794c3f07fcaf46ea8caf21db67c9d02d72eced851cf9f5fb99e03eb78fe70093b09f879021c050718bfd3c659a25798d19fd537a121bf66b150ce0d3979d0c5f38dc4061449a0aa26dd546ea35371d3759a467108337a969d3ad9208765ddb2a49259474fe39161bd166a874b1cbe18ba4a56e1a3a225d08cf4a2a94c0372d720e315093c80bbb839fc79de59b7630fb77edc44df174a6d6c2d5cf8d545e7f69a2ce6e69e3b2f17f37582fa831770445a2e43aded7d2936b4541a8ac6942f59e93755a53decaa61ff7ac614b6b16f55fb22f5fdcfe389bfd38fed956e0a6e3e1d7085ff00dba9e840e1e873b3a76341ba4d1c0fa0d4327e3d5cbe5d518a1140f20e5287c61f0136d82a3b81d6eae7e68f1572341a46c27bc3e883021b0d128610810faa6e49727837c94b4a3f07ef72e8b3d45aaba54d3653fadd574ca29f1ec4959c65ee22ce6f39e786f8554dfc1bcfd97cccf7efc3066faff57d42fbe83fa9595cda0a9572f838cc3219583c9795aa2a748b9264c76af3e6639519c410dcc3cfade3c35fa8eac0f7ade0dc8b348997bb2d342b827b77e4ac99b6090141254f14d4387fb602f8afbb92b2c382ab073c7f3b8e231cc9fab05dee7d8fbd51e5ebd42817739f6febc876d6b8487f5c90e8a99b1651669e89f82a7fce141a0a67fa26d51dbad2d8b939f30b652fc20e6229c8c17e523f5c6d67a89ceae6d5fa63318fc14748a1d3afffe1fd239cdff0f74fe330000ffff5e294d7d2b110000"
	tmp.Length = 4395
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
