package logs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

const (
	FieldKeyMsg            = "message"
	FieldKeyTimestamp      = "timestamp"
	FieldKeyLevel          = "severity"
	FieldKeyApplicationLog = "log"
	FieldKeyFile           = "file"
	FieldKeyFunc           = "function"
)

type LogFormatter struct {
	WithColors bool
}

// Format renders a single log entry
func (f *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}

	applicationLog := make(logrus.Fields, len(entry.Data)+3)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			// Otherwise errors are ignored by `encoding/json`
			// https://github.com/sirupsen/logrus/issues/137
			applicationLog[k] = v.Error()
		default:
			applicationLog[k] = v
		}
	}
	applicationLog[FieldKeyMsg] = entry.Message
	applicationLog[FieldKeyTimestamp] = entry.Time.Format(time.RFC3339)
	applicationLog[FieldKeyLevel] = strings.ToUpper(entry.Level.String())

	if entry.HasCaller() {
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)
		if funcVal != "" {
			applicationLog[FieldKeyFunc] = funcVal
		}
		if fileVal != "" {
			applicationLog[FieldKeyFile] = fileVal
		}
	}

	data := make(logrus.Fields, 1)
	data[FieldKeyApplicationLog] = applicationLog

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	encoder := json.NewEncoder(b)
	encoder.SetEscapeHTML(true)
	if err := encoder.Encode(data); err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON, %v", err)
	}

	if f.WithColors == true {
		buf := &bytes.Buffer{}

		fmt.Fprintf(buf, "\x1b[%dm%s\x1b[0m", levelColor, b.Bytes())
		b = buf
	}

	return b.Bytes(), nil
}
