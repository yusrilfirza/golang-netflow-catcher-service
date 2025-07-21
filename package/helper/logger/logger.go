package logger

import (
	"encoding/json"
	"fmt"
	"netflow-catcher-service/package/helper/contract"
	"os"

	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

type LoggerField map[string]interface{}

type loggerImplementation struct{}

func (i *loggerImplementation) Log(level logrus.Level, message string, field interface{}, err error) {
	log := &logrus.Logger{
		Out:   os.Stdout,
		Level: logrus.DebugLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%] %time% - %msg%\n",
		},
		ReportCaller: true,
	}

	var logMessage string = message

	if field != nil {
		jsonField, marshalErr := json.Marshal(field)
		if marshalErr != nil {
			logMessage += fmt.Sprintf(" [field error: %v]", marshalErr)
		} else {
			logMessage += fmt.Sprintf(" %s", string(jsonField))
		}
	}

	if err != nil {
		logMessage += fmt.Sprintf(" error: %v", err)
	}

	log.Log(level, logMessage)
}

func NewLogger() contract.Logger {
	return &loggerImplementation{}
}
