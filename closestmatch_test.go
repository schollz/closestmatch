package closestmatch

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/schollz/closestmatch/test"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New(test.WordsToTest, []int{3})
	}
}

func BenchmarkSplitOne(b *testing.B) {
	cm := New(test.WordsToTest, []int{3})
	searchWord := test.SearchWords[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cm.splitWord(searchWord)
	}
}

func BenchmarkClosestOne(b *testing.B) {
	cm := New(test.WordsToTest, []int{3})
	searchWord := test.SearchWords[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cm.Closest(searchWord)
	}
}

func BenchmarkClosest3(b *testing.B) {
	cm := New(test.WordsToTest, []int{3})
	searchWord := test.SearchWords[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cm.ClosestN(searchWord, 3)
	}
}

func BenchmarkClosest30(b *testing.B) {
	cm := New(test.WordsToTest, []int{3})
	searchWord := test.SearchWords[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cm.ClosestN(searchWord, 30)
	}
}

func BenchmarkLargeFile(b *testing.B) {
	bText, _ := ioutil.ReadFile("test/books.list")
	wordsToTest := strings.Split(strings.ToLower(string(bText)), "\n")
	cm := New(wordsToTest, []int{3})
	searchWord := "island of a thod mirrors"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cm.Closest(searchWord)
	}
}

func BenchmarkFileLoad(b *testing.B) {
	bText, _ := ioutil.ReadFile("test/books.list")
	wordsToTest := strings.Split(strings.ToLower(string(bText)), "\n")
	cm := New(wordsToTest, []int{3, 4})
	cm.Save("test/books.list.cm.gz")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Load("test/books.list.cm.gz")
	}
}

func BenchmarkFileSave(b *testing.B) {
	bText, _ := ioutil.ReadFile("test/books.list")
	wordsToTest := strings.Split(strings.ToLower(string(bText)), "\n")
	cm := New(wordsToTest, []int{3, 4})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cm.Save("test/books.list.cm.gz")
	}
}

func ExampleMatching() {
	cm := New(test.WordsToTest, []int{2, 3})
	for _, searchWord := range test.SearchWords {
		fmt.Printf("'%s' matched '%s'\n", searchWord, cm.Closest(searchWord))
	}
	// Output:
	// 'cervantes don quixote' matched 'don quixote by miguel de cervantes saavedra'
	// 'mysterious afur at styles by christie' matched 'the mysterious affair at styles by agatha christie'
	// 'charles dickens' matched 'hard times by charles dickens'
	// 'william shakespeare' matched 'the complete works of william shakespeare by william shakespeare'
	// 'war by hg wells' matched 'the war of the worlds by h. g. wells'
}

func ExampleMatchingN() {
	cm := New(test.WordsToTest, []int{1, 2, 3})
	fmt.Println(cm.ClosestN("war by hg wells", 3))
	// Output:
	// [the war of the worlds by h. g. wells the time machine by h. g. wells the yellow wallpaper by charlotte perkins gilman]
}

func ExampleMatchingBigList() {
	bText, _ := ioutil.ReadFile("test/books.list")
	wordsToTest := strings.Split(strings.ToLower(string(bText)), "\n")
	cm := New(wordsToTest, []int{3})
	searchWord := "island of a thod mirrors"
	fmt.Println(cm.Closest(searchWord))
	// Output:
	// island of a thousand mirrors by nayomi munaweera
}

func TestAccuray(t *testing.T) {
	cm := New(test.WordsToTest, []int{3})
	accuracy := cm.Accuracy()
	if accuracy < 92 {
		t.Errorf("Accuracy should be higher than %2.1f", accuracy)
	}
}

func TestSaveLoad(t *testing.T) {
	cm := New(test.WordsToTest, []int{1, 2, 3})
	err := cm.Save("test.txt")
	if err != nil {
		t.Error(err)
	}
	cm2, err := Load("test.txt")
	if err != nil {
		t.Error(err)
	}
	if cm2.Closest("war by hg wells") != cm.Closest("war by hg wells") {
		t.Errorf("Differing answers")
	}
}
