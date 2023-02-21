package remetrics

import (
	"context"
	"remetrics/pkg/conf"
	"sync"
	"testing"
)

var rmts = Remetrics{
	InputCfg:     nil,
	InputData:    nil,
	IsNewVersion: false,
}

func TestRemetrics_CheckPromUp(t *testing.T) {
	type fields struct {
		InputCfg     *conf.Config
		InputData    *conf.Input
		IsNewVersion bool
	}
	type args struct {
		ctx context.Context
		r   Remetrics
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			fields: fields{
				InputCfg:     nil,
				InputData:    nil,
				IsNewVersion: false,
			},
			args: args{
				ctx: nil,
				r:   rmts,
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := &Remetrics{
				InputCfg:     tt.fields.InputCfg,
				InputData:    tt.fields.InputData,
				IsNewVersion: tt.fields.IsNewVersion,
			}
			got, err := rt.CheckPromUp(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckPromUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckPromUp() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemetrics_InitLibs(t *testing.T) {
	type fields struct {
		InputCfg     *conf.Config
		InputData    *conf.Input
		IsNewVersion bool
	}
	type args struct {
		ctx context.Context
		r   Remetrics
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test 01",
			fields: fields{
				InputCfg:     nil,
				InputData:    nil,
				IsNewVersion: false,
			},
			args: args{
				ctx: nil,
				r:   Remetrics{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := &Remetrics{
				InputCfg:     tt.fields.InputCfg,
				InputData:    tt.fields.InputData,
				IsNewVersion: tt.fields.IsNewVersion,
			}
			if err := rt.InitLibs(tt.args.ctx, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("InitLibs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRemetrics_PusherV3(t *testing.T) {
	type fields struct {
		InputCfg     *conf.Config
		InputData    *conf.Input
		IsNewVersion bool
	}
	type args struct {
		ctx context.Context
		r   Remetrics
		idx int
		wg  *sync.WaitGroup
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := &Remetrics{
				InputCfg:     tt.fields.InputCfg,
				InputData:    tt.fields.InputData,
				IsNewVersion: tt.fields.IsNewVersion,
			}
			got, err := rt.PusherV3(tt.args.ctx, tt.args.r, tt.args.idx, tt.args.wg)
			if (err != nil) != tt.wantErr {
				t.Errorf("PusherV3() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PusherV3() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemetrics_Relable(t *testing.T) {
	type fields struct {
		InputCfg     *conf.Config
		InputData    *conf.Input
		IsNewVersion bool
	}
	type args struct {
		ctx context.Context
		r   Remetrics
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := &Remetrics{
				InputCfg:     tt.fields.InputCfg,
				InputData:    tt.fields.InputData,
				IsNewVersion: tt.fields.IsNewVersion,
			}
			got, err := rt.Relable(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Relable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Relable() got = %v, want %v", got, tt.want)
			}
		})
	}
}
