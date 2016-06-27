package main

import "testing"

func TestNewRedisPool(t *testing.T) {
	rp := NewRedisPool()

	conn, err := rp.Dial()
	if err != nil {
		t.Error(err)
	}

	if _, err := conn.Do("PING"); err != nil {
		t.Error(err)
	}
}
