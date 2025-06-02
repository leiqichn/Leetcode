/*
 * Copyright (c) 2023 Lei Qi. All rights reserved.
 * Author: Lei Qi
 * Description:
 * Date: 2023/10/10 上午12:33
 */

package main

import (
	"container/list"
	"fmt"
)

var test int

type Browser struct {
	history  *list.List
	current  *list.Element
	maxLen   int
	homepage string
}

func NewBrowser(maxLen int, homepage string) *Browser {
	history := list.New()
	current := history.PushBack(homepage)
	return &Browser{
		history:  history,
		current:  current,
		maxLen:   maxLen,
		homepage: homepage,
	}
}

func (b *Browser) GetCurrentPage() string {
	return b.current.Value.(string)
}

func (b *Browser) GoBack() string {
	if b.current.Prev() != nil {
		b.current = b.current.Prev()
	}
	test += 1
	return b.GetCurrentPage()
}

func (b *Browser) GoForward() string {
	if b.current.Next() != nil {
		b.current = b.current.Next()
	}
	return b.GetCurrentPage()
}

func (b *Browser) NavigateToNewPage(newPageURL string) string {
	if b.GetCurrentPage() == newPageURL {
		return newPageURL
	}

	// 清除当前页面之后的历史记录
	for e := b.current.Next(); e != nil; e = e.Next() {
		b.history.Remove(e)
	}
	// 将新页面添加到历史记录中
	b.current = b.history.InsertAfter(newPageURL, b.current)
	// 限制浏览器历史记录的最大长度
	for b.history.Len() > b.maxLen {
		front := b.history.Front()
		if front != nil {
			b.history.Remove(front)
		}
	}
	return b.GetCurrentPage()
}

func main() {
	browser := NewBrowser(5, "初始页面")

	fmt.Println("当前页面:", browser.GetCurrentPage())

	// 浏览新页面
	fmt.Println("浏览新页面:", browser.NavigateToNewPage("新页面1"))
	fmt.Println("当前页面:", browser.GetCurrentPage())

	// 浏览更多新页面
	fmt.Println("浏览新页面:", browser.NavigateToNewPage("新页面2"))
	fmt.Println("浏览新页面:", browser.NavigateToNewPage("新页面3"))
	fmt.Println("浏览新页面:", browser.NavigateToNewPage("新页面4"))
	fmt.Println("浏览新页面:", browser.NavigateToNewPage("新页面5"))
	fmt.Println("当前页面:", browser.GetCurrentPage())

	// 后退和前进
	fmt.Println("后退:", browser.GoBack())
	fmt.Println("后退:", browser.GoBack())
	fmt.Println("前进:", browser.GoForward())
}
