package main

import (
	"reflect"
	"testing"
)

// 👇 添加单元测试的测试数据
var testPage = Page {
	Title: "👋 Hi, 这是一个大标题!",
	Body: []byte("这是一段简单的内容."),
}


func Test_loadPage(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name    string
		args    args
		want    *Page
		wantErr bool
	}{
		// 👇 添加单元测试的测试数据
		{"loadPage", args{title: testPage.Title}, &testPage, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := loadPage(tt.args.title)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadPage() = %v, want %v", got, tt.want)
			}

			// 👇 打印测试数据
			t.Logf("获得 body: %#v", string(got.Body))
		})
	}
}

func TestPage_savePage(t *testing.T) {
	type fields struct {
		Title string
		Body  []byte
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// 👇 添加单元测试的测试数据
		{"savePage", fields(testPage), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Page{
				Title: tt.fields.Title,
				Body:  tt.fields.Body,
			}
			if err := p.savePage(); (err != nil) != tt.wantErr {
				t.Errorf("Page.savePage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
