//filteroperation_test.go - тесты для проверки корректности работы конструкторов операций фильтра
package apiv1_test

import (
	"encoding/json"
	"fmt"

	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/export/apiv1"
)

func TestMain(m *testing.M) {
	fmt.Println("START TESTING filteroperation")
	os.Exit(m.Run())
}

func TestNot(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	fmt.Println("TEST Not-operation")
	b, _ := json.Marshal(apiv1.Not(nil))
	assert.Equal(t, len(b), 0)

	b, _ = json.Marshal(apiv1.Not(apiv1.Eq("testField", "testValue")))
	fmt.Println(string(b))
}

func TestGt(m *testing.T) {
	fmt.Println("TEST Gt-operation")
	b, _ := json.Marshal(apiv1.Gt("", ""))
	fmt.Println(string(b))
}

func TestLike(m *testing.T) {
	fmt.Println("TEST Like-operation")
	b, _ := json.Marshal(apiv1.Like("", nil))
	fmt.Println(string(b))
}

func TestBetween(m *testing.T) {
	fmt.Println("TEST Between-operation")
	b, _ := json.Marshal(apiv1.Beetwen("", "", ""))
	fmt.Println(string(b))
}

func TestAnd(m *testing.T) {
	fmt.Println("TEST And-operation")

	b, _ := json.Marshal(apiv1.And(apiv1.Ne("fffff", "vvvv"), apiv1.Eq("gggg", "vvvv3")))
	fmt.Println(string(b))
}

func TestOr(m *testing.T) {
	fmt.Println("TEST Or-operation")

	b, _ := json.Marshal(apiv1.Or())
	fmt.Println(string(b))
}
