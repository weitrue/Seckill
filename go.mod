module github.com/weitrue/Seckill

go 1.15

require (
	github.com/362228416/go-eth v0.0.0-20220411071848-683ab538efd1
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/Shopify/sarama v1.30.0
	github.com/bsm/sarama-cluster v2.1.15+incompatible
	github.com/coreos/etcd v3.3.25+incompatible
	github.com/ethereum/go-ethereum v1.10.18
	github.com/fsnotify/fsnotify v1.4.9
	github.com/gin-contrib/pprof v1.3.0
	github.com/gin-gonic/gin v1.7.7
	github.com/go-playground/locales v0.14.0
	github.com/go-playground/universal-translator v0.18.0
	github.com/go-playground/validator/v10 v10.10.1
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-stack/stack v1.8.1
	github.com/golang-jwt/jwt/v4 v4.4.2
	github.com/golang/protobuf v1.5.2
	github.com/micro/go-micro v1.18.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/olivere/elastic/v7 v7.0.32
	github.com/pkg/errors v0.9.1
	github.com/shopspring/decimal v1.3.1
	github.com/sirupsen/logrus v1.8.1
	github.com/spaolacci/murmur3 v1.1.0
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.8.0
	github.com/xuri/excelize/v2 v2.6.1
	github.com/zeromicro/go-zero v1.4.0
	go.uber.org/zap v1.21.0
	golang.org/x/crypto v0.0.0-20221012134737-56aed061732a
	golang.org/x/sys v0.0.0-20220728004956-3c1f35247d10
	google.golang.org/grpc v1.48.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gorm.io/driver/mysql v1.0.4
	gorm.io/gorm v1.20.12
)

replace (
	github.com/ethereum/go-ethereum v1.10.6 => github.com/ethereum/go-ethereum v1.10.18
	google.golang.org/grpc v1.47.0 => google.golang.org/grpc v1.29.1
	google.golang.org/grpc v1.48.0 => google.golang.org/grpc v1.29.1
	google.golang.org/grpc v1.49.0 => google.golang.org/grpc v1.29.1
)
