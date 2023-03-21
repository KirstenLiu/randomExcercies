package main

import "fmt"

type Wedding struct{
	HasGloom int
	HasBride int
	HasBrideMaids int
	HasBestMen int
	HasBouquet bool
	HasPriest bool
	HasHost bool
	ThemeColor string
}

func NewWedding(options ...WeddingOption) *Wedding {
	const (
		defaultGloom = 1
		defaultHasBride = 1
		defaultHasBrideMaid = 2
		defaultHasBestMen = 2
		defaultHasBouquet = true
		defaultThemeColor = "Green"
	)
	w := &Wedding{
		HasGloom: defaultGloom,
		HasBride: defaultHasBride,
		HasBrideMaids: defaultHasBrideMaid,
		HasBestMen: defaultHasBestMen,
		HasBouquet: defaultHasBouquet,
		HasPriest: false,
		ThemeColor: defaultThemeColor,
	}

	for _, opt := range(options){
		opt(w)
	}

	return w
}

func main(){
	fmt.Println("A new wedding is going to happen...")
	w := NewWedding()
	fmt.Println("this wedding is like as follow:")
	fmt.Printf("There is %d gloom, %d bride, %d bridemaids, %d bestman, %v bouquent, %v priest, %v host, and theme color is %s",
	w.HasGloom, w.HasBride, w.HasBrideMaids, w.HasBestMen, w.HasBouquet, w.HasPriest, w.HasHost, w.ThemeColor)

	w = NewWedding(HasGloomOpt(2), HasBrideOpt(1), HasHost(true), HasBrideOpt(2))
	fmt.Println("this wedding is like as follow after calling options:")
	fmt.Printf("There is %d gloom, %d bride, %d bridemaids, %d bestman, %v bouquent, %v priest, %v host, and theme color is %s",
	w.HasGloom, w.HasBride, w.HasBrideMaids, w.HasBestMen, w.HasBouquet, w.HasPriest, w.HasHost, w.ThemeColor)

}

//The type is "A function that take a Wedding ptr as param"
type WeddingOption func(*Wedding)

func HasGloomOpt(numOfGloom int) WeddingOption{
	return func(w *Wedding) {
		w.HasGloom = numOfGloom
	} 
}

func HasBrideOpt(numOfBride int) WeddingOption{
	return func(w *Wedding) {
		w.HasBride = numOfBride
	}
}

func HasHost(host bool) WeddingOption{
	return func(w *Wedding){
		w.HasHost = host
	}
}

