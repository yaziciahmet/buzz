package buzz

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_InvalidTypeError(t *testing.T) {
	refType := reflect.TypeOf(1)
	str := ""

	err := fmt.Errorf(invalidTypeMsg, refType, str)
	if err.Error() != "invalid type. expected: int received: string" {
		t.FailNow()
	}
}
