module github.com/temphia/temphia

go 1.20

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5

require (
	github.com/alecthomas/kong v0.8.1
	github.com/alecthomas/repr v0.3.0
	github.com/brianvoe/gofakeit/v6 v6.24.0
	github.com/bwmarrin/snowflake v0.3.0
	github.com/dop251/goja v0.0.0-20231027120936-b396bb4c349d
	github.com/dop251/goja_nodejs v0.0.0-20231022114343-5c1f9037c9ab
	github.com/flosch/go-humanize v0.0.0-20140728123800-3ba51eabe506
	github.com/gin-gonic/gin v1.9.1
	github.com/go-git/go-git/v5 v5.10.0
	github.com/gobwas/ws v1.3.1
	github.com/golang-jwt/jwt/v4 v4.5.0
	github.com/graphql-go/graphql v0.8.1
	github.com/hako/branca v0.0.0-20200807062402-6052ac720505
	github.com/hashicorp/go-multierror v1.1.1
	github.com/hashicorp/golang-lru v1.0.2
	github.com/jaevor/go-nanoid v1.3.0
	github.com/joho/godotenv v1.5.1
	github.com/k0kubun/pp v3.0.1+incompatible
	github.com/lib/pq v1.10.9
	github.com/mattn/go-sqlite3 v1.14.18
	github.com/mitchellh/mapstructure v1.5.0
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/postfinance/single v0.0.2
	github.com/rqlite/sql v0.0.0-20221103124402-8f9ff0ceb8f0
	github.com/rs/xid v1.5.0
	github.com/rs/zerolog v1.31.0
	github.com/siddontang/go v0.0.0-20180604090527-bdc77568d726
	github.com/thoas/go-funk v0.9.3
	github.com/tidwall/buntdb v1.3.0
	github.com/tidwall/gjson v1.17.0
	github.com/tidwall/pretty v1.2.1
	github.com/twpayne/go-geom v1.5.3
	github.com/ugorji/go/codec v1.2.11
	github.com/upper/db/v4 v4.7.0
	github.com/webview/webview_go v0.0.0-20230901181450-5a14030a9070
	github.com/ztrue/tracerr v0.4.0
	golang.org/x/crypto v0.15.0
	golang.org/x/oauth2 v0.14.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
	gopkg.in/yaml.v2 v2.4.0
)

require (
	dario.cat/mergo v1.0.0 // indirect
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20230828082145-3c4c8a2d2371 // indirect
	github.com/acomagu/bufpipe v1.0.4 // indirect
	github.com/bytedance/sonic v1.9.1 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/cloudflare/circl v1.3.3 // indirect
	github.com/cyphar/filepath-securejoin v0.2.4 // indirect
	github.com/dlclark/regexp2 v1.10.0 // indirect
	github.com/eknkc/basex v1.0.0 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/fatih/color v1.10.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-git/gcfg v1.5.1-0.20230307220236-3a3c6141e376 // indirect
	github.com/go-git/go-billy/v5 v5.5.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.14.0 // indirect
	github.com/go-sourcemap/sourcemap v2.1.3+incompatible // indirect
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/pprof v0.0.0-20230926050212-f7f687d19a98 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.14.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.2 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgtype v1.14.0 // indirect
	github.com/jackc/pgx/v4 v4.18.1 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/kevinburke/ssh_config v1.2.0 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.8 // indirect
	github.com/pjbgf/sha1cd v0.3.0 // indirect
	github.com/segmentio/fasthash v1.0.3 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/skeema/knownhosts v1.2.0 // indirect
	github.com/tidwall/btree v1.4.2 // indirect
	github.com/tidwall/grect v0.1.4 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/rtred v0.1.2 // indirect
	github.com/tidwall/tinyqueue v0.1.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/xanzy/ssh-agent v0.3.3 // indirect
	golang.org/x/arch v0.3.0 // indirect
	golang.org/x/mod v0.12.0 // indirect
	golang.org/x/net v0.18.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/tools v0.13.0 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
