package backups3

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

type Logstruct struct {
	path string `toml:"path"`
}

const ERR_RUNNING = 5
const ERR_LOG_NOT_FOUND = 6

var LogstructConfig = `
  ##Sample Config
  #path = "/var/log"
`

func (s *Logstruct) SampleConfig() string {
	return LogstructConfig
}

func (s *Logstruct) Description() string {
	return "Collects metrics from script backup to s3"
}

func (s *Logstruct) Gather(acc telegraf.Accumulator) error {

	t := time.Now()
	buf := make([]byte, 2)
	path := "/var/log"
	prefix := []string{"www", "databases"}
	var www, db int
	for i := range prefix {
		p := fmt.Sprintf("%s/backup_%s_%s.log", path, prefix[i], t.Format("2006-01-02"))
		f, err := os.Open(p)
		if err != nil {
			if prefix[i] == "www" {
				www = ERR_LOG_NOT_FOUND
				break
			} else {
				db = ERR_LOG_NOT_FOUND
				break
			}
		}
		defer f.Close()
		s, err := f.Stat()
		if err != nil {
			fmt.Println(err)
		}

		n, err := f.ReadAt(buf, s.Size()-2)
		if n < len(buf) {
			fmt.Println("InvalidFileErr")
		}
		if prefix[i] == "www" {
			www, err = strconv.Atoi(string(buf[:1]))
			if err != nil {
				www = ERR_RUNNING
				// fmt.Println(err)

			}
		} else {
			db, err = strconv.Atoi(string(buf[:1]))
			if err != nil {
				// fmt.Println(err)
				db = ERR_RUNNING
			}
		}

	}

	//Fields association
	fields := make(map[string]interface{})
	fields["www"] = www
	fields["db"] = db

	tags := make(map[string]string)

	acc.AddFields("asterisk", fields, tags)

	return nil

}

func init() {
	inputs.Add("backup_s3", func() telegraf.Input { return &Logstruct{} })
}
