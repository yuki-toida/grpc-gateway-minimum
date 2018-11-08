package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	gw "github.com/yuki-toida/grpc-gateway-sample/proto"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(matcher), runtime.WithForwardResponseOption(filter))
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := gw.RegisterEchoServiceHandlerFromEndpoint(ctx, mux, "localhost:9090", opts); err != nil {
		panic(err)
	}

	if err := gw.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, "localhost:9091", opts); err != nil {
		panic(err)
	}

	if err := http.ListenAndServe(":8080", mux); err != nil {
		glog.Fatal(err)
	}

}

func matcher(headerName string) (string, bool) {
	ok := headerName != "Ignore"
	fmt.Printf("%v %s\n", ok, headerName)
	return headerName, ok
}

func filter(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	w.Header().Set("X-Filter", "FilterValue")
	return nil
}
