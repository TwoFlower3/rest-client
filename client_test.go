package client

import "testing"

func TestGET(t *testing.T) {
	client := NewClient("http://test.com")

	if _, err := client.GET("test", nil, map[string]string{}, map[string]string{}); err != nil {
		t.Fatal(err)
	}
}

func TestGETWithMaps(t *testing.T) {
	client := NewClient("http://test.com")

	if _, err := client.GET("test", nil, map[string]string{"header": "test"}, map[string]string{"query": "test"}); err != nil {
		t.Fatal(err)
	}
}

func TestGETWithoutMaps(t *testing.T) {
	client := NewClient("http://test.com")

	if _, err := client.GET("test", nil, nil, nil); err != nil {
		t.Fatal(err)
	}
}

func TestPOST(t *testing.T) {
	client := NewClient("http://test.com")

	if _, err := client.POST("test", nil, map[string]string{}, map[string]string{}); err != nil {
		t.Fatal(err)
	}
}

func TestPUT(t *testing.T) {
	client := NewClient("http://test.com")

	if _, err := client.PUT("test", nil, map[string]string{}, map[string]string{}); err != nil {
		t.Fatal(err)
	}
}

func TestDELETE(t *testing.T) {
	client := NewClient("http://test.com")

	if _, err := client.DELETE("test", nil, map[string]string{}, map[string]string{}); err != nil {
		t.Fatal(err)
	}
}
