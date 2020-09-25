package main
/**
	算法描述
	设计一个方法，找出任意指定单词在一本书中的出现频率。

你的实现应该支持如下操作：

WordsFrequency(book)构造函数，参数为字符串数组构成的一本书
get(word)查询指定单词在书中出现的频率
示例：

WordsFrequency wordsFrequency = new WordsFrequency({"i", "have", "an", "apple", "he", "have", "a", "pen"});
wordsFrequency.get("you"); //返回0，"you"没有出现过
wordsFrequency.get("have"); //返回2，"have"出现2次
wordsFrequency.get("an"); //返回1
wordsFrequency.get("apple"); //返回1
wordsFrequency.get("pen"); //返回1
提示：

book[i]中只包含小写字母
1 <= book.length <= 100000
1 <= book[i].length <= 10
get函数的调用次数不会超过100000
 */
type WordsFrequency struct {
	root *trieNode
}

/**
	字典树结点
 */
type trieNode struct {
	//单词数目
	count int
	//当前字符
	ch byte
	//孩子
	children [26]*trieNode
	//是否为单词
	mark bool
}


func Constructor(book []string) WordsFrequency {
	//先初始化WordsFrequency
	w := new(WordsFrequency)
	w.root = new(trieNode)
	w.root.mark = false


	//初始化
	for i := 0 ; i < len(book); i ++ {
		insetTrie(w.root,book[i])
	}

	//返回
	return *w
}
func insetTrie(root *trieNode,word string) {
	//添加单词搜索树
	this := root

	for i := 0 ; i < len(word);i ++ {
		//判断child
		if this.children[word[i] - 'a'] != nil {
			//初始化
			this = this.children[word[i] - 'a']

		} else {//初始化
			node := new (trieNode)
			node.ch = word[i]
			node.count = 0
			node.mark = false

			this.children[word[i] - 'a'] = node
			this = node
		}
	}

	//标记为单词
	this.mark = true
	this.count ++ //单词数目+1


}

func (this *WordsFrequency) Get(word string) int {
	//查找单词
	p := this.root
	for i := 0 ; i < len(word);i ++ {
		//判断
		if p.children[word[i] - 'a'] == nil {//不存在
			return 0
		}
		p = p.children[word[i] - 'a']
	}
	if p.mark == false {
		return 0
	}

	//返回
	return p.count


}