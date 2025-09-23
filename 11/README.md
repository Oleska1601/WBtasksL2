Оценка сложности:

```
func sortString(s string) string {
	r := []rune(s)
	slices.Sort(r) //m log m
	return string(r)

}
```
функция **sortString** работает в среднем за O(mlogm) где m - длина слайса (в данному случае - длина слайса рун r, то есть длина нашей искомой строки)


```
func countAnagrams(sl []string) map[string][]string {
	if len(sl) == 0 {
		return nil
	}

	m := make(map[string][]string)
	for i := 0; i < len(sl); i++ {
		s := sortString(sl[i])
		m[s] = append(m[s], sl[i])
	}
	result := make(map[string][]string)
	for _, v := range m {
		if len(v) > 1 {
			key := v[0]
			result[key] = v
			slices.Sort(result[key])
		}
	}
	return result
}
```
ф-ия **countAnagrams**:
первый цикл for для заполнения мапы m:
- мы проходимся по всем элементам входного слайса -> следовательно сложность O(n), где n - длина входного слайса
- в этом цикле for вызывается ф-ия sortString, сложность которой O(mlogm) 
-> таким образом, сложность первого цикла O(n) * O(m logm) -> O(n*mlogm)
второй цикл for для заполнения мапы result:
- мы проходимся по всем элементам мапы m, максимальная длина которой равна длине входного слайса (в случае, если ни одной анаграммы нам не встретится)
-> таким образом, сложность второго цикла O(n)

Итоговая сложность ф-ии O(n *mlogm) + O(n) -> O(n *mlogm)

Таким образом.ю итоговая сложность алгоритма = O(n *mlogm), что удовлетворяет условию задачи

