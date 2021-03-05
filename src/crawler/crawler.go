package crawler

import (
	"fmt"
	"github.com/taark/crawler/src/models"
	"io/ioutil"
	"net/http"
	urlHelper "net/url"
	"regexp"
	"sync"
	"time"
)

var re = regexp.MustCompile(`<title>(.*)</title>`)

//создание http клиента для отправки запросов
var httpClient = http.Client{Timeout: 3 * time.Second}

func Scan(urls []string) []*models.Target {
	var wg sync.WaitGroup //объект для синхронизации конкуретных методов
	var targets []*models.Target
	var m sync.Mutex //объект для блокировок

	/*
		в цикле перебераются все урлы
		все они обрабатываются параллельно
		результат записывается в массив targets



	*/
	for _, url := range urls {
		wg.Add(1)
		urlL := url
		// go перед функцией означает что ее нужно запустить в отдельном "потоке", т.е. "параллельно"
		go func() {
			defer wg.Done()
			title, err := getTitle(urlL)

			t := &models.Target{
				Url:   urlL,
				Title: title,
			}

			if err != nil {
				t.Err = err.Error()
			}

			//блокировка нужна чтобы только одна функция в один момент могла записывать в массив targets
			// т.е. когда один "поток" записывает результат в массив, други потоки приостанавливаются
			//это не лучшие решение, но я его использовал, чтобы продемонстировать работу с mutex
			m.Lock()
			targets = append(targets, t)
			m.Unlock()
		}()
	}

	wg.Wait()

	return targets

}

//здесь ничего особо интересного.

func validateUrl(url string) error {
	_, err := urlHelper.ParseRequestURI(url)
	return err
}

func getTitle(url string) (string, error) {
	if err := validateUrl(url); err != nil {
		return "", err
	}

	content, err := getContent(url)
	if err != nil {
		return "", err
	}

	match := re.FindStringSubmatch(content)
	if len(match) < 2 {
		return "", fmt.Errorf("title not found in %s ", url)
	}

	return match[1], nil

}

func getContent(url string) (string, error) {
	resp, err := httpClient.Get(url)
	if err != nil {
		return "", err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
