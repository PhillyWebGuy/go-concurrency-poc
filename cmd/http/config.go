package main

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/joho/godotenv"

	"github.com/spf13/viper"
)

type configuration struct {
	Env      string
	Database struct {
		Name            string
		User            string
		Password        string
		Host            string
		Port            string
		ConnectionRetry int
	}
	HTTP struct {
		Domain string
		Port   string
	}
}

/*
In summary, this code reads a configuration file named "config.yml"
from the current directory, and unmarshals it into a configuration struct.
It also reads from the environment variables, with periods replaced by underscores.
*/

func mustNewConfiguration() configuration {

	//This calls a function that loads environment variables from a file named ".env" in the ".local" directory.
	loadEnv()

	//This tells Viper to automatically read from the environment variables.
	//Any environment variable that Viper is aware of will be read in this way.
	viper.AutomaticEnv()
	//This sets a replacer that changes all periods (.) in the environment variable keys to underscores (_).
	//This is useful because environment variables typically use underscores instead of periods.
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	//This sets the name of the configuration file that Viper will look for.
	//In this case, it's looking for a file named "config".
	viper.SetConfigName("config")
	//This sets the type of the configuration file that Viper will look for. In this case, it's looking for a YAML file.
	viper.SetConfigType("yml")
	//This adds a path for Viper to look in for the configuration file. In this case, it's the current directory.
	viper.AddConfigPath("./")

	//This tells Viper to read in the configuration file.
	//If there's an error (like the file doesn't exist), it will panic and print the error.
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("unable to decode config file, %v", err))
	}

	config := configuration{}
	//This creates a new instance of a configuration struct (not shown in the code you provided),
	//and then tries to unmarshal the configuration file into that struct.
	//If there's an error (like the file doesn't match the struct), it will panic and print the error.
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Sprintf("unable to decode config file into configuration, %v", err))
	}

	return config
}

func loadEnv() {
	slog.Info("DOTENV: Ready to load .env")
	if err := godotenv.Load("./.local/.env"); err != nil {
		slog.Info("DOTENV: No .env file found or error reading file.")
		return
	}
	slog.Info("DOTENV: .env found.")
}
