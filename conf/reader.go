package conf

import (
	"os"
	"bufio"
	"strings"
	"io"
)

type Properties interface {
	Get(string) string
	Read(confPath string)
}
type Configuration struct {
	Properties map[string]string
}

func (c Configuration) Get(key string) (value string) {
	return c.Properties[key]
}
func (c *Configuration) Read(confPath string) error {
	f, err := os.Open(confPath)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(f)
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		//获得第一个=的下标
		index := strings.IndexRune(s, []rune("=")[0])
		if index != -1 {
			str := []rune(s)
			c.Properties[string(str[:index])] = strings.Replace(string(str[index+1:]),"\r\n","",-1)
		}
	}
	return nil
}
