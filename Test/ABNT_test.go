package Author

import (
	"github.com/Darklabel91/ABNT_AuthorNormalizer"
	"github.com/Darklabel91/ABNT_AuthorNormalizer/Structs"
	"reflect"
	"testing"
)

func TestAbntFormat(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    Structs.DataABNT
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Author.AbntFormat(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("AbntFormat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AbntFormat() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbntFormatCSV(t *testing.T) {
	type args struct {
		rawFilePath      string
		separator        rune
		nameResultFolder string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Author.AbntFormatCSV(tt.args.rawFilePath, tt.args.separator, tt.args.nameResultFolder); (err != nil) != tt.wantErr {
				t.Errorf("AbntFormatCSV() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
