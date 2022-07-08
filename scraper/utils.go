package scraper

import "github.com/gocolly/colly"

// @Author: Feng
// @Date: 2022/5/16 21:29

//Parse 解析context为KV列表切片
func Parse(ctx *colly.Context, ignore map[string]struct{}) (re []KV) {
	ctx.ForEach(func(k string, v interface{}) interface{} {
		if _, ok := ignore[k]; !ok {
			re = append(re, KV{
				Key: k,
				Val: v,
			})
		}
		return nil
	})
	return re
}

//ParseToMap 解析context为KV列表切片
func ParseToMap(ctx *colly.Context, ignore map[string]struct{}) map[string]interface{} {
	re := map[string]interface{}{}
	ctx.ForEach(func(k string, v interface{}) interface{} {
		if _, ok := ignore[k]; !ok {
			re[k] = v
		}
		return nil
	})
	return re
}
