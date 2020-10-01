package in_memory

import (
	"github.com/Mrucznik/U3p5bW9uLUdhamRh/grpc/proto/urls"
	"reflect"
	"testing"
)

var testResponses = []*urls.Response{
	{
		Response:  "abcdefghijklmno,",
		Duration:  0.571,
		CreatedAt: 1559034638,
	},
	{
		Response:  "null", //TODO: nil
		Duration:  5,
		CreatedAt: 1559034938,
	},
}

func TestNewSaver(t *testing.T) {
	tests := []struct {
		name string
		want *Saver
	}{
		{
			"correct saver",
			&Saver{[]*urls.Response{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSaver(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSaver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaver_GetResults(t *testing.T) {
	type fields struct {
		results []*urls.Response
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*urls.Response
		wantErr bool
	}{
		{
			"empty results",
			fields{results: []*urls.Response{}},
			[]*urls.Response{},
			false,
		},
		{
			"some results",
			fields{results: testResponses},
			testResponses,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Saver{
				results: tt.fields.results,
			}
			got, err := s.GetResults()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetResults() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetResults() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaver_Save(t *testing.T) {
	type fields struct {
		results []*urls.Response
	}
	type args struct {
		data *urls.Response
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "correct save",
			fields:  fields{[]*urls.Response{}},
			args:    args{data: testResponses[0]},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Saver{
				results: tt.fields.results,
			}
			if err := s.Save(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
