package in_memory

import (
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/engine"
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/grpc/proto/urls"
	"reflect"
	"testing"
)

func TestNewURLsService(t *testing.T) {
	tests := []struct {
		name string
		want *URLsService
	}{
		{
			name: "correct url service",
			want: &URLsService{
				workers: map[int32]*engine.Worker{},
				ids:     0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewURLsService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewURLsService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURLsService_Create(t *testing.T) {
	type fields struct {
		workers map[int32]*engine.Worker
		ids     int32
	}
	type args struct {
		url      string
		interval int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int32
		wantErr bool
	}{
		{
			name: "correct create",
			fields: fields{
				workers: map[int32]*engine.Worker{
					0: engine.NewWorker("https://httpbin.org/range/15", 60, NewSaver()),
				},
				ids: 1,
			},
			args: args{
				url:      "https://httpbin.org/range/15",
				interval: 60,
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &URLsService{
				workers: tt.fields.workers,
				ids:     tt.fields.ids,
			}
			got, err := i.Create(tt.args.url, tt.args.interval)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURLsService_Delete(t *testing.T) {
	type fields struct {
		workers map[int32]*engine.Worker
		ids     int32
	}
	type args struct {
		id int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "correct delete",
			fields: fields{
				workers: map[int32]*engine.Worker{
					0: engine.NewWorker("https://httpbin.org/range/15", 60, NewSaver()),
				},
				ids: 1,
			},
			args:    args{id: 0},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &URLsService{
				workers: tt.fields.workers,
				ids:     tt.fields.ids,
			}
			if err := i.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestURLsService_Get(t *testing.T) {
	type fields struct {
		workers map[int32]*engine.Worker
		ids     int32
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*urls.Url
		wantErr bool
	}{
		{
			name: "correct get",
			fields: fields{
				workers: map[int32]*engine.Worker{
					0: engine.NewWorker("https://httpbin.org/range/15", 60, NewSaver()),
				},
				ids: 1,
			},
			want: []*urls.Url{{
				Id:       0,
				Url:      "https://httpbin.org/range/15",
				Interval: 60,
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &URLsService{
				workers: tt.fields.workers,
				ids:     tt.fields.ids,
			}
			got, err := i.Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURLsService_History(t *testing.T) {
	type fields struct {
		workers map[int32]*engine.Worker
		ids     int32
	}
	type args struct {
		id int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*urls.Response
		wantErr bool
	}{
		{
			name: "empty history",
			fields: fields{
				workers: map[int32]*engine.Worker{
					0: engine.NewWorker("https://httpbin.org/range/15", 60, NewSaver()),
				},
				ids: 1,
			},
			args:    args{0},
			want:    []*urls.Response{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &URLsService{
				workers: tt.fields.workers,
				ids:     tt.fields.ids,
			}
			got, err := i.History(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("History() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("History() got = %v, want %v", got, tt.want)
			}
		})
	}
}
