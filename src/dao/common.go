package dao

import (
	"assetLibary/configs"
	"assetLibary/services"
	pbStruct "github.com/golang/protobuf/ptypes/struct"
	"gitlab.zixel.cn/go/framework/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
	"reflect"
)

var log = logger.Get()

// CreateError : This function creates a new error
func CreateError(code codes.Code, errorCode int, args ...map[string]interface{}) error {
	message := configs.ErrorCodes[errorCode]
	err := status.Newf(
		code,
		message,
	)

	var _ *structpb.Struct
	if args != nil {
		_ = ConvertMapToStruct(args[0])
	}

	err, wde := err.WithDetails(
		&services.ApiExceptionRpc{
			ErrorDetail: &services.ErrorDetailRpc{
				Code:    int64(errorCode),
				Message: message,
				ServiceInfo: &services.ServiceInfoRpc{
					Name: "Session V2",
					Uuid: "Development",
				},
			},
		})
	if wde != nil {
		return wde
	}
	return err.Err()
}

// ConvertMapToStruct : This function converts map[string]string to struct
func ConvertMapToStruct(stringMap map[string]interface{}) *pbStruct.Struct {
	fields := make(map[string]*pbStruct.Value, len(stringMap))
	for k, v := range stringMap {
		if reflect.TypeOf(v).Kind() == reflect.Map {
			fields[k] = &pbStruct.Value{
				Kind: &pbStruct.Value_StructValue{
					StructValue: ConvertMapToStruct(v.(map[string]interface{})),
				},
			}
			continue
		} else if reflect.TypeOf(v).Kind() == reflect.Slice {
			fields[k] = &pbStruct.Value{
				Kind: &pbStruct.Value_ListValue{
					ListValue: &pbStruct.ListValue{
						Values: []*pbStruct.Value{},
					},
				},
			}
			for _, val := range v.([]interface{}) {
				fields[k].GetListValue().Values = append(fields[k].GetListValue().Values, &pbStruct.Value{
					Kind: &pbStruct.Value_StringValue{
						StringValue: val.(string),
					},
				})
			}
			continue
		} else if reflect.TypeOf(v).Kind() == reflect.Bool {
			fields[k] = &pbStruct.Value{
				Kind: &pbStruct.Value_BoolValue{
					BoolValue: v.(bool),
				},
			}
			continue
		} else if reflect.TypeOf(v).Kind() == reflect.Int {
			fields[k] = &pbStruct.Value{
				Kind: &pbStruct.Value_NumberValue{
					NumberValue: float64(v.(int)),
				},
			}
			continue
		} else if reflect.TypeOf(v).Kind() == reflect.Float64 {
			fields[k] = &pbStruct.Value{
				Kind: &pbStruct.Value_NumberValue{
					NumberValue: v.(float64),
				},
			}
			continue
		} else if reflect.TypeOf(v).Kind() == reflect.String {
			fields[k] = &pbStruct.Value{
				Kind: &pbStruct.Value_StringValue{
					StringValue: v.(string),
				},
			}
			continue
		}
	}
	return &pbStruct.Struct{
		Fields: fields,
	}
}
