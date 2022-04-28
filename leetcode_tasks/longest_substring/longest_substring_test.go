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
