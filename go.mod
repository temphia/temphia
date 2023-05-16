module github.com/temphia/temphia

go 1.20

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5

require (
	github.com/Masterminds/sprig/v3 v3.2.3
	github.com/alecthomas/kong v0.7.1
	github.com/alecthomas/repr v0.1.0
	github.com/antonmedv/expr v1.9.0
	github.com/brianvoe/gofakeit/v6 v6.18.0
	github.com/bwmarrin/snowflake v0.3.0
	github.com/dop251/goja v0.0.0-20220815083517-0c74f9139fd6
	github.com/fergusstrange/embedded-postgres v1.19.0
	github.com/flosch/go-humanize v0.0.0-20140728123800-3ba51eabe506
	github.com/gin-gonic/gin v1.8.1
	github.com/go-git/go-git/v5 v5.6.1
	github.com/gobwas/ws v1.1.0
	github.com/goccy/go-yaml v1.9.5
	github.com/golang-jwt/jwt/v4 v4.4.2
	github.com/hako/branca v0.0.0-20200807062402-6052ac720505
	github.com/hashicorp/golang-lru v0.5.4
	github.com/joho/godotenv v1.5.1
	github.com/k0kubun/pp v3.0.1+incompatible
	github.com/lib/pq v1.10.7
	github.com/mattn/go-sqlite3 v1.14.16
	github.com/mitchellh/mapstructure v1.5.0
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/rqlite/sql v0.0.0-20221103124402-8f9ff0ceb8f0
	github.com/rs/xid v1.4.0
	github.com/rs/zerolog v1.27.0
	github.com/siddontang/go v0.0.0-20180604090527-bdc77568d726
	github.com/tetratelabs/wazero v0.0.0-20220812081006-d7d18a5519e6
	github.com/thoas/go-funk v0.9.2
	github.com/tidwall/buntdb v1.2.9
	github.com/tidwall/gjson v1.14.4
	github.com/tidwall/pretty v1.2.1
	github.com/twpayne/go-geom v1.5.2
	github.com/ugorji/go/codec v1.2.7
	github.com/upper/db/v4 v4.5.4
	github.com/yuin/goldmark v1.5.3
	github.com/ztrue/tracerr v0.3.0
	gitlab.com/mr_balloon/golib v0.0.0-20210813185029-b13a7945a495
	go.etcd.io/bbolt v1.3.6
	golang.org/x/crypto v0.6.0
	golang.org/x/oauth2 v0.0.0-20220808172628-8227340efae7
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/hypirion/go-filecache.v1 v1.0.0-20160810125507-e3e6ef6981f0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/semver/v3 v3.2.0 // indirect
	github.com/Microsoft/go-winio v0.5.2 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20230217124315-7d5c6f04bbb8 // indirect
	github.com/acomagu/bufpipe v1.0.4 // indirect
	github.com/cloudflare/circl v1.1.0 // indirect
	github.com/dlclark/regexp2 v1.7.0 // indirect
	github.com/eknkc/basex v1.0.0 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/fatih/color v1.10.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-git/gcfg v1.5.0 // indirect
	github.com/go-git/go-billy/v5 v5.4.1 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.10.0 // indirect
	github.com/go-sourcemap/sourcemap v2.1.3+incompatible // indirect
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/goccy/go-json v0.9.7 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/huandu/xstrings v1.3.3 // indirect
	github.com/imdario/mergo v0.3.13 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.11.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.2.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.10.0 // indirect
	github.com/jackc/pgx/v4 v4.15.0 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/kevinburke/ssh_config v1.2.0 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/logrusorgru/aurora v0.0.0-20181002194514-a7b3b318ed4e // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.1 // indirect
	github.com/pjbgf/sha1cd v0.3.0 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/skeema/knownhosts v1.1.0 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/tidwall/btree v1.1.0 // indirect
	github.com/tidwall/grect v0.1.4 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/rtred v0.1.2 // indirect
	github.com/tidwall/tinyqueue v0.1.1 // indirect
	github.com/xanzy/ssh-agent v0.3.3 // indirect
	github.com/xi2/xz v0.0.0-20171230120015-48954b6210f8 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
)
