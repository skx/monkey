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
	tmp.Contents = "1f8b08000000000004ffc456616fe4b611fdbebfe2c1682fda7877e5f59d9bbb3abaa0b83a808b360d1c1728eab8302d8d56ac295220293bdbc2ffbd1892d2cabead73bdb6a880bbb5a8e1ccbc376f86ccf3599ee3b2910eb55484d2682fa476a87b5d7a69b483b81752895b45b8dda2a25af4ca2f20db4e514bda5305a9b13e3afa257bea7a4b688dbea3ed6a169dcb15ad6204e9d009eb616a98dec279a12b61aba592b756d871c3798dade9f1205d036fd08a3b42d908bd21c7efbea18fb64ed2ada8969a2a4ea6214becea0ba5a0892ade6de9b697aac28df3c2cb72b5313710b5270baaa4977a033f90b18070ecc637c2ef88e1f0d4de52555185d2745b86e3794f692ac24323cb86914a5daa9e6da48e4e081bd26485a72a31b494da93ed2c79b2035b97666f8abda310f946b69d12dadfc01ba3864d89b1cae82f3c84b224aa2d1a714f137ba99d174a51859a3ca7e8a1e41dc1995fa742f1b331d890c7b20736d237fdedaa346deeee7eca5360d60be3f9ce3cc0d28028e436b2724fd649a3476282b284ae4664be21f631818f07e91bae951d4a3dc92a85c652eaaef7a884173996a6f7fc36d6913d028c20563890334bd8fe207c432d575c28ae256bcf737567a5d1cee3fb731478bd5abf599fbc3bfed5c9eb93b7efbe7af7fa347d3d030ae078f5d5faedf15bfef7e6e4ddd19b93d3c1fb07f621b477b05493b5414626e05e3642578a1c8b89ab58a136163f5cfef6fcbb05fffcf14f97e1f7ece282538dd984cf1cf368c8e0874bb64481f564e5ece202058e4f4794df4aeb3c283626d32f3484b562cbae878e46cd5659fc8039fe31032cf9deeab47675748df9e9ec71747bd9102cb9271eb18c8d2d941a023a9cfdf9c3d9f7975cde18e44958f69022a4a88a3caff6caa3c0d5f5e96c06c81a1914e92ce43dc77bac933582bd8c1c70a56b6333487c1dcde7a319c3494ebbde35990d6f8b08f84a5ecf4fc376401e1ec63f1f7921fc178948594d39f8bdf804669570cf208e70627430a0a331d3182d42bdda815eae638e8f4c489ee326da41f74add84d9d2764a96d28709b70a9562aaf31ce72eb0bf91f7944a0f6a3bbffd86bf8f0a08dcaee2872c2a6097a92355cf5114933c99511608bced29503625ab16cad194ab0f2c0a9e93318d31acd120513620eded76912642f02b86644dcd1b1989f4d43aae31e248657f9da54a96c2b31c791f8f74db53e8f4314a04574be5c966932d09278b8e45741470f09b42112414808fab51353b61722a496e83ab24c900070578ff95bc0e1e585d35b231e12c18cd530ef1e1d86394a0d4f4ba482e633240649b5d26c1066124f5a43d81ff78e05e046e0261a6ae1df1695d43e8ed22d11bb8858c26912edefa1183bacab80a29694ef745e2b889763db923898918c8615db1cfddd711874cc40569ed9a3362cd7368e3519b5e572b0e945a62b93e9d3dee06e0ef4c429566a063c4a3b86eb78194a84ae7794cef03fe372375c6c7901db264e816050e0e428efcfab34ce439be357622f890f6337e1809d8f43ce83ed5925b5c1bbd0c0d0a5155216d479db0c21b1b1c4d1ad63e1d2acce86111f20fd9b27cc638ecccc4c26bfa693c2b924b8bc32211932535cfe3d6fdcafb587423d7e92229c3856a04666a88ae535ba937dce3dcd3b11aa3f6bc4994a5432cd5d3f55da7245550d2f97d456b4597a1d6e5b4642fd62835efd880e9fc795e21a6339e3cbeed507088919a446fd0c3e8e76923fbb61bcdf673180ea74024f7609ee382f8fe442395bc3cd213bbd5469334ba598eac4e4610937936d1660995c2fbe9501ff6a58c17b147d5721d8eff005c2d970c925b32355cac39e7f4fcb469846bf61d36bcfed259b3baa3adcbe6ffd9917341adb9272812154beea1919e5c274a426d4d1b84b6afe1d39af256b689cd480b6db8ddff9afdf8a33b9c67ab2fe7bfd8b5bee9f95868852ff96eb189bccdc3ed25ce3a364843939b318ce38c57afd6d77330a178042947e1eec54fb40905889fc3993e8cf3a48794ed82d7474124047b14318c2b2ef3a70922b63fa93a882024829f93c04524de5b21d567306fff25f3ab2fbf9933fdff2dea8f3e83fa4b2bdb5153af5e0619e7582ac7be123d47ca3561b207f505bea31032a891992737f10f46df93f541cfd3803c36957920bb2c857b761f4a2979130ca62affa4f9c87d7010c5fdecf062e1c6762931b9fda4f570ea97785fe0e0370778f50a25be2e70f09703ec5a233cac4f7650ae8cadb248c3f0947c20bd3e0ed40c4fb42d1bbbb3e5f9c44f98b0297e10731976c62b44a4360db0d85a2fd1d977ddcb740683ff079d6242e7dfff4d3a97c5ff80ce7f060000ffff2458f09945120000"
	tmp.Length = 4677
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
