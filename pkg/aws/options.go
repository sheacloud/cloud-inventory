package aws

import (
	"time"

	"github.com/jinzhu/copier"
)

var Int64 int64 = 0

var CopyOption = copier.Option{
	IgnoreEmpty: true,
	DeepCopy:    true,
	Converters: []copier.TypeConverter{
		{
			SrcType: time.Time{},
			DstType: Int64,
			Fn: func(src interface{}) (interface{}, error) {
				return src.(time.Time).UTC().UnixMilli(), nil
			},
		},
	},
}
