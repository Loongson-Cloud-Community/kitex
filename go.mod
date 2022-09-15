module github.com/cloudwego/kitex

go 1.13

require (
	github.com/apache/thrift v0.13.0
	github.com/bytedance/gopkg v0.0.0-20220531084716-665b4f21126f
	github.com/choleraehyq/pid v0.0.15
	github.com/cloudwego/fastpb v0.0.2
	github.com/cloudwego/frugal v0.1.3
	github.com/cloudwego/netpoll v0.2.6
	github.com/cloudwego/thriftgo v0.2.1
	github.com/golang/mock v1.6.0
	github.com/jhump/protoreflect v1.8.2
	github.com/json-iterator/go v1.1.12
	github.com/tidwall/gjson v1.9.3
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20220829200755-d48e67d00261
	golang.org/x/tools v0.1.1
	google.golang.org/genproto v0.0.0-20210513213006-bf773b8c8384
	google.golang.org/protobuf v1.28.0
	gopkg.in/yaml.v3 v3.0.1
)

replace github.com/bytedance/gopkg v0.0.0-20220531084716-665b4f21126f => github.com/Loongson-Cloud-Community/gopkg v0.0.0-20220918043155-a7bf4087dcd6

replace github.com/cloudwego/netpoll v0.2.6 => github.com/Loongson-Cloud-Community/netpoll v0.2.7-0.20220915043215-f0f548df0481

replace github.com/choleraehyq/pid v0.0.15 => github.com/Loongson-Cloud-Community/pid v0.0.16-0.20220915082924-15274255b2c4
