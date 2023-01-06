Skip to content
Search or jump to…
Pull requests
Issues
Codespaces
Marketplace
Explore
 
@BlackPanthar 
silentnoname
/
supervisor
Private
Code
Issues
Pull requests
Actions
Projects
Security
Insights
supervisor
/
src
/
log
/
log.go
in
main
 

Tabs

8

No wrap
1
package log
2
​
3
import (
4
        "go.uber.org/zap"
5
        "go.uber.org/zap/zapcore"
6
        "gopkg.in/natefinch/lumberjack.v2"
7
        "os"
8
)
9
​
10
var Log *zap.Logger
11
​
12
// InitLog  initializes the logger
13
func InitLog() {
14
        fileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
15
                Filename:   "./nodemonitor.log",
16
                MaxSize:    100,
17
                MaxBackups: 3,
18
                MaxAge:     30,
19
                Compress:   false,
20
        })
21
        encoderConfig := zap.NewProductionEncoderConfig()
22
        encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
23
        encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
24
        encoder := zapcore.NewJSONEncoder(encoderConfig)
25
        fileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(fileWriteSyncer, zapcore.AddSync(os.Stdout)), zapcore.DebugLevel)
26
        Log = zap.New(fileCore, zap.AddCaller())
27
}
28
​
@BlackPanthar
Commit changes
Commit summary
Create log.go
Optional extended description
Add an optional extended description…

conwuliri@gmail.com
Choose which email address to associate with this commit

 Commit directly to the main branch.
 Create a new branch for this commit and start a pull request. Learn more about pull requests.
 
Footer
© 2023 GitHub, Inc.
Footer navigation
Terms
Privacy
Security
Status
Docs
Contact GitHub
Pricing
API
Training
Blog
About
