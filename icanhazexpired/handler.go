package function

import (
	"fmt"
	"github.com/madflojo/hazexpired"
)

// Handle a serverless request
func Handle(req []byte) string {
	host := fmt.Sprintf("%s", req)
	check, err := hazexpired.Expired(host)
	if err != nil {
		return fmt.Sprintf("ERROR - %s", err)
	}
	if check {
		return "EXPIRED"
	}
	return "OK"
}
