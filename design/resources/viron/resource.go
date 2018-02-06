package viron

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("Viron", func() {
	Origin("*", func() {
		Methods("GET")
	})
	Files("/viron_authtype", "staticfile/viron_authtype.json")
	Files("/viron", "staticfile/viron.json")
})
