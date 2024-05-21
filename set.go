package go-set


type Set[T comparable] struct{
  data map[T]struct{}
}

func(self Set[T]) Contains(val T) bool{
  _, ok := self.data[val]
  return ok
}

func(self Set[T]) Add(val T){
  self.data[val] = struct{}{}
}

func(self Set[T]) AddSlice(vals []T){
  for idx := range vals{
    self.Add(vals[idx])
  }
}
func EmptySet[T comparable]() Set[T]{
  return Set[T]{make(map[T]struct{})}
}
func MakeSet[T comparable](vals []T) Set[T]{
  set := EmptySet[T]() 
  for idx := range vals{
    set.Add(vals[idx])
  }
  return set
}

func(self Set[T]) Union(other Set[T]){
  for key, _ := range other.data{
    self.Add(key)
  }
}

func(self Set[T]) FromUnion(other Set[T]) Set[T]{
  set := EmptySet[T]()
  for key, _ := range self.data{
    set.Add(key)
  }
  for key, _ := range self.data{
    set.Add(key)
  }  
  return set
}

func(self Set[T]) Intersection(other Set[T]){
  for key, _ := range other.data{
    if !self.Contains(key){
      delete(self.data, key)
    }
  }
}

func(self Set[T]) FromIntersection(other Set[T]) Set[T]{
  set := EmptySet[T]()
  for key, _ := range other.data{
    if(self.Contains(key)){
      set.Add(key)
    }
  }
  return set
}



