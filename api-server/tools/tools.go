package tools

import (
	"github.com/rs/xid"
)

func GenId() string {
	guid := xid.New()
	return guid.String()
}
