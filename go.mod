module github.com/eden-framework/srv-identity-platform

go 1.14

replace k8s.io/client-go => k8s.io/client-go v0.18.8

require (
	github.com/eden-framework/apollo v0.0.1
	github.com/eden-framework/client v0.0.0-20201022095936-63753150b326
	github.com/eden-framework/context v0.0.2
	github.com/eden-framework/courier v1.0.4
	github.com/eden-framework/eden-framework v1.1.8-0.20201028102439-8a18bdc96161
	github.com/eden-framework/enumeration v1.0.0
	github.com/eden-framework/sqlx v0.0.1
	github.com/eden-framework/timex v0.0.3-0.20201030063221-a30952ce9f2a
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/onsi/ginkgo v1.13.0 // indirect
	github.com/profzone/envconfig v1.4.7-0.20201023033232-7c4ac8aaab6a
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/cobra v0.0.5
	golang.org/x/sys v0.0.0-20201029080932-201ba4db2418 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
)
