package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type MyStruct struct {
	Result  string
	Message string
	Num     int
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func main() {
	mystruct1 := new(MyStruct)
	getJson("https://script.google.com/macros/s/AKfycbysqFfkeAwejuz6hd0RcnMjQfcjACF5YvG-BTKjfgQAsdkeqyI/exec", mystruct1)

	mystruct2 := new(MyStruct)
	getJson("https://script.google.com/a/geniee.co.jp/macros/s/AKfycbzyjnCduTYBWEiqflpMzpKwxtLQdj62lof72hA/exec", mystruct2)

	mystruct3 := new(MyStruct)
	getJson("https://script.google.com/a/geniee.co.jp/macros/s/AKfycbyYfrQAda2z5PqL_Bmt03MurWocyQGiAke7kUDdFQ/exec", mystruct3)

	fmt.Println(mystruct1.Message, mystruct2.Message, mystruct3.Message)
	fmt.Println(mystruct1.Num + mystruct2.Num + mystruct3.Num)

	var wg sync.WaitGroup
	ch := make(chan int, 3)

	wg.Add(3)
	go func() {
		defer wg.Done() // deferキーワードは関数を最後に実行する
		getJson("https://script.google.com/macros/s/AKfycbysqFfkeAwejuz6hd0RcnMjQfcjACF5YvG-BTKjfgQAsdkeqyI/exec", mystruct1)
		ch <- mystruct1.Num
	}()

	go func() {
		defer wg.Done() // deferキーワードは関数を最後に実行する
		getJson("https://script.google.com/a/geniee.co.jp/macros/s/AKfycbzyjnCduTYBWEiqflpMzpKwxtLQdj62lof72hA/exec", mystruct2)
		ch <- mystruct2.Num
	}()

	go func() {
		defer wg.Done() // deferキーワードは関数を最後に実行する
		getJson("https://script.google.com/a/geniee.co.jp/macros/s/AKfycbyYfrQAda2z5PqL_Bmt03MurWocyQGiAke7kUDdFQ/exec", mystruct3)
		ch <- mystruct3.Num
	}()

	wg.Wait() // カウントが0になるまで待つ

	close(ch) // 送信は済んだのでchannelはcloseする。closeしないと下記ループから抜け出せない。
	var result int
	for v := range ch { //channelの中身が一個ずつ読み出される
		result += v
	}
	fmt.Println(result)
}
