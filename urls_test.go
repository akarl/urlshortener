package main

import "testing"

type URL struct {
	in, out string
}

var urls []URL = []URL{
	URL{"http://google.com", "http://google.com"},
	URL{"google.com", "http://google.com"},
}

func TestAddGetURL(t *testing.T) {
	storage := &URLStorage{
		RedisPool: NewRedisPool(),
	}

	for _, url := range urls {
		c, err := storage.AddURL(url.in)
		if err != nil {
			t.Error(err)
		}

		saved_url, err := storage.GetURL(c)
		if err != nil {
			t.Error(err)
		}

		if saved_url != url.out {
			t.Errorf("%v != %s", saved_url, url.out)
		}
	}

}
