module github.com/joeymckenzie/realworld-go-kit/conduit-core

go 1.18

require (
	entgo.io/ent v0.10.1
	github.com/go-kit/kit v0.12.0
	github.com/go-kit/log v0.2.0
	github.com/go-playground/validator/v10 v10.11.0
	github.com/gosimple/slug v1.12.0
	github.com/joeymckenzie/realworld-go-kit/conduit-ent-gen v0.0.0-20220506175452-016436a59380
	github.com/joeymckenzie/realworld-go-kit/conduit-shared v0.0.0-20220506175452-016436a59380
	github.com/mattn/go-sqlite3 v1.14.12
	github.com/stretchr/testify v1.7.1
)

replace (
	github.com/joeymckenzie/realworld-go-kit/conduit-domain => ../conduit-domain
	github.com/joeymckenzie/realworld-go-kit/conduit-ent-gen => ../conduit-ent-gen
	github.com/joeymckenzie/realworld-go-kit/conduit-shared => ../conduit-shared
)

require (
	ariga.io/atlas v0.3.7-0.20220303204946-787354f533c3 // indirect
	github.com/agext/levenshtein v1.2.1 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-chi/chi/v5 v5.0.7 // indirect
	github.com/go-chi/cors v1.2.1 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gosimple/unidecode v1.0.1 // indirect
	github.com/hashicorp/hcl/v2 v2.10.0 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lib/pq v1.10.4 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/mitchellh/go-wordwrap v0.0.0-20150314170334-ad45545899c7 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.12.1 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.32.1 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/zclconf/go-cty v1.8.0 // indirect
	golang.org/x/crypto v0.0.0-20220427172511-eb4f295cb31f // indirect
	golang.org/x/mod v0.5.1 // indirect
	golang.org/x/sys v0.0.0-20220114195835-da31bd327af9 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
