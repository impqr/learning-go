package main

import (
	"github.com/valyala/fasthttp"
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"getBaidu", args{"http://baidu.com"}, []byte("<html>\n<meta http-equiv=\"refresh\" content=\"0;url=http://www.baidu.com/\">\n</html>\n")},
		{"getLocal", args{"http://127.0.0.1:8080"}, []byte("you are request for /")},
	}
	LoadClient()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Get(tt.args.addr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	tests := []struct {
		name string
		want *fasthttp.Client
	}{
		{"newClient", LoadClient()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoadClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPost(t *testing.T) {
	type args struct {
		addr string
		body []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"postBaidu", args{"http://baidu.com", []byte("{}")}, []byte("<html>\n<meta http-equiv=\"refresh\" content=\"0;url=http://www.baidu.com/\">\n</html>\n")},
		{"postLocal", args{"http://127.0.0.1:8080", []byte("{}")}, []byte("you are request for /")},
	}
	LoadClient()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Post(tt.args.addr, tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Post() = %v, want %v", got, tt.want)
			}
		})
	}
}
