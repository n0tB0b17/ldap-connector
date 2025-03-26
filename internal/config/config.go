package config

type Cfg struct {
	LDAPHost     string
	LDAPPort     int
	LDAPUser     string
	LDAPPassword string
	RabbitMQURL  string
	Workers      int
}

func GetCfg() Cfg {
	return Cfg{
		LDAPHost:    "ldap://192.168.1.6",
		LDAPPort:    389,
		RabbitMQURL: "amqp://rabbit:rabbit123@localhost:5672",
		Workers:     10,
	}
}
