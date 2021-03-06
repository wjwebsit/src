package main

/**
	邻接多重表---用于解决邻接表存储无向图的时候的删除问题
	**画图策略：**先列出六条边，每条边都有两个指针域，这两个指针域可以视为等同，因为是无向的，然后根据原图连接指针线。
   举个例子，假如排列好了六条边，从a开始，有01与03，先确定a指向的第一条边，假设是01吧，然后01结点指向a的下一条边，指向的开始位置为0右边的指针域，没有下一条边那么相应指针域为^；
   然后开始搞b……最后都连接好了再美观下位置

 */

/**
	邻接矩阵结构体
 */
type lVertex struct {
	data      string //顶点标识
	firstEdge *lEdge
}

type lEdge struct {
	ivex   int    //边的起始点
	ilink  *lEdge //已边为起始的顶点的下一个指向域
	jvex   int    //变的结束点
	jlink  *lEdge //已边为终止的的顶点的下一个指向域
	mark   int    //是否被标记
	weight int    //权
}

type lChart struct {
	vertex             []lVertex
	numVertex, numEdge int
}

//构造过程略

/** 边集数组存储方式(edgeSet array)
	是利用一维数组存储图中所有边的一种图的表示方法。
    该数组中所含元素的个数要大于等于图中边的条数，每个元素用来存储一条边的起点、终点（对于无向图，可选定边的任一端点为起点或终点）和权（若有的话），
   各边在数组中的次序可任意安排，也可根据具体要求而定。
 */

/**
		结构体和构造过程见 data_24图的深度优先遍历--边集数组
 */
