package main
/**
	算法描述:
	设计一个支持以下两种操作的数据结构：

void addWord(word)
bool search(word)
search(word) 可以搜索文字或正则表达式字符串，字符串只包含字母 . 或 a-z 。 . 可以表示任何一个字母。

示例:

addWord("bad")
addWord("dad")
addWord("mad")
search("pad") -> false
search("bad") -> true
search(".ad") -> true
search("b..") -> true
说明:

你可以假设所有单词都是由小写字母 a-z 组成的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-and-search-word-data-structure-design
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
type WordDictionary struct {
	root *trie
}

/**
	字典树数据结构
 */
type trie struct {
	chr byte //当前结点
	min ,max int //当前节点在内的后面的子长度最大和最小--均为0时表示为空,min==max 只有一个结点
	fg bool //是否为单词
	children [26]*trie
}


/** Initialize your data structure here. */
func Constructor22() WordDictionary {
	//初始化
	w := new(WordDictionary)
	w.root = new(trie)
	w.root.min = 0
	w.root.max = 0
	w.root.fg = false

	//返回
	return *w
}
/**
	初始化
 */
func initTrie() *trie {
	t := new(trie)
	t.min = 0
	t.max = 0
	t.fg = false

	//返回
	return t
}


/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string)  {
	//判断
	if len(word) == 0 {
		return
	}

	//获取单词长度
	length := len(word)

	//获取树节点
	p := this.root

	//添加单词
	for i := 0 ; i < length; i ++ {
		//获取长度
		n := length - i

		//写入
		if p.min == 0 {
			p.min = n
			p.max = n

		} else {
			if p.min > n {//更新当前最小值
				p.min = n
			} else if p.max < n {//更新当前最大值
				p.max = n
			}
		}

		//下发到孩子
		index := word[i] - 'a'

		//判断
		if p.children[index] == nil {
			//初始化
			 p.children[index] = initTrie()
			 p.children[index].chr = word[i]

		}

		//继续
		p = p.children[index]

	}

	//标记为单词
	p.fg = true

}


/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	//查询
	if len(word) == 0 {
		return true
	}
	if this.root .min == 0 {//字典为空
		return false
	}

	//返回
	return mySearch(this.root,0,word)

}
func mySearch(tree *trie,s int,word string) bool{
	//判断
	if len(word) == s {
		return tree != nil && tree.fg == true
	}

	//判断tree是否为nil--word 不为nil
	if tree == nil {
		return false
	}

	//匹配
	if word[s] == '.' {
		//遍历子
		var res = false
		for i := 0 ; i < 26; i ++ {
			if tree.children[i] == nil {
				continue
			}

			res = mySearch(tree.children[i],s + 1,word)

			//判断
			if res {
				return true
			}

		}

		return false

	} else if tree.min > len(word) - s || tree.max < len(word) - s{
		return false

	} else {//继续

		return mySearch(tree.children[word[s] - 'a'],s + 1,word)
	}
}

func main() {

}