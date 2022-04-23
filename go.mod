module github.com/sheacloud/cloud-inventory

go 1.17

require (
	github.com/amzn/ion-go v1.1.3
	github.com/aws/aws-lambda-go v1.30.0
	github.com/aws/aws-sdk-go-v2 v1.16.2
	github.com/aws/aws-sdk-go-v2/config v1.15.3
	github.com/aws/aws-sdk-go-v2/credentials v1.11.2
	github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue v1.9.0
	github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression v1.4.6
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.15.3
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.12.3
	github.com/aws/aws-sdk-go-v2/service/athena v1.15.0
	github.com/aws/aws-sdk-go-v2/service/backup v1.15.4
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.15.5
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.15.4
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.36.0
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.20.5
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.14.3
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.18.3
	github.com/aws/aws-sdk-go-v2/service/iam v1.18.3
	github.com/aws/aws-sdk-go-v2/service/lambda v1.22.0
	github.com/aws/aws-sdk-go-v2/service/rds v1.20.0
	github.com/aws/aws-sdk-go-v2/service/redshift v1.23.0
	github.com/aws/aws-sdk-go-v2/service/route53 v1.20.3
	github.com/aws/aws-sdk-go-v2/service/s3 v1.26.5
	github.com/aws/aws-sdk-go-v2/service/sns v1.17.4
	github.com/aws/aws-sdk-go-v2/service/sqs v1.18.3
	github.com/aws/aws-sdk-go-v2/service/storagegateway v1.17.0
	github.com/aws/aws-sdk-go-v2/service/sts v1.16.3
	github.com/aws/smithy-go v1.11.2
	github.com/awslabs/aws-lambda-go-api-proxy v0.13.2
	github.com/fatih/structtag v1.2.0
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.7
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/google/uuid v1.3.0
	github.com/hashicorp/hcl/v2 v2.12.0
	github.com/jinzhu/copier v0.3.5
	github.com/r3labs/diff/v2 v2.15.1
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.11.0
	github.com/swaggo/files v0.0.0-20210815190702-a29dd2bc99b2
	github.com/swaggo/gin-swagger v1.4.2
	github.com/swaggo/swag v1.8.1
	github.com/xitongsys/parquet-go v1.6.2
	github.com/xitongsys/parquet-go-source v0.0.0-20220315005136-aec0fe3e777c
	go.mongodb.org/mongo-driver v1.9.0
	golang.org/x/tools v0.1.10
)

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/apache/arrow/go/arrow v0.0.0-20211112161151-bc219186db40 // indirect
	github.com/apache/thrift v0.16.0 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.1 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.12.3 // indirect
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.11.5 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.9 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.3 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.10 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.0.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.13.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.1.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.7.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.13.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.11.3 // indirect
	github.com/fsnotify/fsnotify v1.5.3 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/spec v0.20.5 // indirect
	github.com/go-openapi/swag v0.21.1 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.10.1 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.7 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.0-beta.8 // indirect
	github.com/pierrec/lz4/v4 v4.1.14 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/spf13/afero v1.8.2 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.1 // indirect
	github.com/xdg-go/stringprep v1.0.3 // indirect
	github.com/youmark/pkcs8 v0.0.0-20201027041543-1326539a0a0a // indirect
	github.com/zclconf/go-cty v1.10.0 // indirect
	golang.org/x/crypto v0.0.0-20220411220226-7b82a4e95df4 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220106191415-9b9b3d81d5e3 // indirect
	golang.org/x/net v0.0.0-20220421235706-1d1ef9303861 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20220422013727-9388b58f7150 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/xerrors v0.0.0-20220411194840-2f41105eb62f // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/ini.v1 v1.66.4 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

require (
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.15.3
	github.com/aws/aws-sdk-go-v2/service/ecs v1.18.5
	github.com/aws/aws-sdk-go-v2/service/efs v1.17.0
	github.com/golang/snappy v0.0.4 // indirect
	github.com/klauspost/compress v1.15.1 // indirect
)
