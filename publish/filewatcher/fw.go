package filewatcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	pub ".."
	g "../../global"
	w "github.com/cdutwhu/go-wrappers"
	"github.com/fsnotify/fsnotify"
)

var (
	fPln = fmt.Println
	lPln = log.Println
)

// StartFileWatcherAsync :
func StartFileWatcherAsync() {
	defer func() { pub.PH(recover(), pub.CFG.Global.ErrLog) }()

	watcher := pub.Must(fsnotify.NewWatcher()).(*fsnotify.Watcher)

	defer watcher.Close()
	pub.PE(watcher.Add(pub.CFG.Filewatcher.Dir))

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			lPln("event:", event) // CREATE WRITE REMOVE RENAME
			if event.Op&fsnotify.Create == fsnotify.Create {
				time.Sleep(2 * time.Second)
				lPln("created file:", event.Name)

			READ_AGAIN:
				bytes, e := ioutil.ReadFile(event.Name)
				if e != nil && w.Str(e.Error()).Contains("The process cannot access the file because it is being used by another process") {
					fPln("read file failed, trying again ...")
					time.Sleep(1000 * time.Millisecond)
					goto READ_AGAIN
				}
				IDs, _, _, _, _ := pub.Pub2Node(string(bytes), "id", "xapi")
				g.RmIDsInLRU(IDs...)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			lPln("error:", err)
		}
	}
}
