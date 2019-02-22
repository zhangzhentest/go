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

func (p PairList) Swap(i,j int)  {p[i],p[j] = p[j],p[i]}

func (p PairList) Len() int {return len(p)}

func (p PairList) Less(i,j int) bool {return p[j].Value < p[i].Value}

func SplitOnNonLetters(s string) []string{
	//judge char is a letter?
	notALetter := func(char rune) bool {return !unicode.IsLetter(char)}
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

func (wordcount WordCount) Report(){
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
	gap := wordWidth + frequencyWidth - len("Word") - len("Frequency")
	//%*s gap->* %gaps  space numbers
	fmt.Printf("Word %*s%s\n",gap," ","Frwquency")
	for _,word := range words{
		// - left duiqi
		fmt.Printf("%-*s%*d\n",wordWidth,word,frequencyWidth,wordcount[word])
	}
}

func (wordcount WordCount) SortReport(){
	p := make(PairList,len(wordcount))
	i := 0
	for k,v := range wordcount{
		p[i] = Pair{k,v}
		i++
	}
	sort.Sort(p)

	wordWidth,frequencyWidth := 0,0
	for _,pair := range p{
		word,frequency := pair.Key,pair.Value
		//get unicode character numbers
		if width := utf8.RuneCountInString(word);width > wordWidth{
			wordWidth = width
		}
		if width := len(fmt.Sprint(frequency));width > frequencyWidth{
			frequencyWidth = width
		}
	}
	gap := wordWidth + frequencyWidth - len("Word") - len("Frequency")
	//%*s gap->* %gaps  space numbers
	fmt.Printf("Word %*s%s\n",gap," ","Frequency")
	for _,pair := range p{
		// - left duiqi
		fmt.Printf("%-*s %*d\n",wordWidth,pair.Key,frequencyWidth,pair.Value)
	}

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
		line,err := reader.ReadString('\n')
		//delete before string space and end of string space
		for _,word := range SplitOnNonLetters(strings.TrimSpace(line)){
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
		go func(done chan <- struct{},results chan <- Pair,filename string){
			wordcount := make(WordCount)
			wordcount.UpdateFreq(filename)
			for k,v := range wordcount{
				pair := Pair{k,v}
				results <- pair
			}
			done <- struct{}{}
		}(done,results,files[i])
		i++
	}
	
	for working := len(files);working > 0;{
		select{
			case pair := <-results:
				wordcount[pair.Key] += pair.Value
			case <-done:
				working--
		}
	}

DONE:
	for{
		select{	
		case pair := <-results:
			wordcount[pair.Key] += pair.Value
		default:
			break DONE
		}
	}
	close(results)
	close(done)
}
			

	
		
	
