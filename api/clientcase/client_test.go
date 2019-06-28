package clientcase_test

import (
	"testing"

	cc "github.com/yeqown/micro-server-demo/api/clientcase"
	protogen "github.com/yeqown/micro-server-demo/api/protogen"
)

func Test_Echo(t *testing.T) {
	if err := cc.Initial("127.0.0.1:8080"); err != nil {
		t.Errorf("could not Initial: %v", err)
		t.FailNow()
	}

	rpcResp, err := cc.Echo(
		&protogen.FooForm{Foo: "bar"})
	if err != nil {
		t.Errorf("could not cc.Echo: %v", err)
		t.FailNow()
	}

	t.Log(rpcResp)
}
