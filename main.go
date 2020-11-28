package main

import (
	af "grpcPanel/afra_instrument_v1"
)

func main() {

	//af.GetWithCodeList([]string{"8516759016718718", "13082825954602280", "48711126865295010"})
	af.GetWithMarketCode(3)
}