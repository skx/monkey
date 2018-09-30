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
	tmp.Contents = "1f8b08000000000004ffc4577f6fe4b611fd7f3fc560d15eb4f1eecaeb3b3777f1e982c07500f74712f85ca0a8cf856969b4624d910249d9d914feeec50c29aebcde386dda220292b3287266de9b3733dc3c9fe4395c36d2412d154269b417523ba87b5d7a69b403712fa412b70ae1760315d6a2577e0eb2ed14b6a83d562035ac0e0f7f4b96bade22b446dfe1663909c6e51297c18374d009ebc1d4607a0bce0b5d095b2d94bcb5c2a603e7356c4c0f0fd235e00db4e20ea16c845ea3a377dfe0b3a3a3702baca5c68a8269d02299fa4c29d088159db678db4b55c18df3c2cb72b93637206a8f16b0925eea35f8818c390847667c23fc9618728fed2d561556509a6e43703c9d294d85f0d0c8b221a45297aaa73d520723086bd46885c72a32b490daa3ed2c7ab4035b97666f88bd43f67c23db4e09ed6fc01ba3864391b1cae8cf3c086551541b68c43d8ef64bedbc500a2ba8d153881e94bc4370e6cb98287ad606d6e861d103aca56ffadb6569dadcddfd9047c7a417c2f3ad79008b03228e2db1728fd649a31331ac2ca1ab84cc37483646f0e141fa86726587548fa28aae612175d77ba88417392c4cefe92de5912c0210829061266712b1fd59f8065bcab850944bd29ea7ec4e4aa39d87efcfa180d7cbd59bd5f1bba3df1dbf3e7efbee8b77af4fe2d7338002e068f9c5eaedd15bfaefcdf1bbc337c72783f553b221b47760b1466b594686712f1aa12b858ec44459aca036163e5efefefcdb39fdf3dd5f2ef9dfb38b0b0a3544c39fc9e7e110c1c74bda0905ac462b67171750c0d14942f98db4ce0386c224fa850661add890e9a1a2a1a65d59f80033f8e704c0a2efad8e6b5787d7303b993c26b3970d8245f7c4222c42610ba506870ecefe7a7af6fd25a5373879e2962c440fd1ab424fabbdf250c0d5f5c96402206bc840a1ce38ee197c8055dc0dbc5f060e28d3b5b11948781fb6cfd23682138d76bd6b32cb6ff300f84a5ecf4ef838803c38087f3ed202ff2f1011a31a73f027f16f30ab84db8198e004ef40800e53a4c15b807ab505bd5885181f89903c879bb00f74afd40df796b653b2949e3bdc92334554e7399c3b667f2def31a61eb0edfce62bfa9e14c0dc2ec3872c28601ba94355cfa028467112a32410f0b647a66c4c562d94c33157a7240aea93218ce4d66840513680dadbcd3c7604b62b86604d4d070989f4d83aca3184964af63a8b952c852739d2396ae9b647aef4e42580aba5f268b3d19188934447223a641cf4a6a0600931f0b41a54b315268512e536988a92643850009dbf92d76c81d455439602ce78d32cc6101ef29dbcb052e3eb3c9a0cc10004b6c964142c0b23aa279e61fec3c0bd606e983053d70e695ad720f4661ee9656e41862d812e3afa8c415d659485183485fb22715444db9adc9244440ce490aec8e6f66bc22123712cad6d7106ac790eda78a84dafab25398a25b1589d4c1eb70df00f26a28a3dd011e224aedb0d931254e93cb5e97dc0ff61a4ce680cd9214a826ea180e99463a4d79f6522cfe11b634782e7b077f82124405bcf59f7319754e2dae805172888aae2b01d76c20a6f2c1b1a15ac7dda5488d18382e3e768493ec90f193321f11a7f48b3229ab470504462b2a8e65938ba5f79cf4597b88e1749c917aa04ccd420ba4e6da45e538d534d876c24ed7913298b432ce6d3f55da72456a0a4f3fb92d68a2e835a97e394bd98a358bca900e3fcd9cd10d119268f6f3b28c845a226d2cb7a48769e16b26fbbb46d3f873c9c9848aac13c870ba4fb13262a6939d113aad5862db175931c499d842004b3d3d1261195820fe3a63e9c8b11cf438daac58ac73f03578b0581a4928c0537cef94763e9321bbe381e0e54ec94558e9374ec78cf1e0ce1c3307e28164ad78adcd1cb7308790e5fd35d8644f4a3419b1b8d439d9327a11ec4c645872468ee3b0adec3d1208a182a474a8eb8b7d0d4ac418012f42b809b9374e05aba2c5bbafc73b94cb84c515825d1865d4f7052770a90c8f54e99e739b3d9f5de65303d356d27a8f9c034722eafe73005b1a65f5e7ebbba58d1fa273d1d24c4886261c2fbf817dd16f812c74f48c630937989b247fd220c8f01f3777f04c9bcad8da9e04baaf37896462a0b328c9327578a46b866df8d82d65fba502cef70e3b2d97f77afb8c0d6dc23281415f59587467a749d28116a6b5acec7beae1ed794b7b28d25c312b3b8a69efef7ecd3277730cb969fcf7eb3edefa6a7d9df0a5fd205721d1235e32b6a48036dd8f21ea8cb68f56a753d0352173c022a872937610f672d7ce68b5b2039157d8c764eeba9ea23823d653fcca4fd25c332dca9fad0e351d55ce91c08fc649d5314dc9298786f8554bf8079fb93cc2f3fff6a46f4ffafa83ffc05d45f5ad9264dbd7a19641856311d44ce6e8a7691521b23b207f531df410819a8c4cc939f5ba746dfa3f5ace7b1439a8dca3ca05d94c2ed5c7a6348def086b1caa9abfeec5d97ea601ac4cd9ae15f55b09537c9ab84d1153774965009257c2860faf5145ebd8212de1730fddb7474969fc140b934b6ca020dc353d2ade3f51153333ce4ac5c968dddee8d6d2c8cd1e89fc55cf249ee6b91da2753ea253afbae7b994edef06bd0294674fef81fd2b928fe0f74fe2b0000ffffc429159e2a140000"
	tmp.Length = 5162
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
