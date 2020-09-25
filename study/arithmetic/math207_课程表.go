package main
/**
	算法描述
	你这个学期必须选修 numCourse 门课程，记为 0 到 numCourse-1 。

在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：[0,1]

给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？

 

示例 1:

输入: 2, [[1,0]]
输出: true
解释: 总共有 2 门课程。学习课程 1 之前，你需要完成课程 0。所以这是可能的。
示例 2:

输入: 2, [[1,0],[0,1]]
输出: false
解释: 总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0；并且学习课程 0 之前，你还应先完成课程 1。这是不可能的。
 

提示：

输入的先决条件是由 边缘列表 表示的图形，而不是 邻接矩阵 。详情请参见图的表示法。
你可以假定输入的先决条件中没有重复的边。
1 <= numCourses <= 10^5
 */
 /**
 	利用拓扑排序
  */
func canFinish(numCourses int, prerequisites [][]int) bool {
	//判断
	if numCourses == 1 {
		return  true
	}

	//声明入度数组
	degree := make([]int,numCourses)

	//构造map
	edges := make(map[int][]int)

	//入度
	for i := 0;i < len(prerequisites);i ++ {
		//记录入度
		degree[prerequisites[i][0]] ++

		//记录边的指向
		edges[prerequisites[i][1]] = append(edges[prerequisites[i][1]],prerequisites[i][0])
	}

	//声明队列
	var queue []int

	//获取入度为0的入队
	for i := 0;i < numCourses; i ++ {
		if degree[i] == 0 {
			queue = append(queue,i)
		}
	}

	//统计
	sum := 0
	for len(queue) > 0 {
		//出队
		cur := queue[0]
		queue = queue[1:]

		//判断当前的指向
		if len(edges[cur]) > 0 {
			for _,v := range edges[cur] {
				//入度减一
				degree[v] --
				if degree[v] == 0 {
					queue = append(queue,v)
				}

			}

		}
		sum ++

	}

	//返回
	return sum == numCourses

}

