package Abnt

import (
	"github.com/Darklabel91/ABNT_AuthorNormalizer/Abnt"
	"github.com/Darklabel91/ABNT_AuthorNormalizer/AbntCSV"
	"reflect"
	"testing"
)

func TestTransformABNT(t *testing.T) {
	type args struct {
		authorName string
	}
	tests := []struct {
		name    string
		args    args
		want    Abnt.ABNTData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Abnt.TransformABNT(tt.args.authorName)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransformABNT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransformABNT() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransformABNTCSV(t *testing.T) {
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
			if err := AbntCSV.TransformABNTCSV(tt.args.rawFilePath, tt.args.separator, tt.args.nameResultFolder); (err != nil) != tt.wantErr {
				t.Errorf("TransformABNTCSV() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
