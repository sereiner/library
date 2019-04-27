package jsons

import "strings"

// String values encode as JSON strings coerced to valid UTF-8,
// replacing invalid bytes with the Unicode replacement rune.
// The angle brackets "<" and ">" are escaped to "\u003c" and "\u003e"
// to keep some browsers from misinterpreting JSON output as HTML.
// Ampersand "&" is also escaped to "\u0026" for the same reason.

// HTMLEscape 函数将json编码的src中的<、>、&、U+2028 和U+2029字符替换为\u003c、\u003e、\u0026、\u2028、\u2029 转义字符串，
// 以便json编码可以安全的嵌入HTML的<script>标签里。
// 因为历史原因，网络浏览器不支持在<script>标签中使用标准HTML转义，
// 因此必须使用另一种json编码方案。

// Escape 把编码 \\u0026，\\u003c，\\u003e 替换为 &,<,>
func Escape(input string) string {
	r := strings.Replace(input, "\\u0026", "&", -1)
	r = strings.Replace(r, "\\u003c", "<", -1)
	r = strings.Replace(r, "\\u003e", ">", -1)
	r = strings.Replace(r, "\n", "", -1)
	return r
}
