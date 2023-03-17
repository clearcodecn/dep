package dep

import (
	_ "github.com/clearcodecn/dep/db"
	_ "github.com/clearcodecn/dep/leveldb"
	_ "github.com/clearcodecn/dep/log"
	_ "github.com/clearcodecn/dep/resty"
	_ "github.com/clearcodecn/dep/yaml"
)

func Register() {}
