package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strconv"
	"strings"
)

var (
	defCfg      map[string]string
	initialized = false
)

func initialize() {
	viper.SetEnvPrefix("aaa")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	defCfg = make(map[string]string)

	defCfg["server.host"] = "localhost"
	defCfg["server.port"] = "3000"
	defCfg["server.log.level"] = "warn" // valid values are trace, debug, info, warn, error, fatal
	defCfg["server.timeout.write"] = "15 seconds"
	defCfg["server.timeout.read"] = "15 seconds"
	defCfg["server.timeout.idle"] = "60 seconds"
	defCfg["server.timeout.graceshut"] = "15 seconds"

	defCfg["setup.admin.enable"] = "false"
	defCfg["setup.admin.email"] = "admin@hansip"
	defCfg["setup.admin.passphrase"] = "this must be change in the production"

	defCfg["token.issuer"] = "aaa.domain.com"
	defCfg["token.access.duration"] = "5 minutes"
	defCfg["token.refresh.duration"] = "1 year"

	defCfg["token.crypt.key"] = "th15mustb3CH@ngedINprodUCT10N"
	defCfg["token.crypt.method"] = "HS512"

	defCfg["db.type"] = "INMEMORY" // INMEMORY, MYSQL
	defCfg["db.mysql.host"] = "localhost"
	defCfg["db.mysql.port"] = "3306"
	defCfg["db.mysql.user"] = "devuser"
	defCfg["db.mysql.password"] = "devpassword"
	defCfg["db.mysql.database"] = "devdb"
	defCfg["db.mysql.maxidle"] = "3"
	defCfg["db.mysql.maxopen"] = "10"

	defCfg["mailer.type"] = "DUMMY" // DUMMY, SENDMAIL
	defCfg["mailer.from"] = "hansip@aaa.com"
	defCfg["mailer.sendmail.host"] = "localhost"
	defCfg["mailer.sendmail.port"] = "25"
	defCfg["mailer.sendmail.user"] = "sendmail"
	defCfg["mailer.sendmail.password"] = "password"
	defCfg["mailer.templates.emailveri.subject"] = "Please verify your new Hansip account's email"
	defCfg["mailer.templates.emailveri.body"] = "<html><body>Dear New Hansip User<br><br>Your new account is ready!<br>please click this <a href=\"http://hansip.io/activate?code={{.ActivationCode}}\">link to activate</a> your account.<br><br>Cordially,<br>HANSIP team</body></html>"
	defCfg["mailer.templates.passrecover.subject"] = "Passphrase recovery instruction"
	defCfg["mailer.templates.passrecover.body"] = "<html><body>Dear Hansip User<br><br>To recover your passphrase<br>please click this <a href=\"http://hansip.io/activate?code={{.RecoveryCode}}\">link to change your passphrase</a>.<br><br>Cordially,<br>HANSIP team</body></html>"

	for k, _ := range defCfg {
		err := viper.BindEnv(k)
		if err != nil {
			log.Errorf("Failed to bind env \"%s\" into configuration. Got %s", k, err)
		}
	}

	initialized = true
}

func SetConfig(key, value string) {
	viper.Set(key, value)
}

func Get(key string) string {
	if !initialized {
		initialize()
	}
	ret := viper.GetString(key)
	if len(ret) == 0 {
		if ret, ok := defCfg[key]; ok {
			return ret
		}
		log.Debugf("%s config key not found", key)
	}
	return ret
}

func GetBoolean(key string) bool {
	if len(Get(key)) == 0 {
		return false
	}
	b, err := strconv.ParseBool(Get(key))
	if err != nil {
		panic(err)
	}
	return b
}

func GetInt(key string) int {
	if len(Get(key)) == 0 {
		return 0
	}
	i, err := strconv.ParseInt(Get(key), 10, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}

func GetFloat(key string) float64 {
	if len(Get(key)) == 0 {
		return 0
	}
	f, err := strconv.ParseFloat(Get(key), 64)
	if err != nil {
		panic(err)
	}
	return f
}

func Set(key, value string) {
	defCfg[key] = value
}