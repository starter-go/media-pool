package objects

import "testing"

func TestMeta(t *testing.T) {

	m := &Meta{}
	var sum Sum

	sum[0] = 10
	sum[1] = 1
	sum[2] = 2
	sum[3] = 3

	m.Name = "example.txt"
	m.ID = "fffffffffffff"
	m.Sum = sum
	m.Size = 1234
	m.Path = "a/b/c/d/example.txt"
	m.Type = "text/plain"

	str := m.String()
	t.Logf("obj.meta = \n %s", str)
}
