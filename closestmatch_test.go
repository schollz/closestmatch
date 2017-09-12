package closestmatch

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/Yugloocamai/closestmatch/test"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New(test.BooksToTest, []int{3})
	}
}

func BenchmarkSplitOne(b *testing.B) {
	cm := New(test.BooksToTest, []int{3})
	searchWord := test.SearchWords[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cm.splitWord(searchWord)
	}
}

func BenchmarkClosestOne(b *testing.B) {
	bText, _ := ioutil.ReadFile("test/books.list")
	books := test.GetBooks(string(bText))
	cm := New(books, []int{3})
	searchWord := test.SearchWords[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cm.Closest(searchWord)
	}
}

func BenchmarkClosest3(b *testing.B) {
	bText, _ := ioutil.ReadFile("test/books.list")
	books := test.GetBooks(string(bText))
	cm := New(books, []int{3})
	searchWord := test.SearchWords[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cm.ClosestN(searchWord, 3)
	}
}

func BenchmarkClosest30(b *testing.B) {
	bText, _ := ioutil.ReadFile("test/books.list")
	books := test.GetBooks(string(bText))
	cm := New(books, []int{3})
	searchWord := test.SearchWords[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cm.ClosestN(searchWord, 30)
	}
}

func BenchmarkFileLoad(b *testing.B) {
	bText, _ := ioutil.ReadFile("test/books.list")
	books := test.GetBooks(string(bText))
	cm := New(books, []int{3, 4})
	cm.Save("test/books.list.cm.gz")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Load("test/books.list.cm.gz")
	}
}

func BenchmarkFileSave(b *testing.B) {
	bText, _ := ioutil.ReadFile("test/books.list")
	books := test.GetBooks(string(bText))
	cm := New(books, []int{3, 4})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cm.Save("test/books.list.cm.gz")
	}
}

func ExampleMatchingSmall() {
	loveCats := make(map[string]interface{})
	loveCats["love"] = map[string]string{"name": "love"}
	loveCats["loving"] = map[string]string{"name": "loving"}
	loveCats["cat"] = map[string]string{"name": "cat"}
	loveCats["kit"] = map[string]string{"name": "kit"}
	loveCats["cats"] = map[string]string{"name": "cats"}
	cm := New(loveCats, []int{4})
	fmt.Println(cm.splitWord("love"))
	fmt.Println(cm.splitWord("kit"))
	fmt.Println(cm.Closest("kit"))
	// Output:
	// map[love:{}]
	// map[kit:{}]
	// kit

}

func ExampleMatchingSimple() {

	booksLines := strings.Split(strings.ToLower(test.Books), "\n")
	wordsToTest := make(map[string]interface{})
	for _, v := range booksLines {
		wordsToTest[v] = map[string]string{"words": v}
	}
	cm := New(wordsToTest, []int{3})
	for _, searchWord := range test.SearchWords {
		fmt.Printf("'%s' matched '%s'\n", searchWord, cm.Closest(searchWord))
	}
	// Output:
	// 'cervantes don quixote' matched 'don quixote by miguel de cervantes saavedra'
	// 'mysterious afur at styles by christie' matched 'the mysterious affair at styles by agatha christie'
	// 'hard times by charles dickens' matched 'hard times by charles dickens'
	// 'complete william shakespeare' matched 'the complete works of william shakespeare by william shakespeare'
	// 'War by HG Wells' matched 'the war of the worlds by h. g. wells'

}

func ExampleMatchingN() {
	cm := New(test.BooksToTest, []int{4})
	results := cm.ClosestN("war h.g. wells", 3)
	var slice []string
	for _, v := range results {
		slice = append(slice, v.(map[string]string)["name"])
	}
	fmt.Println(slice)
	// Output:
	// [the war of the worlds by h. g. wells the time machine by h. g. wells war and peace by graf leo tolstoy]
}

func ExampleMatchingBigList() {
	bText, _ := ioutil.ReadFile("test/books.list")
	books := test.GetBooks(string(bText))
	cm := New(books, []int{3})
	searchWord := "island of a thod mirrors"
	fmt.Println(cm.Closest(searchWord))
	// Output:
	// island of a thousand mirrors by nayomi munaweera
}

func ExampleMatchingCatcher() {
	bText, _ := ioutil.ReadFile("test/catcher.txt")
	books := test.GetBooks(string(bText))
	cm := New(books, []int{5})
	searchWord := "catcher in the rye by jd salinger"
	for i, match := range cm.ClosestN(searchWord, 3) {
		if i == 2 {
			fmt.Println(match.(map[string]string)["name"])
		}
	}
	// Output:
	// the catcher in the rye by j.d. salinger
}

func ExampleMatchingPotter() {
	bText, _ := ioutil.ReadFile("test/potter.txt")
	books := test.GetBooks(string(bText))
	cm := New(books, []int{5})
	searchWord := "harry potter and the half blood prince by j.k. rowling"
	for i, match := range cm.ClosestN(searchWord, 3) {
		if i == 1 {
			fmt.Println(match.(map[string]string)["name"])
		}
	}
	// Output:
	//  harry potter and the order of the phoenix (harry potter, #5, part 1) by j.k. rowling
}

func TestAccuracyBookWords(t *testing.T) {
	bText, _ := ioutil.ReadFile("test/books.list")
	books := test.GetBooks(string(bText))
	cm := New(books, []int{4, 5})
	accuracy := cm.AccuracyMutatingWords()
	fmt.Printf("Accuracy with mutating words in book list:\t%2.1f%%\n", accuracy)
}

func TestAccuracyBookLetters(t *testing.T) {
	bText, _ := ioutil.ReadFile("test/books.list")
	books := test.GetBooks(string(bText))
	cm := New(books, []int{5})
	accuracy := cm.AccuracyMutatingLetters()
	fmt.Printf("Accuracy with mutating letters in book list:\t%2.1f%%\n", accuracy)
}

func TestAccuracyDictionaryLetters(t *testing.T) {
	bText, _ := ioutil.ReadFile("test/popular.txt")
	words := strings.Split(strings.ToLower(string(bText)), "\n")
	wordsToTest := make(map[string]interface{})
	for _, v := range words {
		wordsToTest[v] = map[string]string{"word": v}
	}
	cm := New(wordsToTest, []int{2, 3, 4})
	accuracy := cm.AccuracyMutatingWords()
	fmt.Printf("Accuracy with mutating letters in dictionary:\t%2.1f%%\n", accuracy)
}

func TestSaveLoad(t *testing.T) {
	bText, _ := ioutil.ReadFile("test/books.list")
	books := test.GetBooks(string(bText))
	type TestStruct struct {
		cm *ClosestMatch
	}
	tst := new(TestStruct)
	tst.cm = New(books, []int{5})
	err := tst.cm.Save("test.gob")
	if err != nil {
		t.Error(err)
	}

	tst2 := new(TestStruct)
	tst2.cm, err = Load("test.gob")
	if err != nil {
		t.Error(err)
	}
	answer2 := tst2.cm.Closest("war of the worlds")
	answer1 := tst.cm.Closest("war of the worlds")
	if answer1 != answer2 {
		t.Errorf("Differing answers: '%s' '%s'", answer1, answer2)
	}
}
