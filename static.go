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
	tmp.Contents = "1f8b08000000000004ffc4587f6f1bc711fd5f9f624ab40919913c51b61bbb321d04ae03b868d340568ba2b20badeee6c8ad96bb87dd3dc94ca1ef5ebcd9bde389a695366dd1036c897bb3f3e3cd9b1f54511c15055dac75a05a1ba6d2d9a8b40d54b7b68cdad940ea5669a3ae0dd3f5962aae556be294f4a631bc611bb9226d697172f22b686a5acfb471f686b7f3a3a45ccf799e2ce8408df2915c4daef514a2b295f2d5cce86baf7c7fe16d4d5bd7d29d0e6b8a8e36ea86a95c2bbbe280cf71cd9f5c1db85b71ad2d577066cd9ea1ea4b63c83257b8edf9bad5a6a2ab1055d4e57ce5ae48d5913d71a5a3b62b8a1d185352016ae25ac51d3030cf9b6bae2aaea874cd16e144dc295dc574b7d6e51a916a5b9a1632da26254c2bb6ec55e42a2334d336b26f3c47f61d5a17eea08b6d60b17ca5378d51365e5174ce7497326295b35f4652c6b3aab6b456b73c90d73644650c575473848b918cbe610aee37395178568e561c69d612ad745cb7d7f3d26d8a70f3b1c886c117c4f3bdbb23cf5d44e25b8fca2dfba09ded811166295bf591c53543c7207cbad3718d5cf92ed503afb2699a69dbb4912a15554133d7467ceaf3088d4488206558c039cab1fd41c5356f907165904b702f22bb47a5b321d20f6f69494fe68ba78b672f4e7ffdecc9b3e72fbe7ef1e42cbf7d43b4243a9d7fbd787efa1cff9e3e7b71f2f4d959a7fd3574281b0379aed97ba19193b8676b652bc3016442162baa9da77717bf7dfbfd143ffef8a70bf9f9e6fc1cae266fe4356c9e741ebcbb80242d69313879737e4e4b3a3deba3fc4efb10895361027e654979afb650dd5534d5901aa71734a17f1c11798eadb7f9ecf2e4034dce8eee7bb5176b26cfe181469aa5c256c6740603bdf9cbeb373f5c20bdc9c803b3d0902d64ab86234e5b136949971fce8e8e88744d63326cc7e2f7845ed1224b93c8eb8401325d3b3f264d2f93f8a41743385969d386f5d8cba7690af8527f989cc975227d7c9c7ebdc781fc9780c85e0d31f8bdfa1790352aec85d88793ac13023ae93d4dd652a897bba0678be4e33d00290aba4a72645b63aea4b76c1aa34b1da5c3cd255380ba28e86d10f457fa9673ea89374ddc7e83f73d0304db797a314e0cd8791ad8d4135a2e077e02511084a26f59201b82552b137888d56b90027d32b9d19b759658956b621bfd769a3b82e8559db3aec64544a2236f02724ca9a5425fe3b9d2a58aa023eea1a5fb96a5d27b2b29b85a9bc87e3cb892e304e940a21389039f0c2d854212787f9a58b323265cc974eb54654a4a38b424dcbfd41f4403d855d3b877782c4293ec437a60bbb7224ccd1fa75965728628a10d9599b0428ccc9e7c47f04f03f75cb011c05c5d07c6b4ae49d9ed34c32bd8924e22092e5cfd04415b8d9185ec34dc7d143814d1ae26772001880e1cf00a3a776ffb3874064ea8b52bce146b519075916ad7da6a0e43b924668bb3a3fb5d03fc9dcb51e51e1810714faeebad8092581922daf4a1c0ffeeb41d630cf9ce4b84ee6949a391f8888f3f894451d077ce0f082f6eefe1834808a26f85f739972871ebec4c0a94545589db811be555745e140d0ad63f6c2a40f47829fe8bb7a04f6f07ca5c4abce58ffdacc82a3d1d2f3330e3cce649ba7a98799f92aec73a2f925a16aa3e3057936a1ab3d576851a474da76cf4dc8b2e43968758ce67689bc668aec8e8100f256da39a31d5b61ca6ecd11ce5e2ed0b30cf9ffd0c01ce3479e2a6a1254cf4d06478850fbd9e87851c374d2f761843194e02246ab028e89cb13f710f258e7b7852b5fa24925b37e808762282e4cc5e473bca51197a356ceaddbdecf134d5a8992d64fc4be06636439028c95c70c39cbf731ecb6c7a136438a0d89155f1133c0e22732086f4a21b3ff005e95ac01c3e7c1a4251d0b7d86540a21f1dfbc259eeea1c9694b953db900d82d0d2770cbda4d38e14d955f11486a4b7606ad6a4c8287c0b90e6a403850d96658fe55fcae548ca9495379a7d927a1027ba530a09a6f7ca5cb014777255d1cbfc1b46bd6c60f22424bb812a47801ec59e3abf389ca530f986bc7977a71af1e95699968152bc735d226cc51f391c2a9c7087ca5153bace28017df567850488b7d732d1707a3d3855695dc3f143ee218b325bf10a4ab0d740d144f6bbcf22a3311a54f6213dd000e542503f15afba6aba27368177d0257845491748ffece991303a3d449faaca05df9ba6ae36f4ae3270efa732946a05dc2f0a2916494f6a78522007f3e17c1c1476a342a02510edaaf817929641fdf404429c805f8a2807b197011c1fcc013886d51b99429c0382ee0cc803094adb4e268f9e2da6a48788263bd099789bd12a0aa214d1f1529cbc0732d4b4318c69847622f3c2b91b1a4d45704a23f9c981e6f3f77624467201c09161013cd87dd72aac0fadbe387f6cf39ddff0368c27ffd9027cce1b77cb64585508e86ead238746954cb5771b61c1a1f5239f99e8f526534072ea7985e5e36fe3f7efc3f1643cff6af2cbdd22e25a2ca91b154b7cd359a52e9e6a2dd504047a8a64ecc638bd5c7c9810dae07e053cc05752b7fb2ed84fa7eced14f4eec7538ee0c07cea962770748f931db1f7c6535a46d8d43292c44ffaec408217323b05f8e895363f0379ff59e4e75f7d3301fcff2de84f7e06f4175e6f7a4e7df17890a9c9e474009cfd14ed478a9c00ec8e7d827722c2984c8fcc83bf0bbc76f696fdb0ab258358e28cbb633f2b55d8fb769624e6d189c090e5e85c32380e7104a9c75f0b500723192487bb18844a4add49be8be126fa107a5b49af9634fa76445f7c4125bd5cd2e8af23da95863c9d8272ee7c354e30744f89f5f8c9a940d33d3056cecbb5dfc976dd4f7a5fb62f642ee5a6ccf00c6d1e11e97bdb6370b64df3389c22f0ff80530de0fcf1df8473b6fc1fc0f9cf000000ffff918f8d94d3160000"
	tmp.Length = 5843
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
