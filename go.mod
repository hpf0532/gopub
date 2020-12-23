module github.com/hpf0532/gopub

go 1.12

require (
	github.com/astaxie/beego v1.12.1
	github.com/bndr/gojenkins v1.0.1
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-ldap/ldap/v3 v3.2.4
	github.com/go-sql-driver/mysql v1.5.0
	github.com/hpf0532/sshexec v0.0.0-20200228123807-89cda2655e44
	github.com/julienschmidt/httprouter v1.3.0
	github.com/linclin/gopub v0.0.0-20200228133029-96dbe260850b
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/toolkits/net v0.0.0-20160910085801-3f39ab6fe3ce
	golang.org/x/crypto v0.0.0-20201217014255-9d1352758620
)

replace (
	github.com/hpf0532/sshexec v0.0.0-20200228123807-89cda2655e44 => /home/hpf/sshexec/
	github.com/linclin/gopub v0.0.0-20200228133029-96dbe260850b => ./
	golang.org/x/crypto v0.0.0-20200221231518-2aa609cf4a9d => github.com/golang/crypto v0.0.0-20200221231518-2aa609cf4a9d
)
