package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lengthOfLongestSubstringTwoDistinct(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstringTwoDistinct("eceba"))
	assert.Equal(t, 5, lengthOfLongestSubstringTwoDistinct("ccaabbb"))
	assert.Equal(t, 6, lengthOfLongestSubstringTwoDistinct("cacaabbba"))
	assert.Equal(t, 0, lengthOfLongestSubstringTwoDistinct(""))
	assert.Equal(t, 3, lengthOfLongestSubstringTwoDistinct("aaa"))
	assert.Equal(t, 2, lengthOfLongestSubstringTwoDistinct("ab"))
	assert.Equal(t, 6, lengthOfLongestSubstringTwoDistinct("aaabbbcccddd"))
	assert.Equal(t, 10, lengthOfLongestSubstringTwoDistinct("abccbbcccaaacaca"))
}

func Test_lengthOfLongestSubstringTwoDistinctApproach2(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstringTwoDistinctApproach2("eceba"))
	assert.Equal(t, 5, lengthOfLongestSubstringTwoDistinctApproach2("ccaabbb"))
	assert.Equal(t, 6, lengthOfLongestSubstringTwoDistinctApproach2("cacaabbba"))
	assert.Equal(t, 0, lengthOfLongestSubstringTwoDistinctApproach2(""))
	assert.Equal(t, 3, lengthOfLongestSubstringTwoDistinctApproach2("aaa"))
	assert.Equal(t, 2, lengthOfLongestSubstringTwoDistinctApproach2("ab"))
	assert.Equal(t, 6, lengthOfLongestSubstringTwoDistinctApproach2("aaabbbcccddd"))
	assert.Equal(t, 10, lengthOfLongestSubstringTwoDistinctApproach2("abccbbcccaaacaca"))
}
