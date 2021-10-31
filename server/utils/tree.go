package utils

type TreeNode map[string]interface{}

func GetTree(list []TreeNode, parentName string, idName string, parent string) []TreeNode {
	treeMap := getTreeMap(list, parentName)

	parents := treeMap[parent]

	for i := 0; i < len(parents); i++ {
		getChildrenList(parents[i], treeMap, idName)
	}

	return parents
}

func getChildrenList(item TreeNode, treeMap map[string][]TreeNode, idName string) TreeNode {

	id := item[idName]
	idStr := Strval(id)
	tt := treeMap[idStr]
	for i := 0; i < len(tt); i++ {
		getChildrenList(tt[i], treeMap, idName)
	}

	item["children"] = tt

	return item
}

func getTreeMap(list []TreeNode, parentName string) map[string][]TreeNode {
	treeMap := make(map[string][]TreeNode)
	for _, v := range list {
		_parentName := v[parentName]
		parentNameStr := Strval(_parentName)
		treeMap[parentNameStr] = append(treeMap[parentNameStr], v)
	}

	return treeMap
}
