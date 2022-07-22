package do

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// s string, i int, b bool

func Test_Do(t *testing.T) {
	tests := map[string]struct {
		enterString string
		enterInt int
		enterBool bool
		expString string
		expErr string
	}{
		"successBoolTrue" : {
			enterString : "a",
			enterInt : 3,
			enterBool : true,
			expString : "[3A]",
			expErr : "",
		},
		"successBoolFalse" : {
			enterString : "a",
			enterInt : 3,
			enterBool : false,
			expString : "[3a]",
			expErr : "",
		},
		"successInt_13_21_24" : {
			enterString : "a",
			enterInt : 13,
			enterBool : true,
			expString : "A",
			expErr : "",
		},
		"invalidString" : {
			enterString : "text",
			enterInt : 3,
			enterBool : true,
			expString : "",
			expErr : "invalid s",
		},
		"invalidInt" : {
			enterString : "a",
			enterInt : 133,
			enterBool : true,
			expString : "",
			expErr : "invalid s",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := Do(tt.enterString, tt.enterInt, tt.enterBool)
			assert.Equal(t, tt.expString, actual)

			if tt.expErr != "" {
				assert.EqualError(t, err, tt.expErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}