package gcurl

import (
	"bufio"
	"bytes"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func ParseCookie(r io.Reader) ([]*http.Cookie, error) {
	scanner := bufio.NewScanner(r)
	cookies := []*http.Cookie{}

	for scanner.Scan() {
		// 每行格式为：domain \t 是否HttpOnly \t cookie路径path \t 是否secure \t 有效期expires \t cookie名 \t cookie值
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}
		split := strings.Split(line, "\t")
		if len(split) < 7 {
			continue
		}

		expiresSplit := strings.Split(split[4], ".")
		expiresSec, err := strconv.Atoi(expiresSplit[0])
		if err != nil {
			return nil, err
		}
		expiresNsec := 0
		if len(expiresSplit) > 1 {
			expiresNsec, err = strconv.Atoi(expiresSplit[1])
			if err != nil {
				expiresNsec = 0
			}
		}
		cookie := &http.Cookie{
			Name:     split[5],
			Value:    split[6],
			Path:     split[2],
			Domain:   split[0],
			Expires:  time.Unix(int64(expiresSec), int64(expiresNsec)),
			Secure:   strings.ToLower(split[3]) == "true",
			HttpOnly: strings.ToLower(split[1]) == "true",
		}
		cookies = append(cookies, cookie)
	}
	return cookies, nil
}

func ParseCookie4String(s string) ([]*http.Cookie, error) {
	return ParseCookie(bytes.NewReader([]byte(s)))
}

func ParseCookie4File(cookieFile string) ([]*http.Cookie, error) {
	f, err := os.Open(cookieFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ParseCookie(f)
}
