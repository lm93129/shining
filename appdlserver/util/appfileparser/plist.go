package appfileparser

import (
	"appdlserver/model"

	"howett.net/plist"
)

type Assets struct {
	Kind string `xml:"kind" plist:"kind"`
	Url  string `xml:"url" plist:"url"`
}
type Metadata struct {
	BundleIdentifier string `xml:"bundle_identifier" plist:"bundle-identifier"`
	BundleVersion    string `xml:"bundle_version" plist:"bundle-version"`
	Kind             string `xml:"kind" plist:"kind"`
	Title            string `xml:"title" plist:"title"`
}
type Item struct {
	Assets   []Assets `xml:"assets" plist:"assets"`
	Metadata Metadata `xml:"metadata" plist:"metadata"`
}
type sparseBundleHeader struct {
	Items []Item `xml:"items>item" plist:"items"`
}

func EnCodePlist(item model.AppManage, singUrl string) ([]byte, error) {
	data := sparseBundleHeader{Items: []Item{{
		Assets: []Assets{
			{
				Kind: "software-package",
				Url:  singUrl,
			},
		},
		Metadata: Metadata{
			BundleIdentifier: item.BundleId,
			BundleVersion:    item.Version,
			Kind:             "software",
			Title:            item.Name,
		},
	},
	}}

	//buf := &bytes.Buffer{}
	//encoder := plist.NewEncoder(buf)
	//err := encoder.Encode(data)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(buf.String())

	plistdata, err := plist.MarshalIndent(data, 1, "\t")
	return plistdata, err
}
