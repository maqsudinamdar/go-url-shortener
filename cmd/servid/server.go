package servid

import (
	"flag"
	"fmt"
	"net/http"
	"url-shortener/internal/platform/db"
	"url-shortener/internal/url"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	defaultConfigPath	= 	"./configs"
	configFilePathUsage    = "config file directory. Config file must be named 'conf_{env}.yml'."
	configFilePathFlagName = "configFilePath"
	envUsage               = "environment for app, prod, dev, test"
	envDefault             = "dev"
	envFlagname            = "env"

)

var configFilePath string
var env string

func config() {
	logger()
	flag.StringVar(&configFilePath, configFilePathFlagName, defaultConfigPath, configFilePathUsage)
	flag.StringVar(&env, envFlagname, envDefault, envUsage)
	flag.Parse()
	configuration(configFilePath, env)
}


func logger() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors: true,
		DisableColors: false,
		FullTimestamp: true,
	})
}


type App struct {
	*http.Server
	r 			*chi.Mux
	db 			*sqlx.DB
	urlRouter	*url.Router
}

func NewApp() *App {
	config()
	router := chi.NewRouter()
	database := setUpDb(viper.GetString("database.URL"))
	urlsRouter := url.NewRouter(router, database)
	server := &App {
		r: 			router,
		db:			database,
		urlRouter: 	urlsRouter,
	}
	server.routes()
	return server
}

func (a *App) routes() {
	a.urlRouter.Routes()
	showRoutes(a.r)
}

func (a *App) Start() {
	log.Fatal(http.ListenAndServe(viper.GetString("server.port"), a.r))
}

func showRoutes(r *chi.Mux) {
	log.Info("registered routes: ")
	walkFunc := func(method string, route string, handler http.Handler, m ...func(http.Handler) http.Handler) error {
		log.Infof("%s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(r, walkFunc); err != nil {
		log.Infof("Logging err: %s\n", err.Error())
	}
}


func configuration(path string, env string) {
	if flag.Lookup("test.v") != nil {
		env = "test"
		path = "./../../configs"
	}
	log.Println("Environment is: " + env + " configFilePath is: " + path)
	viper.SetConfigName("conf_" + env)
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("fatal: %+v", err))
	}
}

func setUpDb(dbUrl string) *sqlx.DB {
	mysql, err := db.New(dbUrl)
	if err != nil {
		log.Fatal(fmt.Errorf("fatal: %+v", err))
	}
	return mysql
}