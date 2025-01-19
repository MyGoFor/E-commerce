package module

import (
	"database/sql/driver"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/json"
)

type stringArray []string

func (sa stringArray) Value() (driver.Value, error) {
	return json.Marshal(sa)
}

func (sa *stringArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan StringArray: %v", value)
	}
	return json.Unmarshal(bytes, sa)
}

type Product struct {
	Id          uint32      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string      `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Picture     string      `protobuf:"bytes,4,opt,name=picture,proto3" json:"picture,omitempty"`
	Price       float32     `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty"`
	Categories  stringArray `protobuf:"bytes,6,rep,name=categories,proto3" json:"categories,omitempty" gorm:"type:json"`
}
