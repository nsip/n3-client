package filewatcher

import (
	"fmt"
	"io/ioutil"
	"time"

	s ".."
	w "github.com/cdutwhu/go-wrappers"
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
	s.PE(watcher.Add(s.CFG.Filewatcher.Dir))

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
				if e != nil && w.Str(e.Error()).Contains("The process cannot access the file because it is being used by another process") {
					fPln("read file failed, trying again ...")
					time.Sleep(1000 * time.Millisecond)
					goto READ_AGAIN
				}

				s.ToNode(string(bytes), "id", "xapi")
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			s.LPln("error:", err)
		}
	}
}
