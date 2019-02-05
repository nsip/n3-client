package xjy

import (
	"io/ioutil"
	"testing"
)

func TestJstr2Y(t *testing.T) {

	ystr := Jstr2Y(`{"actor": {
		"name": "Team PB",
		"mbox": "mailto:teampb@example.com",
		"member": [
			{
				"name": "Andrew Downes",
				"account": {
					"homePage": "http://www.example.com",
					"name": "13936749"
				},
				"objectType": "Agent"
			},
			{
				"name": "Toby Nichols",
				"openid": "http://toby.openid.example.org/",
				"objectType": "Agent"
			},
			{
				"name": "Ena Hills",
				"mbox_sha1sum": "ebd31e95054c018b10727ccffd2ef2ec3a016ee9",
				"objectType": "Agent"
			}
		],
		"objectType": "Group"
	}
}`)

	ioutil.WriteFile("test.yaml", []byte(ystr), 0666)
}
