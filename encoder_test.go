package wsjtx

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func Test_encodeHeartbeat(t *testing.T) {
	type args struct {
		msg HeartbeatMessage
	}
	wantBin, _ := hex.DecodeString("adbccbda00000002000000000000000657534a542d580000000300000005322e322e3200000006306439623936")
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "encodeHeartbeat",
			args: args{msg: HeartbeatMessage{
				Id:        "WSJT-X",
				MaxSchema: 3,
				Version:   "2.2.2",
				Revision:  "0d9b96"}},
			want:    wantBin,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := encodeHeartbeat(tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeHeartbeat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encodeHeartbeat() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encodeClear(t *testing.T) {
	type args struct {
		msg ClearMessage
	}
	wantBin, _ := hex.DecodeString("adbccbda00000002000000030000000657534a542d5802")
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "encodeClear",
			args:    args{msg: ClearMessage{"WSJT-X", 2}},
			want:    wantBin,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := encodeClear(tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeClear() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encodeClear() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encodeReply(t *testing.T) {
	type args struct {
		msg ReplyMessage
	}
	wantBin, _ := hex.DecodeString("adbccbda00000002000000040000000657534a542d580259baf8fffffffb3fc99999a000000000000516000000017e0000000e4a4132454a50204e3442502037330000")
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "encodeReply",
			args: args{msg: ReplyMessage{
				Id:               "WSJT-X",
				Time:             39435000,
				Snr:              -5,
				DeltaTimeSec:     0.20000000298023224,
				DeltaFrequencyHz: 1302,
				Mode:             "~",
				Message:          "JA2EJP N4BP 73",
				LowConfidence:    false,
				Modifiers:        0,
			}},
			want:    wantBin,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := encodeReply(tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeReply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encodeReply() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encodeClose(t *testing.T) {
	type args struct {
		msg CloseMessage
	}
	wantBin, _ := hex.DecodeString("adbccbda00000002000000060000000657534a542d58")
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "encodeClose",
			args:    args{msg: CloseMessage{"WSJT-X"}},
			want:    wantBin,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := encodeClose(tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeClose() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encodeClose() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encodeReplay(t *testing.T) {
	type args struct {
		msg ReplayMessage
	}
	wantBin, _ := hex.DecodeString("adbccbda00000002000000070000000657534a542d58")
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "encodeReplay",
			args:    args{msg: ReplayMessage{"WSJT-X"}},
			want:    wantBin,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := encodeReplay(tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeReplay() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encodeReplay() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encodeHaltTx(t *testing.T) {
	type args struct {
		msg HaltTxMessage
	}
	wantBin, _ := hex.DecodeString("adbccbda00000002000000080000000657534a542d5800")
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "encodeHaltTx",
			args:    args{msg: HaltTxMessage{"WSJT-X", false}},
			want:    wantBin,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := encodeHaltTx(tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeHaltTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encodeHaltTx() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encodeFreeText(t *testing.T) {
	type args struct {
		msg FreeTextMessage
	}
	wantBin, _ := hex.DecodeString("adbccbda00000002000000090000000657534a542d5800000010f09f988a20646520f09f87baf09f87b801")
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "encodeFreeText",
			args:    args{msg: FreeTextMessage{"WSJT-X", "😊 de 🇺🇸", true}},
			want:    wantBin,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := encodeFreeText(tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeFreeText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encodeFreeText() got = %v, want %v", got, tt.want)
			}
		})
	}
}
