package main

/**
	字典树二进制结点
 */
type Trade struct {
	fg bool
	left,right *Trade
	char byte
	count int //重复次数
	p *Trade
}

/**
	字典树声明
 */
type TradeTree struct {
	Root *Trade
}

/**
	初始化字典树
 */
func initTradeTree() *TradeTree {
	t := new(TradeTree)
	t.Root = new(Trade)
	return t
}

/**
	添加数字至数字字典树---不遍历符号位
 */
func (t *TradeTree) insert(num int,size int) {
	//获取当前的值
	tree := t.Root

	//遍历
	for i := 0; i < size; i ++ {
		//获取当前的值
		if num >> (size - i) & 1 == 1 {//当前的高位为1
			if tree.right == nil {
				tree.right = new(Trade)
				tree.right.fg = false
				tree.right.char = '1'
				tree.right.p = tree
			}
			tree = tree.right

		} else {//当前的高位为0
			if tree.left == nil {
				tree.left = new(Trade)
				tree.left.fg = false
				tree.left.char = '0'
				tree.left.p = tree
			}
			tree = tree.left
		}
	}

	//结束--标记为结点
	tree.fg = true
	tree.count += 1
}

/**
	批量添加数字
 */
func (t *TradeTree) insertNums(nums []int,size int) {
	//循环添加
	for _,num := range nums {
		t.insert(num,size)
	}
}

/**
	删除数字---向上删除
 */
func (t *TradeTree) delete(num int,size int) {
	tree := t.Root
	for i := 0; i < size; i ++ {
		//获取当前的值
		if num >> (num - i) & 1 == 1 {
			tree = tree.right
		} else {
			tree = tree.left
		}
	}

	//删除
	tree.count -= 1
	if tree.count == 0 && tree.left == nil && tree.right == nil{//删除只能删除叶子结点
		//---级联删除
		p := tree.p

		for p != nil {
			if p.left == tree {
				p.left = nil
			} else {
				p.right = nil
			}
			if p.left == nil && p.right == nil && p.count == 0 {
				tree = p
				p = p.p

			} else {
				break
			}
		}

		if p == nil {
			//初始化根
			t.Root = new(Trade)
		}
	}
}





