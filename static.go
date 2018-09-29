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
	tmp.Contents = "1f8b08000000000004ffc4567f6fdc3612fd7f3fc5c3e22ed1c6fbd38e2ff1399bc3c1e70302dca581b3058a3a2e4c4ba315618a1448caeeb6f0772f86a4b46b2771d2a04509246b51c3f9f1de9ba166b3c16c8655251d4aa908b9d15e48ed50b63af7d268077123a412578a70b54141a568951f43d68da29ab4a7025263319fff9d3d35ad25d4465fd3663a88cee594a63182746884f530254c6be1bcd085b0c544c92b2b6c7fe04d898d69712b5d056f508b6b425e09bd26c7cfbea28f8eeea45b502935159c4c4596d8d553a5a0890a3e6de9aa95aac0a5f3c2cb7cba369710a5270b2aa4977a0ddf81318670ecc657c26f81e1f0545f51515081dc341b2ec7f399dc1484db4ae615572a75ae5ab6913a3a21ac4993159e8a84d0446a4fb6b1e4c97668adcc27536c1d85c897b26e94d0fe12de18d51d4a8815463ff510ca922836a8c40dedd84bedbc508a0a94e439450f25af09cefc3311c56b6db0268f490baca5afdaab696eea99bbfe799602b35eb89eb7e61696ba8a426e3d2a37649d34ba0726284be8a2afcc57c43e76cac7adf41573653baa77b24aa13191ba693d0ae1c50c13d37a7eea79648f005710190ee00c526dff17bea29a19178ab964ed796677901bed3cdebdc11207d3c5f3c5e1d1fe3f0e0f0e5f1ebd383a384e6f4f8125b03f7db178b9ff92ff3d3f3c9a3f3f3ceebc9fb00fa1bd83a592ac0d3232a1ee492574a1c8b19898c502a5b178bffacf9bb763fef9eefb55f83d3d3be3546336e135c79c7719bc5fb1259658ecec9c9e9d6189fde3becaff4aeb3c283626c32f3484b562c3aebb8e46c956597c81117e1d00967c6b75da3b9f5f60743cb8ebddae2a822577cf2326b1b185525d4087d31f4e4edfad98de18e45e58f69022a4a88a3cefb6ca6389f38be3c10090253228d259c87b84d758246b047b193160a64b633348bc8ae6a3de8ccb494e9bd655990d4fe358f0b9bc181d87e380dcdb8b7fdef146f82f0291b2dac5e07fe22b9055c23d28b12f2746071734ef338dd162a9e7dba2278b98e31d03329be132da41b74a5d86d952374ae6d28709370d4c31d4b3194e980f1e516b7943ba1f8b301a24f20aa4bddd8c533332e7bd44582e5d5f7aaa1dc38b38cdd85f63a990b9f0ac043ec7d3d4b6149aac5756e06c5a4ae5c9663b47a2c87afee60175665f6119d873a4ca480bef46c2b69ae05412d39daba486500e96e0f3e7f222f865624b647dc259301aa51ce2ba17258824051d2797311920aa825d26ad044e1271e94c9049bcebce0236013053968ef8a22c21f4669ce005630b194d225c7cf423047591b1654a9ad365e17f1638d6efb61db62031101d38582e63f4e4945712964cc08516d8f645ac753683361ea5697531e540e9d064713cb8dbce9e54ba481387872c7f49c870a3b26a424b9a12a269d446ea35cb9595f540a9de24a5a62996b0726dd32849059474fe5398d5a2c950eabc2bef8b9825097dbdd87cdd60c921b224b75e241cabf7735f4ebe6e7ab32f69a86be233aacd0d41912818bedb4a7a728dc809a5357500cd79be64eee390f694b7b2ce12c931b3359618fe947df8e0f646d9f4d9e86fc34038bf342d375a2d7cce83723d0e6a1985511cd5c306c95998acdcfc19ef9e2f2e46087eee40ca51b84878257d3048f17518505d83f0c70141a46cc7bcdfab3f5560a351aa81b3b44b0c43cefcf060680cd27c5078bd335a61b197e6829a2cbac9a026134e89759eb2b4a17b398bd9ac03de5b21d537206f3f8bfcf4d9bf460cff1f05fdfc1ba05f5959f79a7af27891b127131d0cce438a1e56ca9c30e19dfa18e5047106d52373efb3e2c4e81bb23ee87937208f00656ec94e72e11edc3029256f8241524850c5a3f3b1eb75ee836114f7a7c665709423e986ef133ec91391bf4a72bc5e62f8ef219e3c418e574b0c7f1c62db1a61b13ed9413e35b6c8220cddca599107fb019a6e45dbbcb25b5b1627af302d52fc20e63c9c8c43f99e7a636b3d0667db348fc3190cfe0a38c50e9cbffc4e3827cb3f01cedf020000ffff6d280bb8120f0000"
	tmp.Length = 3858
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
