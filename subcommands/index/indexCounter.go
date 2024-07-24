
package index

type indexCounter struct {
	levelCounts [6]int
}
func (self *indexCounter) increment(level int) {
	self.levelCounts[level]++
	for i := level + 1; i < len(self.levelCounts); i++ {
		self.levelCounts[i] = 0
	}
}
