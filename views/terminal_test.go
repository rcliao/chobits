package views

import "testing"

func TestConvertIntToArray(t *testing.T) {
	var content = `
    #
    #
    #
    #
    #
`
	var expected = []string{"    #", "    #", "    #", "    #", "    #"}
	var actual = convertIntToArray(content)
	compareArray(expected, actual, "TestConvertIntToArray", t)
}

func TestConcatArray(t *testing.T) {
	var array1 = []string{"    #", "    #", "    #", "    #", "    #"}
	var array2 = []string{"    #", "    #", "    #", "    #", "    #"}
	var expected = []string{"    #     #", "    #     #", "    #     #", "    #     #", "    #     #"}
	var actual = concatArrayHorizontal(array1, array2)
	compareArray(expected, actual, "TestConcatArray", t)
}

func compareArray(arr1, arr2 []string, context string, t *testing.T) {
	if len(arr1) != len(arr2) {
		t.Errorf("%s: Expected array length %d but but got %d", context, len(arr1), len(arr2))
		return
	}

	for i, e := range arr2 {
		if arr1[i] != e {
			t.Errorf("%s: Expected '%s' but got '%s'", context, e, arr1[i])
		}
	}
}
