package main

import "fmt"

func uangKakak(tabPertama int, n int) int {
    total := tabPertama
    tabunganMingguIni := tabPertama

    for i := 1; i < n; i++ {
        tabunganMingguIni += 2500
        total += tabunganMingguIni
    }

    return total
}

func main() {
    var tabPertama, n, jumlah int
    fmt.Scan(&tabPertama, &n)
    jumlah = uangKakak(tabPertama, n)
    fmt.Println(jumlah)
}
