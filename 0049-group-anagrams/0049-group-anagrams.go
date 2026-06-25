func groupAnagrams(strs []string) [][]string {
    resMap := map[string][]string{}
    for _, s := range strs {
        sBytes := []byte(s)
        slices.Sort(sBytes)
        key := string(sBytes)
        resMap[key] = append(resMap[key], s)
    }

    groups := make([][]string, 0, len(resMap))
    for _, v := range resMap {
        groups = append(groups, v)
    }
    return groups
}