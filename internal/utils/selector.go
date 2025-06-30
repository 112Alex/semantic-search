package utils

import "golang.org/x/net/html"

// CSSSelector возвращает простой CSS-путь до узла.
// Формат: tag:nth-child(idx) > tag:nth-child(idx)... от <html> до нода.
func CSSSelector(n *html.Node) string {
	if n == nil {
		return ""
	}
	var parts []string
	for cur := n; cur != nil && cur.Type != html.DocumentNode; cur = cur.Parent {
		if cur.Type == html.ElementNode || cur.Type == html.TextNode {
			idx := 1
			for sib := cur.PrevSibling; sib != nil; sib = sib.PrevSibling {
				if sib.Type == cur.Type && sib.Data == cur.Data {
					idx++
				}
			}
			parts = append([]string{cur.Data + ":nth-child(" + itoa(idx) + ")"}, parts...)
		}
	}
	res := ""
	for i, p := range parts {
		if i > 0 {
			res += " > "
		}
		res += p
	}
	return res
}

func itoa(i int) string {
	// маленькая helper без strconv для скорости написания
	if i == 0 {
		return "0"
	}
	digits := ""
	for i > 0 {
		digits = string('0'+(i%10)) + digits
		i /= 10
	}
	return digits
}
