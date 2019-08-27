package commands

import (
	"fmt"
	"github.com/taqboz/tombo_gdn/cli/data"
	"github.com/taqboz/tombo_gdn/cli/helpers"
	"github.com/taqboz/tombo_gdn/internal/app/tombo_gdn/pkg"
)

func Check(input string, n int) error {
	// クローリング
	for i := 0; i < n; i ++ {
		if err := getPaths(input); err != nil {
			return err
		}
	}
	fmt.Println("complete index")

	// 各PATHの確認
	if err := checkStatus(input); err != nil {
		return err
	}

	return nil
}

func getPaths(input string) error {
	if err := getDocument(input); err != nil {
		return err
	}

	fmt.Println("check complete")
	data.FindNotCheckedPath()
	return nil
}

func getDocument(input string) error {

	//var eg errgroup.Group
	//c := make(chan struct{}, 10)
	fmt.Println(len(data.NotCheckedPaths))
	for _, v := range data.NotCheckedPaths {
		//v2 := v
		//eg.Go(func() error {
		//	//c <- struct{}{}
		//	//defer func() {
		//	//	<-c
		//	//}()

			a, err := pkg.ResolveURL(input, v)
			fmt.Println(a)
			if err != nil {
				return err
			}

			doc, st, err := helpers.GetRequestBasicAuth(a)
			data.AddPath(&data.Path{Path:v, Status:st, Doc:doc})

			//return nil
		//})
	}

	//if err := eg.Wait(); err != nil {
	//	return err
	//}

	data.NotCheckedPaths = []string{}
	return nil
}

var ErrLinks = make([]*ErrLink, 0)

type ErrLink struct {
	Destination *LinkInfo
	Sources []string
}

type LinkInfo struct {
	Path string
	Status int
}

func checkStatus(input string) error {
	e := make([]*LinkInfo, 0)
	for _, v := range data.CheckedPaths {
		if v.Status != 200 {
			e = append(e, &LinkInfo{v.Path,v.Status})
		}
	}


	for _, v := range e {
		add := &ErrLink{Destination:v, Sources:[]string{}}
		for _, v2 := range data.CheckedPaths {
			if pkg.IsContain(v2.ErrLinks, v.Path) {
				fmt.Println(v2.Path)
				add.Sources = append(add.Sources, v2.Path)
			}
		}
		ErrLinks = append(ErrLinks, add)
	}

	return nil
}


