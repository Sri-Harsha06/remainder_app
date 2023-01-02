package customlogger

import (
    "log"
    "os"
    "sync"
)


// func init(){
//     file,err :=os.OpenFile("file.log",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0666)
//     if(err!=nil){
// 		log.Fatal(err)
// 	}
// 	InfoLogger=log.New(file,"INFO:",log.LstdFlags|log.Lshortfile)
// 	WarningLogger=log.New(file,"WARNING:",log.LstdFlags|log.Lshortfile)
// 	ErrorLogger=log.New(file,"ERROR:",log.LstdFlags|log.Lshortfile)
// }
type loggers struct {
    filename string
    InfoLogger    *log.Logger
    WarningLogger *log.Logger
    ErrorLogger   *log.Logger
}

var logger *loggers
var once sync.Once

func GetInstance() *loggers {
    once.Do(func() {
        logger = createLogger("file.log")
    })
    return logger
}

func createLogger(fname string) *loggers {
    file,err :=os.OpenFile(fname,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0666)
    if(err!=nil){
		log.Fatal(err)
	}
    return &loggers{
        filename: fname,
        InfoLogger: log.New(file,"INFO:",log.LstdFlags|log.Lshortfile),
        WarningLogger: log.New(file,"WARNING:",log.LstdFlags|log.Lshortfile),
        ErrorLogger: log.New(file,"ERROR:",log.LstdFlags|log.Lshortfile)}
}