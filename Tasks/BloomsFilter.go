package main

import (
	"hash"
	"hash/fnv"
)

type Interface interface {
	Add(item []byte)     // Adds the item into the Set
	Test(item []byte) bool  // Check if items is maybe in the Set
}



// BloomFilter probabilistic data structure definition
type BloomFilter struct {
	bitset []bool      // The bloom-filter bitset
	k      uint         // Number of hash values
	n      uint         // Number of elements in the filter
	m      uint         // Size of the bloom filter bitset
	hashFuncs[]hash.Hash64           // The hash functions
}

// Returns a new BloomFilter object,
func New(size) *BloomFilter {
	return &BloomFilter{
		bitset: make([]bool, size),
		k: 3,  // we have 3 hash functions for now
		m: size,
		n: uint(0),
		hashfns:[]hash.Hash64{murmur3.New64(),fnv.New64(),fnv.New64a()},
	}
}

func (bf *BloomFilter) Add(item []byte) {
	hashes := bf.hashValues(item)
	i := uint(0)

	for {
		if i >= bf.k {
			break
		}

		position := uint(hashes[i]) % bf.m
		bf.bitset[uint(position)] = true

		i+= 1
	}

	bf.n += 1
}

func (bf *BloomFilter) hashValues(item []byte) []uint64  {
	var result []uint64

	for _, hashFunc := range bf.hashfns {
		hashFunc.Write(item)
		result = append(result, hashFunc.Sum64())
		hashFunc.Reset()
	}

	return result
}


