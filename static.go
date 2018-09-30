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
	tmp.Contents = "1f8b08000000000004ffc457ff6fdcc6f1fdfdfe8a81f0f9387791eea893edc68e7c0e02d701dc2f4920ab4051c58556e4f0b8f57297d85d4ab914fadf8b37bbe451b2a2b4698b12b0252e7767e6bd79333b2a8a5951d079a303d5da3095ce46a56da0bab765d4ce0652d74a1b756598ae765471ad7a138f48b79de1966de48ab4a5f5f1f1ffc352d77ba6d6d98fbc5bcd9271bde255f2a00375ca477235b9de5388ca56ca574ba3afbcf2e3817735ed5c4f373a34141db5ea2353d928bbe580f7d8f0274727e1565c6bcb158269d8334c7d660c59e60aa73d5ff5da547419a28aba5c6ddd25a93ab227ae74d4764b7120e388548099d8a8b82706eeb9bde2aae28a4ad7ed0027e24ce92aa69b46970d906a5b9a1e7bb44d4698b66cd9abc8556668a96d64df798eec07b6cedd8321f681c5f3a56e3ba36cbca4e89c190e65c62a673f8ba48c6755eda851d73cd9af6d88ca18aea8e688102319fd9129b82f73a2f06c1d6d39d2b227daead8f457abd2b545f8f863911d432fc0f3adbb21cf0322896d64e59a7dd0ce8ec488b294ad4664b161d898c0a71b1d1be4ca0fa99e44955dd352dbae8f54a9a80a5aba3ee26dcc232c120141cab09033cbd8fea862c32d32ae0c7209ed456477563a1b227dff8e36f474b57eb67efef2e437cf9f3e7ff1f28b974f4ff3d7b7441ba293d517eb17272ff0efd9f397c7cf9e9f0ed6dfc086b23190e79abd171939c1bd6c94ad0c07880959aca8769ede9ffff6ddb747f8f1dd9fcee5e7dbb333849aa291cff0793c44f0fe1c3b6943ebc9cadbb333dad0c9e988f21bed43244e8509fa9525e5bddac1f450d15463d73c7da005fd7d46e439f6dee6b58be30fb4389ddd8e66cf1b26cfe18e455aa6c256c60c0e03bdfdf39bb7df9f23bdc9c91db7b0903d64af8623567b136943171f4e6733225dd39c0cdbb9c4bda0d7b4cebb49f6ebc401325d3b3f274dafd2f6c5b80d70b2d1ae0fcddccbdb51027ca13f2c4ee538913e3c4cbfde6241fe4b44e4a8a61cfc41fd13cc1a15ee411ce124ef0440c763a4c95b827ab107bd5ca7186f414851d065da47b637e6527a4bdb195dea281d6e259902d54541ef82b0bfd5d79c534fdc7671f715be8f0a106e57e9c33c29601f6960532f68b399c4094621108abe67a16c4a56ad4ce029576f200af4c914c6e8d659625536c436fadd51ee0862570dc1ba1a078144476e03724ca9a5c25ee7b9d2a58a9023cea1a5fb9ea5d2472f095cad4d643f9f1cc938213a88e85870e0cdd0462424c0c7d5a49abd30114a96db602a4b52e0d08670fe427f100b50574df331e0b96c5ae418d203dfa317516a7e3dca26533044896d98cc82156164f5e433c27fba70cf841b21ccd57560dcd63529bb3bcaf40ab7a4d39644178e7ec2a0ade6c8420e1ae13e4a1c8a685f937b9240c4400e74059bfbaf230e9d891369ed8b33612d0ab22e52ed7a5bade02897c4727d3abbdd37c0dfb98c2af7c000c4a3b8ae76424a52658868d30f01ff9bd3768e6bc80f5102baa70d1d1c488c78fd45268a82be717e227809fb1e3f4042d8fa4e749f738912b7ce2ea540495595841db8535e45e7c5d0a460fddda602460f3712bf440bf98c7e60cca5c45bfe71bc2bb2494f879b4ccc3cab79918e3eacbc4f4537729d07492d03d508ccd5a4baceecb4dda2c651d3291ba3f6a2cb94e54b2ce733f45d6734576474880f25ad55dd9c6a5b4e53f6688e72f18e0598ef9ffb19029de9e6896d471bb818a9c9f48a1e463b770b39b6ddb8ed610ee572122251834541678cf989472ab13cd293aad5a72db975438e502710a460ee75b4594665e8f5b4a90fe772c447a946cd722dd7bf0037cb2540a22473c14d73fede790cb3e94b90cb01c58eac4a9cd071903d0f60481f86eb07b1205d6bb8c3cba7108a82bec62c0311fde4d817cef250e7f0a4cc8dda85ec108296be63e8159d0ca2c8a14aa47024bd05b7664d8a8cc25f01d29c74a0d06258f618fea55c6652a6acbcd1ecd3ae3b38d19d1224b8be57e6c2a58493ab8a5ee5df70d5cb04264f6272b8506509d4a319cae4974bb928e8bbdf9316d45be72afa12559a0fe3421439a5cbe0ce40d0a8d03c340f60fdb17160f5917761bef8f7a682336edd35936155a12bdc343a72e854c9547bd70a9b0ff5e4bc66a2d76d16bc08c4f3161df9aff31f7e08878bf9eaf3c5ffedbbb3eb7173b72a9618ffb649da0b1930531eb0614f7ca26e8ed58bf58705411b744b6c028fc9497b246de9b38c5d89e4b16473b447581f6b362378a068871be561c18b88eed56cead06c6aa95309847eb64a11853414213e7aa5cdaf60deff2cf3abcfbf5a80feff14f5c7bf82fa73afdb51534f1e0799ae9a9c0e90733f45f791a20981ec417dc27712c29cccc8cc9d3f96de387bcd3e8a9ea70e71b31977c37e59aa706f64cd2145271ba62a474ffcc54915757090c42d9a91bf89682f6fc8aba4c9809a5a4baa84925e6fe8e0eb037af2844a7ab5a183bf1c4cceca33182857ce57f344c3f09498199e9e0835c30367e5aa6cfc7e6fea63e8648729528c29385aca49e9c499da3b77cc6374f65df7389db2e17f41a79ad0f9d3bf48e772f35fa0f31f010000ffff4a270d8de8130000"
	tmp.Length = 5096
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
