package main

import "fmt"

type StringIntMap struct {
	Name  []string
	Value []int
}

func (m *StringIntMap) Add(key string, value int) {
	m.Name = append(m.Name, key)
	m.Value = append(m.Value, value)
}

func (m *StringIntMap) Get(key string) (int, bool) {
	for i, v := range m.Name {
		if v == key {
			return m.Value[i], true
		}
	}
	return 0, false
}

// Remove method returns true if there was value for provided key
func (m *StringIntMap) Remove(key string) bool {
	for i, v := range m.Name {
		if v == key {
			m.Name = append(m.Name[:i], m.Name[i+1:]...)
			m.Value = append(m.Value[:i], m.Value[i+1:]...)
			return true
		}
	}
	return false
}

func (m *StringIntMap) Exists(key string) bool {
	for _, v := range m.Name {
		if v == key {
			return true
		}
	}
	return false
}

func (m *StringIntMap) Copy() StringIntMap {
	c := StringIntMap{
		Name:  make([]string, len(m.Name)),
		Value: make([]int, len(m.Value)),
	}

	copy(c.Name, m.Name)
	copy(c.Value, m.Value)

	return c
}

func main() {
	sim := new(StringIntMap)

	sim.Add("Aleksandr", 1)
	fmt.Println(sim.Get("Aleksandr"))
	fmt.Println(sim.Exists("Aleksandr"))
	fmt.Println(sim.Copy())
	fmt.Println(sim.Remove("Aleksandr"))

}
