package docuconf

import (
	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"log"
)

var k = koanf.New(".")

func LoadDotEnv[T interface{}](filePath string, envStruct T) (T, error) {
	f := file.Provider(filePath)
	if err := k.Load(f, dotenv.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	err := k.UnmarshalWithConf("", &envStruct, koanf.UnmarshalConf{Tag: "env"})
	if err != nil {
		return envStruct, err
	}
	return envStruct, nil
}
