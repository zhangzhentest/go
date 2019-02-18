package wordcount

import(
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Pair struct{
	Key string
	Value int
}

type PairList []Pair

func (p PairList) Swap(i,j int)  {p[i],p[j] = [j],p[i]}

func (p PairList) Len() int {return len(p)}

func (p PairList) Less(i,j int) {return p[j].Value < p[i].Value}

func SplitOnNonLetters(s string) []string{
	//judge char is a letter?
	notALetter := func(char rune) bool {return !uncoid.IsLetter(char)}
	//if false -->string[i] true not in string[i] 
	return strings.FieldsFunc(s,notALetter)
}

type WordCount map[string]int

func (source WordCount) Merge(wordcount WordCount) WordCount{
	for k,v := range wordcount{
		source[k] += v
	}
	return source
}

/*func (wordcount WordCount) Report(){
	words := make([]string,0,len(wordcount))
	wordWidth,frequencyWidth := 0,0
	for word,frequency := range wordcount{
		words = append(words,word)
		//get unicode character numbers
		if width := utf8.RuneCountInString(word);width > wordWidth{
			wordWidth = width
		}
		if width := len(fmt.Sprint(frequency));width > frequencyWidth{
			frequencyWidth = frequency
		}
	}
	sort.Strings(words)
*/

func (wordcount WordCount) SortReport(){
	p := make([]PairList,len(wordcount))
	i := 0
	for k,v := range wordcount{
		p[i] = pair{k,v}
		i++
	}
	sort.Sort(p)
}

func (wordcount WordCount) UpdateFreq(filename string){
	var file *os.File  //os.File * type
	var err error
	if file,err = os.Open(filename);err != nil{
		log.Println("failed to open the file:",err)
		return
	}
	defer file.Close()
	
	//read file to buf
	reader := bufio.NewReader(file)
	for{
		//get '\n' before data
		line,err := reader.ReadString("\n")
		//delete before string space and end of string space
		for _,word := SplitOnNonLetters(strings.TrimSpace(line)){
			wordcount[strings.ToLower(word)] ++
		}
		if err != nil{
			if err !=io.EOF{
				log.Println("failed to finish reading the file:",err)
			}
			break
		}
	}
}

func (wordcount WordCount) WordFreqCounter(files []string){
	results := make(chan Pair,len(files))
	done := make(chan struct{},len(files))
	for i := 0;i<len(files);{
		go func(done chan,results chan,filename string){
			wordcount := make(WordCount)
			wordcount.UpdateFreq(filename)
			for k,v := range wordcount{
				pair := Pair{k,v}
				results <- pair
			}
			done <- struct{}{}
		}(done,results,files[i]
			

	
		
	
