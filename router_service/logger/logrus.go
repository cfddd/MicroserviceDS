package logger

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path"
)

// 颜色
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct{}

var Log *logrus.Logger

// Format 实现Formatter(entry *logrus.Entry) ([]byte, error)接口
func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同的level去展示颜色
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
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	prefix := viper.GetString("logger.prefix")

	//自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		//自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		//自定义输出格式
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", prefix, timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s\n", prefix, timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

func InitLogger() {
	mLog := logrus.New()                                    //新建一个实例
	mLog.SetOutput(os.Stdout)                               //设置输出类型
	mLog.SetReportCaller(viper.GetBool("logger.show_line")) //开启返回函数名和行号
	mLog.SetFormatter(&LogFormatter{})                      //设置自己定义的Formatter
	level, err := logrus.ParseLevel(viper.GetString("logger.level"))
	if err != nil {
		level = logrus.InfoLevel
	}
	mLog.SetLevel(level) //设置最低的Level
	InitDefaultLogger()
	Log = mLog
	Log.Info("日志模块初始化成功")
}

func InitDefaultLogger() {
	// 全局log
	logrus.SetOutput(os.Stdout)                               //设置输出类型
	logrus.SetReportCaller(viper.GetBool("logger.show_line")) //开启返回函数名和行号
	logrus.SetFormatter(&LogFormatter{})                      //设置自己定义的Formatter
	level, err := logrus.ParseLevel(viper.GetString("logger.level"))
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level) //设置最低的Level
}
