package utility

import (
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"io/ioutil"
	"strings"
)

// Win1251ToUTF8 - Русские фразы из базы нужно конвертировать из win1251 в utf8 этим:
func Win1251ToUTF8(str string) (string, error) {
	tr := transform.NewReader(strings.NewReader(str), charmap.Windows1251.NewDecoder())
	buf, err := ioutil.ReadAll(tr)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}
