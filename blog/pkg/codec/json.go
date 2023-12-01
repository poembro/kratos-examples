package codec

import (
	"net/http"

	common "github.com/go-kratos/examples/blog/api/common/v1"
	"github.com/go-kratos/kratos/v2/encoding/json"
	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func init() {
	json.MarshalOptions.UseProtoNames = true
	json.MarshalOptions.UseEnumNumbers = true
}

// NormalizeErrorResponse 统一处理错误相应
func NormalizeErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	se := errors.FromError(err)

	ret := &common.HttpResponse{
		Code:    se.Code,
		Message: se.GetMessage(),
		Success: false,
	}

	ret.Data, _ = anypb.New(&common.Empty{})

	body, _ := json.MarshalOptions.Marshal(ret)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(body)
}

// NormalizeResponse 统一处理相应
func NormalizeResponse(w http.ResponseWriter, r *http.Request, data interface{}) error {
	var err error
	var body []byte

	ret := &common.HttpResponse{
		Success: true,
		Message: "请求成功",
	}

	if ret.Data, err = anypb.New(data.(proto.Message)); err != nil {
		return err
	}

	if body, err = json.MarshalOptions.Marshal(ret); err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(body)
	return err
}
