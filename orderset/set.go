package orderset

//用golang 实现set集合，利用map特性去重(或者循环遍历，不过效率低)

func GetSet(array []interface{}) ([]interface{}) {
	var mp = make(map[interface{}]bool)
	for _, v := range array {
		mp[v] = true
	}

	retArr := make([]interface{}, len(mp))
	var i int
	for k, _ := range mp {
		retArr[i] = k
		i++
	}
	return retArr
}
