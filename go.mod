module github.com/sheacloud/cloud-inventory

go 1.17

require (
	github.com/aws/aws-lambda-go v1.28.0
	github.com/aws/aws-sdk-go-v2 v1.13.0
	github.com/aws/aws-sdk-go-v2/config v1.8.1
	github.com/aws/aws-sdk-go-v2/credentials v1.4.1
	github.com/aws/aws-sdk-go-v2/service/athena v1.11.0
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.12.0
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.27.0
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.17.0
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.11.0
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.15.0
	github.com/aws/aws-sdk-go-v2/service/iam v1.15.0
	github.com/aws/aws-sdk-go-v2/service/lambda v1.16.0
	github.com/aws/aws-sdk-go-v2/service/rds v1.15.0
	github.com/aws/aws-sdk-go-v2/service/redshift v1.18.0
	github.com/aws/aws-sdk-go-v2/service/route53 v1.16.0
	github.com/aws/aws-sdk-go-v2/service/s3 v1.15.1
	github.com/aws/aws-sdk-go-v2/service/sts v1.7.0
	github.com/aws/smithy-go v1.10.0
	github.com/awslabs/aws-lambda-go-api-proxy v0.12.0
	github.com/fatih/structtag v1.2.0
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.7
	github.com/google/uuid v1.3.0
	github.com/hashicorp/hcl/v2 v2.11.1
	github.com/jinzhu/copier v0.3.4
	github.com/r3labs/diff/v2 v2.14.5
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.10.1
	github.com/swaggo/files v0.0.0-20210815190702-a29dd2bc99b2
	github.com/swaggo/gin-swagger v1.4.0
	github.com/swaggo/swag v1.7.8
	github.com/xitongsys/parquet-go v1.6.2
	github.com/xitongsys/parquet-go-source v0.0.0-20211228015320-b4f792c43cd0
	golang.org/x/tools v0.1.7
)

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/agext/levenshtein v1.2.1 // indirect
	github.com/apache/arrow/go/arrow v0.0.0-20200730104253-651201b0f516 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.5.0 // indirect
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.3.2 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.4 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.2.0 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.2.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.6.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.4.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.6.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.7.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.4.0 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/spec v0.20.4 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mitchellh/go-wordwrap v0.0.0-20150314170334-ad45545899c7 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/pierrec/lz4/v4 v4.1.8 // indirect
	github.com/spf13/afero v1.6.0 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/ugorji/go/codec v1.1.7 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	github.com/zclconf/go-cty v1.8.0 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/mod v0.5.0 // indirect
	golang.org/x/net v0.0.0-20210813160813-60bc85c4be6d // indirect
	golang.org/x/sys v0.0.0-20211210111614-af8b64212486 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/ini.v1 v1.66.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

require (
	github.com/apache/thrift v0.15.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.12.0
	github.com/aws/aws-sdk-go-v2/service/ecs v1.15.0
	github.com/aws/aws-sdk-go-v2/service/efs v1.12.0
	github.com/golang/snappy v0.0.4 // indirect
	github.com/klauspost/compress v1.13.5 // indirect
)
