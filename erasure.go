package main

/*
#cgo CFLAGS: -mavx -mavx2
#cgo LDFLAGS: -lm
#include <stdio.h>
#include <immintrin.h>
int *AVX4K(const int a[1024], const int b[1024],  int c[1024]){
    int m = 0xFFFFFFFF;
    __m256i loada;
    __m256i loadb;
    __m256i loadc;
    for (int i = 0; i < 32; i++){
        loada = _mm256_lddqu_si256((__m256i*)(a+32*i));
        loadb = _mm256_loadu_si256((__m256i*)(b+32*i));
        loadc = _mm256_xor_si256(loada,loadb);
        _mm256_storeu_si256((__m256i*)(c+32*i),loadc);

        loada = _mm256_loadu_si256((__m256i*)(a+32*i+8));
        loadb = _mm256_loadu_si256((__m256i*)(b+32*i+8));
        loadc = _mm256_xor_si256(loada,loadb);
        _mm256_storeu_si256((__m256i*)(c+32*i+8),loadc);

        loada = _mm256_loadu_si256((__m256i*)(a+32*i+16));
        loadb = _mm256_loadu_si256((__m256i*)(b+32*i+16));
        loadc = _mm256_xor_si256(loada,loadb);
        _mm256_storeu_si256((__m256i*)(c+32*i+16),loadc);

        loada = _mm256_loadu_si256((__m256i*)(a+32*i+24));
        loadb = _mm256_loadu_si256((__m256i*)(b+32*i+24));
        loadc = _mm256_xor_si256(loada,loadb);
        _mm256_storeu_si256((__m256i*)(c+32*i+24),loadc);
    }
    return 0;
}

int *AVX4M(const int a[1024*1024], const int b[1024*1024],  int c[1024*1024]){
    for(int i =0; i<1024; i++){
        AVX4K(a+i*1024,b+i*1024,c+i*1024);
    }
    return 0;
}
*/
import "C"
import (
	"unsafe"
)

func Xor4M(da [1024 * 1024]int32, db [1024 * 1024]int32, dc [1024 * 1024]int32) {
	a0 := (*C.int)(unsafe.Pointer(&da[0]))
	b0 := (*C.int)(unsafe.Pointer(&db[0]))
	c0 := (*C.int)(unsafe.Pointer(&dc[0]))
	C.AVX4M(a0, b0, c0)
}

func Erasure() {

}
