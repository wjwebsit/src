package main


func isMatch2(s string, p string) bool {
	//声明缓存
	dp := make(map[int]map[int]bool)

	return myMatch(p,0,s,0,&dp)
}

func myMatch(p string ,i int,s string, j int,dp *map[int]map[int]bool) bool { //p匹配s
	//p到尽头了
	if i >= len(p) {//p到尽头了
		return j == len(s)
	}

	//判断缓存
	if _,ok := (*dp)[i][j];ok {
		return (*dp)[i][j]
	}

	//第一个是否匹配
	FG := false
	if j < len(s) && (p[i] == s[j] || p[i] == '.') {
		FG = true
	}

	//p[i + 1] 是否为‘*’
	if i + 1 < len(p) && p[i + 1] == '*' {
		//匹配0次或多次
		FG = myMatch(p,i + 2,s,j,dp) || (FG && myMatch(p,i,s,j + 1,dp))

	} else {//DFS
		FG = FG && myMatch(p, i + 1,s, j + 1,dp)
	}

	//写入缓存
	if (*dp)[i] == nil {//二维map初始化
		(*dp)[i] = make(map[int]bool)
	}

	//写入
	(*dp)[i][j] = FG

	//返回
	return FG
}
