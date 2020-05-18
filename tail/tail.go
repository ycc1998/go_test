package tail

import (
	"time"
	"github.com/hpcloud/tail"
	"test/global"
	
)


var tailFile *tail.Tail


// Init :
func Init(filename string) (err error){

    tailFile, err = tail.TailFile(filename, tail.Config{
        ReOpen:    true,
        Follow:    true,
        Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
        MustExist: false,
        Poll:      true,
    })

    if err != nil {
        return
	}
	
	go run()
	return
}


func run(){
	for {
        select{
		case msg, ok := <- tailFile.Lines:
			if !ok {
				time.Sleep(50 * time.Millisecond)
				continue
			}
			
			global.AllUrl <- msg.Text
		default:
			time.Sleep(50 * time.Millisecond)
		}
    }
}

