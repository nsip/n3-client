package preprocess

import (
	"os"
	"os/exec"

	g "github.com/nsip/n3-client/global"
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
	pe(os.Chdir(jqDir))
}

// FmtJSONStr : <json string> must not have single quote <'>
func FmtJSONStr(json string, jqDirs ...string) string {
	defer func() { pe(os.Chdir(g.OriExePath)) }()
	prepareJQ(jqDirs...)
	if !IsJSON(json) {
		return ""
	}
	sJSON := S(json).Replace("'", "\\'") //        *** deal with <single quote> in "echo" ***
	cmdstr := "echo $" + sJSON.MkQuotes(QSingle).V() + ` | ./jq .`
	cmd := exec.Command("bash", "-c", cmdstr)
	output := must(cmd.Output()).([]byte)
	return string(output)
}

// FmtJSONFile : <file> is the <relative path> to <jq>
func FmtJSONFile(file string, jqDirs ...string) string {
	defer func() { pe(os.Chdir(g.OriExePath)) }()
	prepareJQ(jqDirs...)
	cmdstr := "cat " + file + ` | ./jq .`
	cmd := exec.Command("bash", "-c", cmdstr)
	output := must(cmd.Output()).([]byte)
	return string(output)
}
