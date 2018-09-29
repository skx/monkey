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
	tmp.Contents = "1f8b08000000000004ffc4567f6fdc3612fd7f3fc5c3e22ed1c6fbd38e2ff1399bc3c1e70302dca581b3058a3a2e4c4ba315618a1448caeeb6f0772f86a4b46b2771d2a04509246b51c399376fde0c359b0d6633ac2ae9504a45c88df6426a87b2d5b997463b881b2195b85284ab0d0a2a45abfc18b26e14d5a43d15901a8bf9fcefeca9692da136fa9a36d341742ea7348d11a44323ac8729615a0be7852e842d264a5e5961fb036f4a6c4c8b5be92a78835a5c13f24ae835397ef6157d7474076e41a5d45430988a2cb1aba74a4113157cdad2552b55814be78597f9746d2e214a4f1654482ff51abe23630ce1d88daf84df12c3e1a9bea2a2a002b969369c8ee733b92908b795cc2bce54ea5cb56c237574425893262b3c1589a189d49e6c63c993edd85a994f426c1d85c897b26e94d0fe12de18d51d4a8c15463ff510ca922836a8c40dedd84bedbc508a0a94e419a28792d70467fe990ac56b6db0268f490baca5afdaab696eea99bbfe79960207bdf05a9bcee523969cf95b730b4b5dee218b9ebf1bb24e1add5318342874d173e02b621f3b44e156fa8aab6a3b51ece04f203191ba693d0ae1c50c13d37a7eea2bce1e01ce206a21d038482cfc5ff88a6ad686505c7556a9671d0c72a39dc7bb3758e260ba78be383cdaffc7e1c1e1cba3174707c7e9ed29b004f6a72f162ff75ff2bfe78747f3e787c79df713f621b477b05492b5417026e43da9842e1439961dd7bb40692cdeaffef3e6ed987fbefb7e157e4fcfce186a44135e73cc7987e0fd8a2db1c46267e7f4ec0c4bec1ff759fe575ae741b185997ea121ac151b76ddf53e4ab6cae20b8cf0eb00b0e45babd3def9fc02a3e3c15def7655112cb97b1e3189234028d5057438fde1e4f4dd8acb1b83dc0bcb1e5284145591e7dd56792c717e713c1800b24406453a0bb847788d45b246b0979103ae74696c068957d17cd49b713ac969d3ba2ab3e1691c133e9717a3e3701c907b7bf1cf3bde08ff452212aa5d0efe27be825925dc8314fb7462747042f31e698c16533ddf263d59448c774cc86c86cb6807dd2a7519a650dd28994b1f66e134548aa99ecd70c2f5e061b69637a4fb010aa34122af40dadbcd383523d7bc9708cba5eb4b4fb5637a11e71efb6b2c1532179e95c0e778eeda964293f5ca0a359b965279b2d9ce9128b2be7ef3c03a575f6119aae74895b12cbc1b0bb6d504434995ee5c25358474b0049f3f9717c12f17b644d603ce82d1286188eb5e94209214749c5c4630405405bb4c5a093549854b67824ce2ad7816b8098499b274c4576a09a137e3442f985bc86812e9e2a31f31a88b8c2d136886cbc2ff2c71acdf6d3b6c4962223a72b05cc6e8c929af242c99880b2db0ed8b98eb6c066d3c4ad3ea62ca81d2a1c9e27870b79d3d297591260e0f59fee690e1ee65d58496342544d3a88dd46b962b2beb8152bd494a4d532c71e5daa651920a28e9fca738ab4593a1d47997de17394b12fa7ab1f9bac1924364496ebd483856efe7be9c7cddf4665fd250d7c467549b1b822251307db795f4e41a91134a6bea409af37cc9dce721ed296f659da52247646b2c31fc29fbf0c1ed8db2e9b3d1df86a1e0fcd2b4dc68b5f0390fcaf538a865144671540f1b246761b272f367bc7bbeb81821f8b9032947e122e195f4c124c5d76140750dc21f070491d08e79bf577fcac046a39403a3b44b0c03667e78303406693e28bcde19adb0d84b73414d16dd6450930943629d279436742fa398cd3ae2bd15527d03f3f6b3cc4f9ffd6bc4f4ff51d4cfbf81fa959575afa9278f27197b329583c97958a28799724db8e09dfa98e5447106d53373efb3e2c4e81bb23ee87937208f00656ec94e72e11edc30099237c1202924a8e2d1f9d8f53af7c1308afb53e33238ca9174c3f7099fe489c85f25395e2f31fcf7104f9e20c7ab25863f0eb16d8db0589fec209f1a5b6491866ee5acc883fd404db7a26d5ed9ad2d8b93579816297e10731e4ec6a17c4fbdb1b51ea3b36d9ac7e90c067f059d6287ce5f7e279d93e59f40e76f010000ffff9c6e300d3c0f0000"
	tmp.Length = 3900
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
