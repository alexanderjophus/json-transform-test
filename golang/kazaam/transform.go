package main

import "github.com/qntfy/kazaam"

// {
// 	"fixture": "Antonio Schmitt+Alison Crist",
// 	"startDate": "2019-06-13T20:37:43.160Z",
// 	"sport": "high fiving",
// 	"markets": [
// 	   {
// 		  "id": 998,
// 		  "type": "slowest-sports-person",
// 		  "options": [
// 			 {
// 				"id": 750,
// 				"name": "Antonio Schmitt",
// 				"odds": 1.4
// 			 },
// 			 {
// 				"id": 593,
// 				"name": "Alison Crist",
// 				"odds": "8/9"
// 			 }
// 		  ]
// 	   }
// 	],
// 	"uuid": "103"
// }

var (
	spec = `
	[
		{
			"operation": "shift",
			"inplace" : true,
			"spec": {
			  "startDate": "timings.start"
			}
		},
		{
			"operation": "delete",
			"spec": {
			  "paths": ["timings"]
			}
		}
	]`

	kazaamTransform *kazaam.Kazaam
)

func setupKzm() error {
	var err error
	kazaamTransform, err = kazaam.NewKazaam(spec)
	return err
}

func transform(in string) (string, error) {
	return kazaamTransform.TransformJSONStringToString(in)
}
