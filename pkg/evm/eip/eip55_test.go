package eip

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestToCheckSumAddress(t *testing.T) {
	cases := []string{
		"0x5afefbf75772bb15af45b9dbca8402f3374062b3",
		//"0x62d17DE1fbDF36597F12F19717C39985A921426e",
		//"0x6F702345360D6D8533d2362eC834bf5f1aB63910",
	}
	for _, c := range cases {
		t.Run(c, func(t *testing.T) {
			res, err := ToCheckSumAddress(strings.ToLower(c))
			assert.Nil(t, err)
			assert.Equal(t, res, c)
		})
	}
}
