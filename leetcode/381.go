package main

/*
leetcode 381
设计一个支持在平均 时间复杂度 O(1) 下， 执行以下操作的数据结构。

注意: 允许出现重复元素。

insert(val)：向集合中插入元素 val。
remove(val)：当 val 存在时，从集合中移除一个 val。
getRandom：从现有集合中随机获取一个元素。每个元素被返回的概率应该与其在集合中的数量呈线性相关。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/insert-delete-getrandom-o1-duplicates-allowed
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	mp    map[int]int
	slice []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	a := map[int]int{}
	b := []int{}
	return RandomizedSet{mp: a, slice: b}

}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.mp[val]
	if !ok {
		if this.slice == nil {
			this.mp[val] = 0
		} else {
			this.mp[val] = len(this.slice)
		}
		this.slice = append(this.slice, val)
		return true
	} else {
		return false
	}

}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.mp[val]
	if ok {
		this.slice[len(this.slice)-1], this.slice[index] = this.slice[index], this.slice[len(this.slice)-1]
		//重点  容易遗漏  切莫忘记
		this.mp[this.slice[index]] = index
		this.slice = this.slice[:len(this.slice)-1]
		delete(this.mp, val)
		return true
	} else {
		return false
	}

}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	if this.slice == nil {
		return 0
	}
	lenth := len(this.slice)
	if lenth > 1 {
		r := rand.Intn(lenth)
		return this.slice[r]
	}
	return this.slice[0]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 ["RandomizedSet","insert","remove","insert","getRandom","remove","insert","getRandom"]
[[],[1],[2],[2],[],[1],[2],[]]
*/

func main() {
	a := Constructor()
	fmt.Printf("%+v\n", a)
	a.Insert(0)
	fmt.Printf("%+v\n", a)
	a.Insert(1)
	fmt.Printf("%+v\n", a)
	a.Remove(0)
	fmt.Printf("%+v\n", a)
	a.Insert(2)
	fmt.Printf("%+v\n", a)
	a.Remove(1)
	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", a.GetRandom())
}

//["RandomizedSet","insert","insert","remove","insert","remove","getRandom"]
//[[],[0],[1],[0],[2],[1],[]]
