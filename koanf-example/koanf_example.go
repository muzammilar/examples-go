package main

import (
	"fmt"
	"strings"

	"github.com/knadh/koanf"
)

// https://docs.github.com/en/copilot/using-github-copilot/getting-started-with-github-copilot

// Struct to hold the configuration
type Configuration struct {
	Ghardelli    string `koanf:"ghardelli" yaml:"ghardelli" json:"ghardelli"`
	Kinder       string
	Hersheys     string `koanf:"hersheys" yaml:"hersheys" json:"hersheys"`
	KitKat       kitKat `koanf:"kit_kat" yaml:"kit_kat" json:"kit_kat"`
	Lindt        Lindt
	MarsInc      MarsInc      `koanf:"mars_inc" yaml:"mars_inc" json:"mars_inc"`
	FerreroGroup FerreroGroup `koanf:"ferrero_group" yaml:"ferrero_group" json:"ferrero_group"`
}

type kitKat struct {
	Flavor string `koanf:"flavor" yaml:"flavor" json:"flavor"`
}

type Lindt struct {
	Lindor string `koanf:"lindor" yaml:"lindor" json:"lindor"`
}

type MarsInc struct {
	SnickerChocolate string `koanf:"snicker_chocolate" yaml:"snicker_chocolate" json:"snicker_chocolate"`
	Mars             string `koanf:"mars" yaml:"mars" json:"mars"`
}

type FerreroGroup struct {
	FerreroRocher FerreroRocher `koanf:"ferrero_rocher" yaml:"ferrero_rocher" json:"ferrero_rocher"`
}

type FerreroRocher struct {
	Raffaello string `koanf:"raffaello" yaml:"raffaello" json:"raffaello"`
}

func main() {
	var sb strings.Builder
	// load the configuration
	msg := NewMessage()

	// get all the tags and all the variables
	tagsSlice := GetFlattenedTags("msgtag", "msg", msg)

	// print the information about the embedded struct
	sb.WriteString(fmt.Sprintf("Details of the `Message` Struct:\n"))
	sb.WriteString(fmt.Sprintf("Variable Name\t\t\t Type\t Tag\n"))
	for _, i := range tagsSlice {
		sb.WriteString(fmt.Sprintf("%s\n", i))
	}
	fmt.Println(sb.String())

}

func getVarFromEnvName(s string) string {
	// read environmental variable and remove the prefix (and convert them to lower)
	s = strings.TrimPrefix(s, environmentVarPrefix)
	s = strings.ToLower(s)
	// convert all `__` to a `.` so that nested structs could be loaded via env variables
	// This allows something like MYAPP_ROGER_FED__RAFAEL to be converted to roger_fed.rafael
	koanfOption := strings.ReplaceAll(s, "__", ".")
	return koanfOption
}

func loadConfig(confPath string) (*Configuration, error) {
	var config Configuration
	var k = koanf.New(".")
	if err := k.Load(file.Provider(confPath), json.Parser()); err != nil {
		return nil, err
	}

	// Load values from environment variables MYAPP_*
	// For example, MYAPP_POSTGRES_USER will replace postgres_user loaded from config file
	_ = k.Load(env.Provider(environmentVarPrefix, ".", getVarFromEnvName), nil)

	// Note the Koanf allows you to have nested configs, but we are not using that here.
	// If you do use nested, use: k.UnmarshalWithConf("", &config, koanf.UnmarshalConf{Tag: "koanf", FlatPaths: true})
	if err := k.Unmarshal("", &config); err != nil {
		return nil, err
	}

	if err := validateConfig(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func validateConfig(config *Configuration) error {

}
