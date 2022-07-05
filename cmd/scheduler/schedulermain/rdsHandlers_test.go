package schedulermain

import "testing"

func Test_getDBInstanceStatus(t *testing.T) {
	type args struct {
		output string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test_available",
			args: args{
				output: "DBInstanceStatus: \"available\"",
			},
			want:    "available",
			wantErr: false,
		},
		{
			name: "test_running",
			args: args{
				output: "DBInstanceStatus: \"running\"",
			},
			want:    "running",
			wantErr: false,
		},
		{
			name: "test_invalid",
			args: args{
				output: "DBInstanceStatus: ",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getDBInstanceStatus(tt.args.output)
			if (err != nil) != tt.wantErr {
				t.Errorf("getDBInstanceStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getDBInstanceStatus() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDBClusterStatus(t *testing.T) {
	type args struct {
		output string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test_available",
			args: args{
				output: "Status: \"available\"",
			},
			want:    "available",
			wantErr: false,
		},
		{
			name: "test_running",
			args: args{
				output: "Status: \"running\"",
			},
			want:    "running",
			wantErr: false,
		},
		{
			name: "test_invalid",
			args: args{
				output: "Status: ",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getDBClusterStatus(tt.args.output)
			if (err != nil) != tt.wantErr {
				t.Errorf("getDBClusterStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getDBClusterStatus() got = %v, want %v", got, tt.want)
			}
		})
	}
}
