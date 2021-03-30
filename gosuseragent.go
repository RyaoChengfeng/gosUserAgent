package gosUserAgent

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

// 来源https://gist.githubusercontent.com/pzb/b4b6f57144aea7827ae4/raw/cf847b76a142955b1410c8bcef3aabe221a63db1/user-agents.txt
func init() {
	rand.Seed(time.Now().UnixNano())
}


func getUserAgentList() (io.ReadCloser, error) {
	file,err := os.Open("user-agent.txt")
	if err != nil {
		return nil, err
	}
	return file, nil
}

func GetRandomUserAgent() (string, error) {
	var lines []string
	body, err := getUserAgentList()
	if err != nil {
		return "", fmt.Errorf("get user agent lists failed, got %v", err)
	}

	scanner := bufio.NewScanner(body)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	random := randomRange(0, len(lines)-1)
	if len(lines)==0 {
		return "", nil
	}
	userAgent := lines[random]

	return userAgent, nil
}

func randomRange(min, max int) int {
	return rand.Intn(max-min+1) + min
}
