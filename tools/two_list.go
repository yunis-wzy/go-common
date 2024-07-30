package tools

func TwoList(src, dst []string) ([]string, []string, []string) {
	// 源map
	SrcMap := make(map[string]byte)
	// 并集map
	AllMap := make(map[string]byte)
	// 交集
	var set []string
	// 针对源列表建立map
	for _, v := range src {
		SrcMap[v] = 0
		AllMap[v] = 0
	}
	// 目标列表元素如果没有存储进去则为重复元素，放入交集部分
	for _, v := range dst {
		l := len(AllMap)
		AllMap[v] = 1
		if l == len(AllMap) {
			set = append(set, v)
		}
	}
	// 遍历交集，在交集中找到并删除，删除后就是对称差集
	for _, v := range set {
		delete(AllMap, v)
	}
	// 右差集
	var add []string
	// 左差集
	var del []string
	// 遍历堆成差集，如果存在源map中则为左差集（减少的部分），如果不存在则为右差集（增加的部分）
	for v := range AllMap {
		_, exist := SrcMap[v]
		if exist {
			del = append(del, v)
		} else {
			add = append(add, v)
		}
	}
	// 返回三部分
	return add, del, set
}