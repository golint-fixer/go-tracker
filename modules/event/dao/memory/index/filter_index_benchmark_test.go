package index

import (
	"math/rand"
	"testing"
	"time"

	eventEntity "github.com/index0h/go-tracker/modules/event/entity"
	visitEntity "github.com/index0h/go-tracker/modules/visit/entity"
	"github.com/index0h/go-tracker/share/types"
	"github.com/index0h/go-tracker/share/uuid"
)

func Benchmark_FilterIndex_FindAllByVisit_3(b *testing.B) {
	filterIndexFindAllByVisit(3, b)
}

func Benchmark_FilterIndex_FindAllByVisit_5(b *testing.B) {
	filterIndexFindAllByVisit(5, b)
}

func Benchmark_FilterIndex_FindAllByVisit_10(b *testing.B) {
	filterIndexFindAllByVisit(10, b)
}

func Benchmark_FilterIndex_FindAllByVisit_15(b *testing.B) {
	filterIndexFindAllByVisit(15, b)
}

var filterIndexSymbols = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func filterIndexFindAllByVisit(countKeys uint, b *testing.B) {
	b.StopTimer()

	rand.Seed(time.Now().UTC().UnixNano())

	events := filterIndexGenerateEvents(uint(b.N), countKeys)
	visits := filterIndexGenerateVisits(uint(b.N), countKeys)

	index := NewFilterIndex()
	index.Refresh(events)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = index.FindAllByFields(visits[i].Fields())
	}
}

func filterIndexGenerateVisits(countVisits uint, countData uint) []*visitEntity.Visit {
	result := make([]*visitEntity.Visit, countVisits)

	for i := uint(0); i < countVisits; i++ {
		fileds := filterIndexGenerateKeyValue(countData)

		result[i], _ = visitEntity.NewVisit(uuid.New().Generate(), int64(0), uuid.New().Generate(), "", fileds)
	}

	return result
}

func filterIndexGenerateEvents(count uint, countData uint) []*eventEntity.Event {
	result := make([]*eventEntity.Event, count)

	for i := uint(0); i < count; i++ {
		filters := filterIndexGenerateKeyValue(countData)

		result[i], _ = eventEntity.NewEvent(uuid.New().Generate(), true, types.Hash{}, filters)
	}

	return result
}

func filterIndexGenerateKeyValue(count uint) (result types.Hash) {
	result = make(types.Hash, count)

	for i := uint(0); i < count; i++ {
		result[filterIndexGenerateString()] = filterIndexGenerateString()
	}

	return result
}

func filterIndexGenerateString() string {
	count := 5
	result := make([]byte, count)

	var number int

	for i := 0; i < count; i++ {
		number = rand.Intn(len(filterIndexSymbols) - 1)

		result[i] = filterIndexSymbols[number]
	}

	return string(result)
}
