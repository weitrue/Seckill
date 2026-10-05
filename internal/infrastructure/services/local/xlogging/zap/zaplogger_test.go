package xzap

var testLoggerConf = LogConf{
	ServiceName: "testService",
	Mode:        "file",
	Path:        "test",
	Level:       "info",
	Compress:    false,
	KeepDays:    1,
}
