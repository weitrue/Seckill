/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2023/3/30 4:07 PM
 * Description:
 **/

package merkle

import (
	"reflect"
	"testing"
)

func TestGetRoot(t *testing.T) {
	type args struct {
		addresses []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "OK",
			args: args{
				addresses: []string{
					"0xe01511d7333A18e969758BBdC9C7f50CcF30160A",
					"0x62d17DE1fbDF36597F12F19717C39985A921426e",
					"0x6F702345360D6D8533d2362eC834bf5f1aB63910",
				},
			},
			want:    "0x9593ef2d207a3738fb385a662acab9077e8ea343fa0867400bbfa5539350b46c",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRoot(tt.args.addresses)
			if err != tt.wantErr {
				t.Errorf("GetRoot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetRoot() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetProof(t *testing.T) {
	type args struct {
		address   string
		addresses []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr error
	}{
		{
			name: "OK",
			args: args{
				addresses: []string{
					"0xe01511d7333A18e969758BBdC9C7f50CcF30160A",
					"0x62d17DE1fbDF36597F12F19717C39985A921426e",
					"0x6F702345360D6D8533d2362eC834bf5f1aB63910",
				},
				address: "0xe01511d7333A18e969758BBdC9C7f50CcF30160A",
			},
			want: []string{
				"0xea3a488603068aaf2632f108365edcd62563e193024c6af02b498c8b9b9a2120",
				"0x28d889ab829c62f3fddd900df2440f7766be4278537d601a0d6a5949963f5374",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetProof(tt.args.address, tt.args.addresses)
			if err != tt.wantErr {
				t.Errorf("GetProof() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProof() got = %v, want %v", got, tt.want)
			}
		})
	}
}
