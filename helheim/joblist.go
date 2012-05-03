package helheim


// Linear list of Jobs.
type JobList []*Job

// Add Job to list.
func(l*JobList)Append(j*Job){
	*l = append(*l, j)
}

// Remove Job from list.
func(l*JobList)Remove(job*Job){
	index:=-1
	for i, j := range *l{
		if j == job{
			index =i
			break
		}
	}
	*l=append((*l)[:index], (*l)[index+1:]...)
}
