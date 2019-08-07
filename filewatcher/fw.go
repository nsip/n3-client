package filewatcher

import (
	"io/ioutil"
	"time"

	g "github.com/nsip/n3-client/global"
	pub "github.com/nsip/n3-client/publish"
	w "github.com/cdutwhu/go-wrappers"
	"github.com/fsnotify/fsnotify"
)

// StartFileWatcherAsync :
func StartFileWatcherAsync() {
	defer func() { ph(recover(), g.Cfg.ErrLog) }()

	watcher := must(fsnotify.NewWatcher()).(*fsnotify.Watcher)

	defer watcher.Close()
	pe(watcher.Add(g.Cfg.FileWatcher.Dir))

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

				if IDs, _, _, _, _, e := pub.Pub2Node(g.CurCtx, string(bytes), "xapi"); e != nil {
					fPln(e)
					return
				} else {
					for _, id := range IDs {
						fPln(id, "is sent")
					}
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			lPln("error:", err)
		}
	}
}
