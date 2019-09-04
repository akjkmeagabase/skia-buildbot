package search

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	assert "github.com/stretchr/testify/require"
	"go.skia.org/infra/go/fileutil"
	"go.skia.org/infra/go/gcs/gcs_testutils"
	"go.skia.org/infra/go/paramtools"
	"go.skia.org/infra/go/testutils"
	"go.skia.org/infra/go/testutils/unittest"
	"go.skia.org/infra/go/util"
	"go.skia.org/infra/golden/go/types"
)

const (
	// TEST_STORAGE_DIR_SEARCH_API is the path in the testdata bucket where
	// the test data files are stored.
	TEST_STORAGE_DIR_SEARCH_API = "gold-testdata"

	// TEST_DATA_DIR_PARSE is the directory where test data for the parse
	// function are downloaded.
	TEST_DATA_DIR_PARSE = "testdata_parse"

	// QUERIES_FNAME_SEARCH_API contains the file name of the list of queries
	// that were extracted from the Gold application log.
	QUERIES_FNAME_SEARCH_API = "live_queries.txt"
)

func TestParseCTQuery(t *testing.T) {
	unittest.SmallTest(t)
	testQuery := CTQuery{
		RowQuery: &Query{
			Pos:            true,
			Neg:            false,
			Head:           true,
			Unt:            true,
			IncludeIgnores: true,
			QueryStr:       "source_type=gm&param=value",
			Limit:          20,
		},
		ColumnQuery: &Query{
			Pos:            true,
			Neg:            false,
			Head:           true,
			Unt:            true,
			IncludeIgnores: true,
			QueryStr:       "source_type=gm&param=value",
		},

		Match: []string{"gamma_correct"},
	}

	jsonBytes, err := json.Marshal(&testQuery)
	assert.NoError(t, err)

	var ctQuery CTQuery
	assert.NoError(t, ParseCTQuery(ioutil.NopCloser(bytes.NewBuffer(jsonBytes)), 9, &ctQuery))
	exp := url.Values{"source_type": []string{"gm"}, "param": []string{"value"}}
	assert.True(t, util.In(types.PRIMARY_KEY_FIELD, ctQuery.Match))
	assert.Equal(t, exp, ctQuery.RowQuery.Query)
	assert.Equal(t, exp, ctQuery.ColumnQuery.Query)
	assert.Equal(t, int32(9), ctQuery.ColumnQuery.Limit)

	testQuery.RowQuery.QueryStr = ""
	jsonBytes, err = json.Marshal(&testQuery)
	assert.NoError(t, err)
	assert.Error(t, ParseCTQuery(ioutil.NopCloser(bytes.NewBuffer(jsonBytes)), 10, &ctQuery))
}

func TestParseQuery(t *testing.T) {
	unittest.SmallTest(t)
	assertQueryValidity(t, true, "fdiffmax=-1&fref=false&frgbamax=-1&head=true&include=false&issue=2370153003&limit=50&match=gamma_correct&match=name&metric=combined&neg=false&pos=false&query=source_type%3Dgm&sort=desc&unt=true")
	assertQueryValidity(t, true, "fdiffmax=-1&fref=false&frgbamax=-1&head=true&include=false&limit=50&match=gamma_correct&match=name&metric=combined&neg=false&pos=false&query=source_type%3Dgm&sort=desc&unt=true")
	assertQueryValidity(t, false, "fdiffmax=abc&fref=false&frgbamax=-1&head=true&include=false&limit=50&")

	q := &Query{}
	err := clearParseQuery(q, "fdiffmax=-1&fref=false&frgbamax=-1&head=true&include=false&issue=2370153003&limit=50&match=gamma_correct&match=name&metric=combined&neg=false&pos=false&query=source_type%3Dgm&sort=desc&unt=true")
	assert.NoError(t, err)

	assert.Equal(t, &Query{
		Metric:         "combined",
		Sort:           "desc",
		Match:          []string{"gamma_correct", "name"},
		BlameGroupID:   "",
		Pos:            false,
		Neg:            false,
		Head:           true,
		Unt:            true,
		IncludeIgnores: false,
		QueryStr:       "",
		Query: url.Values{
			"source_type": []string{"gm"},
		},
		RQueryStr:       "",
		RQuery:          paramtools.ParamSet{},
		ChangeListID:    "2370153003",
		DeprecatedIssue: 2370153003,
		PatchsetsStr:    "",
		Patchsets:       []int64(nil),
		IncludeMaster:   false,
		FCommitBegin:    "",
		FCommitEnd:      "",
		FRGBAMin:        0,
		FRGBAMax:        -1,
		FDiffMax:        -1,
		FGroupTest:      "",
		FRef:            false,
		Offset:          0,
		Limit:           50,
		NoDiff:          false,
		NewCLStore:      false,
	}, q)
}

// TODO(kjlubick): replace this test with one that 1) reads from testdata/ and
// 2) explicitly has a list of valid and invalid queries.
func TestParseQueryLarge(t *testing.T) {
	unittest.LargeTest(t)

	// Reuse the paths from the SearchAPI benchmarks.
	cloudQueriesPath := TEST_STORAGE_DIR_SEARCH_API + "/" + QUERIES_FNAME_SEARCH_API + ".gz"
	localQueriesPath := TEST_DATA_DIR_PARSE + "/" + QUERIES_FNAME_SEARCH_API
	defer testutils.RemoveAll(t, TEST_DATA_DIR_PARSE)

	// Download the list of queries.
	assert.NoError(t, gcs_testutils.DownloadTestDataFile(t, gcs_testutils.TEST_DATA_BUCKET, cloudQueriesPath, localQueriesPath))

	// Load the list of of live queries.
	queries, err := fileutil.ReadLines(localQueriesPath)
	assert.NoError(t, err)

	q := &Query{}
	wrongQueries := 0
	for _, qStr := range queries {
		err := clearParseQuery(q, qStr)
		if err != nil {
			wrongQueries++
		}
	}

	// Accept as long as 10% of all queries are wrong.
	errFraction := float64(wrongQueries) / float64(len(queries))
	assert.True(t, errFraction < 0.1, fmt.Sprintf("Fraction of wrong queries is too high: %f > %f", errFraction, 0.1))
}

func assertQueryValidity(t *testing.T, isCorrect bool, qStr string) {
	assertFn := assert.NoError
	if !isCorrect {
		assertFn = assert.Error
	}
	q := &Query{}
	assertFn(t, clearParseQuery(q, qStr))
}

func clearParseQuery(q *Query, qStr string) error {
	*q = Query{}
	r, err := http.NewRequest("GET", "/?"+qStr, nil)
	if err != nil {
		return err
	}
	return ParseQuery(r, q)
}
