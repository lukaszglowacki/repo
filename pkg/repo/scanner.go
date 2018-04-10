package repo

import (
	"bufio"
	"io"

	"github.com/lukaszglowacki/repo/pkg/log"
)

func scanOutput(std io.ReadCloser) {
	scanner := bufio.NewScanner(std)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		log.Info(m)
	}
}
