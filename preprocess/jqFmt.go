package preprocess

import (
	"os"
	"os/exec"
)

func prepareJQ(jqDirs ...string) {
	jqPath, jqDir := "./jq", "./"
	for _, jqDir = range jqDirs {
		if _, err := os.Stat(jqDir + "jq"); err == nil {
			jqPath = jqDir + "jq"
			break
		}
	}
	if _, err := os.Stat(jqPath); err != nil {
		panic("jq is not found")
	}
	PE(os.Chdir(jqDir))
}

// FmtJSONStr :
func FmtJSONStr(json string, jqDirs ...string) string {
	prepareJQ(jqDirs...)
	if !IsJSON(json) {
		return ""
	}
	cmdstr := "echo " + Str(json).MkQuotes(QSingle).V() + ` | jq .`
	cmd := exec.Command("bash", "-c", cmdstr)
	output := Must(cmd.Output()).([]byte)
	return string(output)
}

// FmtJSONFile : <file> is the <relative path> to <jq>
func FmtJSONFile(file string, jqDirs ...string) string {
	prepareJQ(jqDirs...)
	cmdstr := "cat " + file + ` | jq .`
	cmd := exec.Command("bash", "-c", cmdstr)
	output := Must(cmd.Output()).([]byte)
	return string(output)
}
