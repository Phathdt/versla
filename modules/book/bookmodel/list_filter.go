package bookmodel

import "versla/sdk/sdkcm"

type ListFilter struct {
	Author string `json:"author" form:"author"`
}

type ListParams struct {
	sdkcm.Paging `json:",inline"`
	*ListFilter  `json:",inline"`
}
