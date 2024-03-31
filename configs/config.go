package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var cfg *conf

type conf struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JwtExperesIn  int    `mapstructure:"JWT_EXPIRESIN"`
	TokenAuth     *jwtauth.JWTAuth
}

// Função para carregar as configurações a partir de um arquivo informando o path.
// Normalmente colocamos o arquivo ".env" dentro da pasta cmd/aplicacao.
func LoadConfig(path string) (*conf, error) {
	viper.SetConfigName("app_config") //podemos trabalhar com várias configurações
	viper.SetConfigType("env")        // Tipo de configuração que vamos trabalhar. Pode ser variável (env), yaml, toml, json, etc.
	viper.AddConfigPath(path)         // Caminho do arquivo de onde serão carregadas as configurações.
	viper.SetConfigFile(".env")
	viper.AutomaticEnv() //Utiliza as variéis de ambiente já configuradas, caso existam. Isso é importante quando temos um configmap no kubernetes, por exemplo.
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg, err

}

/*
Para carregar as configurações, vamos utilizar o pacote viper. Ele é muito famoso na comunidade.
github.com/spf13/viper

Pacote com o roteador. Trabalharemos com ele no decorrer do capítulo.
"github.com/go-chi/jwtauth"





*/
