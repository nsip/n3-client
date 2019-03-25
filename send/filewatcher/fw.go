package filewatcher

import (
	"fmt"
	"io/ioutil"
	"time"

	s ".."
	u "github.com/cdutwhu/go-util"
	"github.com/fsnotify/fsnotify"
)

var (
	fPln = fmt.Println
)

// StartFileWatcherAsync :
func StartFileWatcherAsync() {
	defer func() { s.PH(recover(), s.CFG.Global.ErrLog) }()

	watcher := s.Must(fsnotify.NewWatcher()).(*fsnotify.Watcher)

	defer watcher.Close()
	s.PE(watcher.Add(s.CFG.Filewatcher.DirSif))
	s.PE(watcher.Add(s.CFG.Filewatcher.DirXapi))

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			s.LPln("event:", event) // CREATE WRITE REMOVE RENAME
			if event.Op&fsnotify.Create == fsnotify.Create {
				time.Sleep(2 * time.Second)
				s.LPln("created file:", event.Name)

			READ_AGAIN:
				bytes, e := ioutil.ReadFile(event.Name)
				if e != nil && s.SC(e.Error(), "The process cannot access the file because it is being used by another process") {
					fPln("read file failed, trying again ...")
					time.Sleep(1000 * time.Millisecond)
					goto READ_AGAIN
				}

				str := u.Str(string(bytes))
				if str.IsJSON() {
					s.Xapi(str.V())
				} else {
					s.Sif(str.V())
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			s.LPln("error:", err)
		}
	}
}
