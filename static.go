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
	tmp.Contents = "1f8b08000000000004ffc4567f6fdcb811fd7f3fc5c3a24db4f1feb4e326aeb3090ad7050cb469e06c81a28e0bd3d2684598220592b2bb2dfcdd8b2129eddaf13977b93b9c80642d723833efcd9ba166b3c16c8655251d4aa908b9d15e48ed50b63af7d26807712ba412d78a70bd4141a568951f43d68da29ab4a7025263319fff9e3d35ad25d446dfd0663a88cee594a63182746884f530254c6be1bcd085b0c544c96b2b6c7fe0acc4c6b4b893ae8237a8c50d21af845e93e3775fd1574777d22da8949a0a4ea6224beceaa552d044059fb674dd4a55e0ca79e1653e5d9b2b88d2930515d24bbd86efc818433876e32be1b7c47078aaafa928a8406e9a0dc3f17c263705e1ae9279c548a5ce55cb36524727843569b2c25391189a48edc936963cd98ead957932c5d651887c25eb4609edafe08d51dda1c45861f44b0fa12c8962834adcd28ebdd4ce0ba5a840499e53f450f286e0cc1f53a1f8591bacc963d2026be9abf67a9a9b7ae66efe334b81592f8ce7a3b983a50e51c8ad67e596ac9346f7c40465095df4c87c45ec63073eeea4afb856b62bf54e562934265237ad4721bc9861625acf6f7d1dd923c008628503398384ed6fc2575473c585e25ab2f63c5777901bed3c3e9d618983e9e2f5e2f068ff0f8707876f8fde1c1d1ca7dd536009ec4fdf2cdeeebfe57faf0f8fe6af0f8f3bef27ec4368ef60a9246b838c4cc03da9842e1439161357b140692c3eaffe7cf671cc3f7fffc72afc9e9e9f73aa319bb0cd31e75d069f576c8925163b2ba7e7e75862ffb847f917699d07c5c664fa8586b0566cd875d7d128d92a8b1b18e17f03c0926fad4e6b17f34b8c8e07f7bddb5545b0e41e78c42436b650aa0be870facf93d34f2b2e6f0cf2202c7b48115254459e575be5b1c4c5e5f16000c8121914e92ce43dc27b2c923582bd8c1c70a54b633348bc8be6a3de8ce124a74debaacc86b771047c212f47c7e13820f7f6e29ff7bc10fe8b44a4ac7639f8abf811cc2ae11e41ece1c4e86040f33ed3182d42bdd8829e2c628ef74cc86c86ab6807dd2a7515664bdd28994b1f26dc34548aa99ecd70e602fb6b794ba9f4a0baf19b0fbcdf2b20703b8d1b5954c0365347aa1c61b9dcc993196581c0db960265bb64954239dae5ea8445c17332a6d187351a24f20aa4bddd8cd344087e4597ac29f92023919e6ac735461ca9ecafb154c85c7896239fe3916e5b0a9dde4789e04aa93cd96ce748c2c9a26311cd030e7e5358060905e0fd6a54cd56989c4a925be72a4932c0c1127cfe425e060facae12599f70168c462987f870ec3e4a506a7a1d2797311920b2cd2e93608330927ad299c07fbc70cf0337813053968ef8b62e21f4669ce80ddc42469348171ffd8a415d645c859434a7fb2c71dc44db9edc92c44474e4b0aed8e776b7c7211371415adbe68c58673368e3519a5617530e945a62b2381edc6f076082decb297dcec870adb36ac25c302544d3a88dd46b561a2beb9152bd494a4da33471e5daa651920a28e9fc539cd5a2c950eabc83f74dce92847a19a429f8248fc19baf1b2c394496e4d68b8463f57e1ecac9d74d6ff62d0db1121e0f924ab8eaa939c2ebcf8d91e90d6d5c36fa79d3e49c6a734b50240aaee35d253db946e484d29a3a54cf79be721f1624ad296f65ddcdb848d11a4b0cff9d7df9e2f646d9f4d5e877c3a03cde342d777c2d7cced7c67a1c643b0a175394311ba47ee8959bf1eac5e27284e0e71ea41c856b959f2454ae56dc0ee3baeb54fe54228894ed98d7fb364c086c344a18384bbbc430e4cc2f8fa6d7206947e1fdee00b7d84b034a4d16dd88529309a7c40d97b2b4618c7016b35947bcb742aaef60defe20f3d3571f464cff2f45fdfc3ba85f5959f79a7af13cc8381c5239989cc7257a8c946bc205efd4c72c278a33a89e99071f592746df92f541cfbb01791629734776920bf7e8aa4b2979130c9242822a9e1dd4ddd0e13e1846713f35b783a31c49377cb1f1491ecdfc8d96e3fd12c33f0df1e20572bc5b62f8af21b6ad111ed6273bc8a7c61659a4a17b7256e4c17ea0a67ba26d5ed9ad2d8b939f30b652fc20e63c9c8cb7c303f5c6d67a8eceb6699ea73318fc16748a1d3afffb13e99c2c7f053aff1f0000ffffbd03129920100000"
	tmp.Length = 4128
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
